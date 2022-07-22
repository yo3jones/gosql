package generic

import (
	"github.com/yo3jones/gosql/pkg/prtbldr"
)

type (
	// CommonFactory is a prtbldr.SQLPartBuilderFactory for creating builders
	// of common parts.
	CommonFactory struct{}

	// TableNameBuilder is a prtbldr.SQLPartBuilder for building a table name
	// in the generic dialect.
	TableNameBuilder struct {
		part *prtbldr.TableNamePart
	}

	// ColumnDefnBuilder is a prtbldr.SQLPartBuilder for building a column
	// definition in the generic dialect.
	ColumnDefnBuilder struct {
		part *prtbldr.ColumnDefnPart
	}
)

var (
	_ prtbldr.SQLPartBuilderFactory = (*CommonFactory)(nil)
	_ prtbldr.SQLPartBuilder        = (*TableNameBuilder)(nil)
	_ prtbldr.SQLPartBuilder        = (*ColumnDefnBuilder)(nil)
)

// NewPartBuilder returns a new SQLPartBuilder for the given partType
// and part.
//nolint:ireturn//factory needs to return an interface
func (*CommonFactory) NewPartBuilder(
	_ prtbldr.SQLPartType,
	part prtbldr.SQLPart,
) prtbldr.SQLPartBuilder {
	if part, ok := part.(*prtbldr.TableNamePart); ok {
		return &TableNameBuilder{part}
	}

	return nil
}

func (builder *TableNameBuilder) Build(
	res prtbldr.SQLPartBuilderResult,
	opts ...prtbldr.SQLPartBuilderOption,
) {
	var (
		part   = builder.part
		name   = part.Name
		schema = part.Schema
	)

	if schema != "" {
		res.PrintWithOptions(opts, "%s.%s", schema, name)
	} else {
		res.PrintfWithOptions(opts, name)
	}
}

func (*TableNameBuilder) Type() prtbldr.SQLPartType {
	return prtbldr.TableName
}

func (builder *ColumnDefnBuilder) Build(
	res prtbldr.SQLPartBuilderResult,
	opts ...prtbldr.SQLPartBuilderOption,
) {
}
