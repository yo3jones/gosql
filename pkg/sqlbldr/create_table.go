package sqlbldr

import (
	"fmt"

	"github.com/yo3jones/gosql/pkg/prtbldr"
)

type (
	// CreateTableSQLBuilder describes methods for building a CREATE TABLE SQL
	// statement.
	CreateTableSQLBuilder interface {
		SQLBuilder

		// Temporary marks the table as TEMPORARY.
		Temporary() CreateTableSQLBuilder

		// IfNotExists ads the condition IF NOT EXISTS to the CREATE TABLE
		// statement.
		IfNotExists() CreateTableSQLBuilder

		// TableName specifies the table name of the CREATE statement.
		TableName(
			tableName string,
			opts ...CreateTableTableNameOption,
		) CreateTableSQLBuilder

		// Col adds a column definition to the CREATE TABLE statement.
		Col(name string) CreateTableSQLBuilder
	}

	// ICreateTableSQLBuilder implements CreateTableSQLBuilder
	//
	// CreateTableSQLBuilder describes methods for building a CREATE TABLE SQL
	// statement.
	ICreateTableSQLBuilder struct {
		factory prtbldr.SQLPartBuilderFactory
		part    *prtbldr.CreateTablePart
	}

	// CreateTableSQLBuilderOption describes options accepted by
	// CreateTableSQLBuilder.TableName().
	CreateTableTableNameOption interface {
		// isCreateTableSQLBuilderOption return true
		isCreateTableSQLBuilderOption()
	}
)

//nolint:grouper//can't group a single variable
var _ CreateTableSQLBuilder = (*ICreateTableSQLBuilder)(nil)

// Temporary marks the table as TEMPORARY.
//nolint:ireturn//need to return interface for builder pattern
func (builder *ICreateTableSQLBuilder) Temporary() CreateTableSQLBuilder {
	builder.part.Temporary = true

	return builder
}

// IfNotExists ads the condition IF NOT EXISTS to the CREATE TABLE
// statement.
//nolint:ireturn//need to return interface for builder pattern
func (builder *ICreateTableSQLBuilder) IfNotExists() CreateTableSQLBuilder {
	builder.part.IfNotExists = true

	return builder
}

// TableName specifies the table name of the CREATE statement.
//nolint:ireturn//need to return interface for builder pattern
func (builder *ICreateTableSQLBuilder) TableName(
	name string,
	opts ...CreateTableTableNameOption,
) CreateTableSQLBuilder {
	tableName := &prtbldr.TableNamePart{
		Name:   name,
		Schema: "",
	}

	for _, opt := range opts {
		opt.isCreateTableSQLBuilderOption()

		if opt, ok := opt.(*SchemaNameOption); ok {
			tableName.Schema = opt.name
		}
	}

	builder.part.TableName = tableName

	return builder
}

// Col adds a column definition to the CREATE TABLE statement.
//nolint:ireturn//need to return interface for builder pattern
func (builder *ICreateTableSQLBuilder) Col(name string) CreateTableSQLBuilder {
	colDefn := prtbldr.NewColumnDefnPart()
	colDefn.Name = name

	builder.part.AddColumnDefn(colDefn)

	return builder
}

//revive:disable:function-result-limit
func (builder *ICreateTableSQLBuilder) Build(
	_ ...SQLBuilderOption,
) (string, []any, error) {
	res := prtbldr.NewResult(builder.factory)

	partBuilder := builder.factory.NewPartBuilder(
		prtbldr.CreateTable,
		builder.part,
	)

	partBuilder.Build(res)

	var (
		sql string
		err error
	)

	if sql, err = res.Result(); err != nil {
		return sql, nil, fmt.Errorf("%w error while building CREATE TABLE", err)
	}

	return sql, nil, nil
}

//revive:enable:function-result-limit

// NewCreateTableBuilder returns a new CreateTableSQLBuilder initialized with
// factory which is used for dialect specific building.
func NewCreateTableBuilder(
	factory prtbldr.SQLPartBuilderFactory,
) *ICreateTableSQLBuilder {
	return &ICreateTableSQLBuilder{
		factory: factory,
		part:    prtbldr.NewCreateTable(),
	}
}
