// package prtbldr provides how parts of SQL statements are built.
//
// package sqlpartbldr is not ment to be used, users should interact with
// sqlbldr. This package should be used for implementors of different SQL
// dialects.
package prtbldr

import (
	"fmt"
	"strings"
)

type (
	// SQLPartBuilder describes a Build method which writes the SQL part to the
	// given writer.
	SQLPartBuilder interface {
		// Build writes the string of the SQL part to the given writer. It
		// returns a non null err in the case that an error occurred during the
		// building of the part.
		Build(
			res SQLPartBuilderResult,
			opts ...SQLPartBuilderOption,
		)
	}

	// SQLPartBuilderOption describes an option that can be passed to the
	// SQLPartBuilder Build method.
	SQLPartBuilderOption interface {
		// isSQLPartBuilderOption is only used to help development to make the
		// SQLBuilderOption type safe.
		isSQLPartBuilderOption()
	}

	// SQLPartBuilderFactory descibes a factory that will create SQLPartBuilders
	// for the given partType and part.
	SQLPartBuilderFactory interface {
		// NewPartBuilder returns a new SQLPartBuilder for the given partType
		// and part.
		NewPartBuilder(partType SQLPartType, part SQLPart) SQLPartBuilder
	}

	// SQLPart describes a SQL Part.
	SQLPart interface {
		// Type returns the SQLPartType this SQLPart represents.
		Type() (partType SQLPartType)
	}

	// SQLPartType specifies the type of a SQL part.
	SQLPartType uint8

	// SQLPartBuilderResult is the object used to store the result of
	// SQLPartBuilders. The result can store the string SQL as well as all
	// the errors that might have occurred during the building of many parts.
	SQLPartBuilderResult interface {
		// Print is a convenience method for printing a formatted string to
		// the SQL string. If an error occurs while writing it will be added
		// to the result.
		Print(a ...any) (n int)

		// Printf is a convenience method for printing a formatted string to
		// the SQL string. If an error occurs while writing it will be added
		// to the result.
		Printf(format string, a ...any) (n int)

		// PrintPrefix is a convenience method for printing any prefix
		// specified in the given opts.
		PrintPrefix(opts []SQLPartBuilderOption) (n int)

		// PrintWithOptions is a convenience method for printing a formatted
		// string to the SQL string.
		//
		// This method will also handle any pretty printing options passed in.
		//
		// If an error occurs while writing it will be added
		// to the result.
		PrintWithOptions(opts []SQLPartBuilderOption, a ...any) (n int)

		// PrintfWithOptions is a convenience method for printing a formatted
		// string to the SQL string.
		//
		// This method will also handle any pretty printing options passed in.
		//
		// If an error occurs while writing it will be added
		// to the result.
		PrintfWithOptions(
			opts []SQLPartBuilderOption,
			format string,
			a ...any,
		) (n int)

		// AppendError will store the given error so it can be viewed in
		// aggregate from the result.
		AppendError(err error)

		// Build is a convenience method for building the given part.
		Build(partType SQLPartType, part SQLPart, opts ...SQLPartBuilderOption)

		// Result returns the SQL built by many SQLPartBuilders. If there were
		// at least one error appended during the building, err will be non
		// nil.
		Result() (sql string, err error)
	}

	// ISQLPartBuilderResult implements [prtbldr.SQLPartBuilder].
	//
	// SQLPartBuilderResult is the object used to store the result of
	// SQLPartBuilders. The result can store the string SQL as well as all
	// the errors that might have occurred during the building of many parts.
	ISQLPartBuilderResult struct {
		factory SQLPartBuilderFactory
		sb      *strings.Builder
		errs    []error
	}

	// PrefixBuilderOption is a SQLPartBuilderOption for adding a prefix to
	// the build SQL string.
	PrefixBuilderOption struct {
		Prefix string
	}
)

const (
	// Table name part.
	TableName SQLPartType = iota + 1

	// CREATE TABLE statement.
	CreateTable
	// Beginning of the CREATE TABLE statement.
	CreateTableName
	// CREATE TABLE TEMPORARY part.
	CreateTableTemporary
	// CREATE TABLE IF NOT EXISTS part.
	CreateTableNameIfNotExist
	// CREATE TABLE AS part.
	CreateTableAs
	// Column definition of a create table statement.
	ColumnDefn
)

var (
	_ SQLPartBuilderResult = (*ISQLPartBuilderResult)(nil)
	_ SQLPartBuilderOption = (*PrefixBuilderOption)(nil)

	ErrSQLPartBuilder = fmt.Errorf("SQL_PART_BUILDER_ERR")
	ErrUnimplemented  = fmt.Errorf("UNIMPLEMENTED_SQL_PART_BUILDER_ERR")
)

// String returns the string representation of a SQLPartType.
//revive:disable:cyclomatic Need complexity for switch statement
func (partType SQLPartType) String() string {
	switch partType {
	case TableName:
		return "TableName"
	case CreateTable:
		return "CreateTable"
	case CreateTableName:
		return "CreateTableName"
	case CreateTableTemporary:
		return "CreateTableTemporary"
	case CreateTableNameIfNotExist:
		return "CreateTableNameIfNotExist"
	case CreateTableAs:
		return "CreateTableAs"
	case ColumnDefn:
		return "ColumnDefn"
	default:
		return "Unknown"
	}
}

//revive:enable:cyclomatic

func NewResult(factory SQLPartBuilderFactory) *ISQLPartBuilderResult {
	return &ISQLPartBuilderResult{
		factory: factory,
		sb:      &strings.Builder{},
		errs:    []error{},
	}
}

// Print is a convenience method for printing a formatted string to
// the SQL string. If an error occurs while writing it will be added
// to the result.
func (res *ISQLPartBuilderResult) Print(args ...any) int {
	var (
		writeN int
		err    error
	)

	if writeN, err = fmt.Fprint(res.sb, args...); err != nil {
		res.AppendError(err)
	}

	return writeN
}

// Printf is a convenience method for printing a formatted string to
// the SQL string. If an error occurs while writing it will be added
// to the result.
func (res *ISQLPartBuilderResult) Printf(format string, args ...any) int {
	var (
		writeN int
		err    error
	)

	if writeN, err = fmt.Fprintf(res.sb, format, args...); err != nil {
		res.AppendError(err)
	}

	return writeN
}

// PrintPrefix is a convenience method for printing any prefix
// specified in the given opts.
func (res *ISQLPartBuilderResult) PrintPrefix(opts []SQLPartBuilderOption) int {
	for _, opt := range opts {
		if opt, ok := opt.(*PrefixBuilderOption); ok {
			return res.Print(opt.Prefix)
		}
	}

	return 0
}

// PrintWithOptions is a convenience method for printing a formatted
// string to the SQL string.
//
// This method will also handle any pretty printing options passed in.
//
// If an error occurs while writing it will be added
// to the result.
func (res *ISQLPartBuilderResult) PrintWithOptions(
	opts []SQLPartBuilderOption,
	a ...any,
) int {
	writeN := 0

	writeN += res.PrintPrefix(opts)
	writeN += res.Print(a...)

	return writeN
}

// PrintfWithOptions is a convenience method for printing a formatted
// string to the SQL string.
//
// This method will also handle any pretty printing options passed in.
//
// If an error occurs while writing it will be added
// to the result.
func (res *ISQLPartBuilderResult) PrintfWithOptions(
	opts []SQLPartBuilderOption,
	format string,
	a ...any,
) int {
	writeN := 0

	writeN += res.PrintPrefix(opts)
	writeN += res.Printf(format, a...)

	return writeN
}

// AppendError will store the given error so it can be viewed in
// aggregate from the result.
func (res *ISQLPartBuilderResult) AppendError(err error) {
	res.errs = append(res.errs, err)
}

// Build is a convenience method for building the given part.
func (res *ISQLPartBuilderResult) Build(
	partType SQLPartType,
	part SQLPart,
	opts ...SQLPartBuilderOption,
) {
	builder := res.factory.NewPartBuilder(partType, part)

	if builder == nil {
		builder = &NotImplementedBuilder{
			partType: partType,
			part:     part,
		}
	}

	builder.Build(res, opts...)
}

// Result returns the SQL built by many SQLPartBuilders. If there were
// at least one error appended during the building, err will be non
// nil.
func (res *ISQLPartBuilderResult) Result() (string, error) {
	if len(res.errs) == 0 {
		return res.sb.String(), nil
	}

	errBuilder := &strings.Builder{}
	for _, err := range res.errs {
		//nolint:revive//just combining errors
		fmt.Fprintf(errBuilder, "\n  %v", err)
	}

	return res.sb.String(), fmt.Errorf(
		"%w %s",
		ErrSQLPartBuilder,
		errBuilder.String(),
	)
}

func (*PrefixBuilderOption) isSQLPartBuilderOption() {}
