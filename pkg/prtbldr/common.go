package prtbldr

type (
	// TableNamePart represents a SQL part for a table
	// [SCHEMA_NAME.]TABLE_NAME.
	TableNamePart struct {
		Name   string // Name of the table
		Schema string // Optional Schema name
	}
)

//nolint:grouper // only one variable
var _ SQLPart = (*TableNamePart)(nil)

func (*TableNamePart) Type() SQLPartType {
	return TableName
}
