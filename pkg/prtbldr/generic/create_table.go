package generic

import (
	"github.com/yo3jones/gosql/pkg/prtbldr"
)

type (
	// CreateTableFactory is a prtbldr.SQLPartBuilderFactory that creates
	// builders for CREATE TABLE statement parts.
	CreateTableFactory struct{}

	// CreateTableBuilder is a prtbldr.SQLPartBuilder for building
	// CREATE TABLE statements in the generic dialect.
	CreateTableBuilder struct {
		part *prtbldr.CreateTablePart
	}

	// CreateTableTemporaryBuilder is a prtbldr.SQLPartBuilder for building
	// the TEMPORARY part of a CREATE TABLE statemet in the generic dialect.
	CreateTableTemporaryBuilder struct {
		part *prtbldr.CreateTablePart
	}

	// CreateTableIfNotExistsBuilder is a prtbldr.SQLPartBuilder for building
	// the IF NOT EXÎSTS part of a CREATE TABLE statement in the generic
	// dialect.
	CreateTableIfNotExistsBuilder struct {
		part *prtbldr.CreateTablePart
	}
)

var (
	_ prtbldr.SQLPartBuilderFactory = (*CreateTableFactory)(nil)
	_ prtbldr.SQLPartBuilder        = (*CreateTableBuilder)(nil)
	_ prtbldr.SQLPartBuilder        = (*CreateTableTemporaryBuilder)(nil)
	_ prtbldr.SQLPartBuilder        = (*CreateTableIfNotExistsBuilder)(nil)
)

// NewPartBuilder returns a new SQLPartBuilder for the given partType
// and part.
//nolint:ireturn//factory has to return an interface
func (factory *CreateTableFactory) NewPartBuilder(
	partType prtbldr.SQLPartType,
	part prtbldr.SQLPart,
) prtbldr.SQLPartBuilder {
	if part, ok := part.(*prtbldr.CreateTablePart); ok {
		return factory.newCreateTableTypeBuilder(partType, part)
	}

	return nil
}

// newCreateTableTypeBuilder returns the builder that builds an
// prtbldr.CreateTablePart.
//nolint:ireturn//factory has to return an interface
//revive:disable:cyclomatic Need complexity for switch statement
func (*CreateTableFactory) newCreateTableTypeBuilder(
	partType prtbldr.SQLPartType,
	part *prtbldr.CreateTablePart,
) prtbldr.SQLPartBuilder {
	//nolint:exhaustive//can only be a subset of types
	switch partType {
	case prtbldr.CreateTable:
		return &CreateTableBuilder{part}
	case prtbldr.CreateTableTemporary:
		return &CreateTableTemporaryBuilder{part}
	case prtbldr.CreateTableNameIfNotExist:
		return &CreateTableIfNotExistsBuilder{part}
	}

	return nil
}

// revive:enable:cyclomatic

func (builder *CreateTableBuilder) Build(
	res prtbldr.SQLPartBuilderResult,
	opts ...prtbldr.SQLPartBuilderOption,
) {
	res.PrintWithOptions(opts, "CREATE")

	// build TEMPORARY
	res.Build(
		prtbldr.CreateTableTemporary,
		builder.part,
		&prtbldr.PrefixBuilderOption{Prefix: " "},
	)

	// build IF NOT EXISTS
	res.Build(
		prtbldr.CreateTableNameIfNotExist,
		builder.part,
		&prtbldr.PrefixBuilderOption{Prefix: " "},
	)

	res.Print(" TABLE")

	// build the table name
	res.Build(
		prtbldr.TableName,
		builder.part.TableName,
		&prtbldr.PrefixBuilderOption{Prefix: " "},
	)
}

func (builder *CreateTableTemporaryBuilder) Build(
	res prtbldr.SQLPartBuilderResult,
	opts ...prtbldr.SQLPartBuilderOption,
) {
	if builder.part.Temporary {
		res.PrintWithOptions(opts, "TEMPORARY")
	}
}

func (builder *CreateTableIfNotExistsBuilder) Build(
	res prtbldr.SQLPartBuilderResult,
	opts ...prtbldr.SQLPartBuilderOption,
) {
	if builder.part.IfNotExists {
		res.PrintWithOptions(opts, "IF NOT EXISTS")
	}
}
