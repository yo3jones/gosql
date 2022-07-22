package prtbldr

type (
	// CreateTablePart represents a CREATE TABLE statement.
	CreateTablePart struct {
		TableName   *TableNamePart
		columnDefns []*ColumnDefnPart
		Temporary   bool
		IfNotExists bool
	}

	// ColumnDefnPart represents a column definition of a CREATE TABLE
	// statement.
	ColumnDefnPart struct {
		Name string
	}
)

var (
	_ SQLPart = (*CreateTablePart)(nil)
	_ SQLPart = (*ColumnDefnPart)(nil)
)

// NewCreateTable returns a new initialized CreateTablePart.
func NewCreateTable() *CreateTablePart {
	return &CreateTablePart{
		TableName:   nil,
		Temporary:   false,
		IfNotExists: false,
		columnDefns: []*ColumnDefnPart{},
	}
}

// AddColumnDefn adds a column definition part to the CREATE TABLE statement.
func (part *CreateTablePart) AddColumnDefn(colDefn *ColumnDefnPart) {
	part.columnDefns = append(part.columnDefns, colDefn)
}

func (*CreateTablePart) Type() SQLPartType {
	return CreateTable
}

// NewColumnDefnPart returns a new initialized ColumnDefnPart.
func NewColumnDefnPart() *ColumnDefnPart {
	return &ColumnDefnPart{
		Name: "",
	}
}

func (*ColumnDefnPart) Type() SQLPartType {
	return ColumnDefn
}
