// package sqlbldr provides functionality for building SQL statements
package sqlbldr

import (
	"fmt"
)

type (
	// SQLBuilder describes a Build method will return a string sql statement
	// along with any values passed in while building. The err return value
	// will be non null if an error occurred.
	SQLBuilder interface {
		// Build builds a SQL statement string and returns it. If any values
		// were used during the building, they are returned in values. If
		// any error occurred, err will be non nil.
		Build(opts ...SQLBuilderOption) (sql string, values []any, err error)
	}

	// SQLBuilderOption describes an option that can be passed to the
	// SQLBuilder Build method.
	SQLBuilderOption interface {
		// isSQLBuilderOption is only used to help development to make the
		// SQLBuilderOption type safe.
		isSQLBuilderOption() bool
	}

	SchemaNameOption struct {
		name string
	}
)

var (
	ErrSQLBuilder = fmt.Errorf("SQL_BUILDER_ERROR")

	_ CreateTableTableNameOption = (*SchemaNameOption)(nil)
)

// Schema returns a SchemaNameOption with the given name.
func Schema(name string) *SchemaNameOption {
	return &SchemaNameOption{name}
}

func (*SchemaNameOption) isCreateTableSQLBuilderOption() {}
