package mysql_test

import (
	"testing"

	"github.com/yo3jones/gosql/pkg/prtbldr/mysql"
	"github.com/yo3jones/gosql/pkg/sqlbldr"
)

type (
	createTableTest struct {
		t       *testing.T
		builder sqlbldr.SQLBuilder
		name    string
		expect  string
	}
)

func (tc *createTableTest) run() {
	var (
		gotSQL string
		err    error
		t      = tc.t
	)

	if gotSQL, _, err = tc.builder.Build(); err != nil {
		t.Fatal(err)
	}

	if gotSQL != tc.expect {
		t.Errorf(
			"expected generated sql to be \n%s\n but got \n%s\n",
			tc.expect,
			gotSQL,
		)
	}
}

func TestCreateTable(t *testing.T) {
	t.Parallel()

	tests := []*createTableTest{
		{
			name: "simple",
			builder: sqlbldr.NewCreateTableBuilder(mysql.NewFactory()).
				TableName("foo").
				Col("bar").
				Col("baz"),
			expect: "CREATE TABLE foo (foo, baz)",
		},
		{
			name: "with temporary",
			builder: sqlbldr.NewCreateTableBuilder(mysql.NewFactory()).
				TableName("foo").
				Temporary(),
			expect: "CREATE TEMPORARY TABLE foo",
		},
		{
			name: "with if not exists",
			builder: sqlbldr.NewCreateTableBuilder(mysql.NewFactory()).
				TableName("foo").
				IfNotExists(),
			expect: "CREATE IF NOT EXISTS TABLE foo",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.t = t

			tc.run()
		})
	}
}
