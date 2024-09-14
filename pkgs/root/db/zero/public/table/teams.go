//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Teams = newTeamsTable("public", "teams", "")

type teamsTable struct {
	postgres.Table

	// Columns
	ID             postgres.ColumnString
	CreatedAt      postgres.ColumnTimestampz
	UpdatedAt      postgres.ColumnTimestampz
	Name           postgres.ColumnString
	Description    postgres.ColumnString
	Slug           postgres.ColumnString
	OrganizationID postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TeamsTable struct {
	teamsTable

	EXCLUDED teamsTable
}

// AS creates new TeamsTable with assigned alias
func (a TeamsTable) AS(alias string) *TeamsTable {
	return newTeamsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TeamsTable with assigned schema name
func (a TeamsTable) FromSchema(schemaName string) *TeamsTable {
	return newTeamsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TeamsTable with assigned table prefix
func (a TeamsTable) WithPrefix(prefix string) *TeamsTable {
	return newTeamsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TeamsTable with assigned table suffix
func (a TeamsTable) WithSuffix(suffix string) *TeamsTable {
	return newTeamsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTeamsTable(schemaName, tableName, alias string) *TeamsTable {
	return &TeamsTable{
		teamsTable: newTeamsTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newTeamsTableImpl("", "excluded", ""),
	}
}

func newTeamsTableImpl(schemaName, tableName, alias string) teamsTable {
	var (
		IDColumn             = postgres.StringColumn("id")
		CreatedAtColumn      = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn      = postgres.TimestampzColumn("updated_at")
		NameColumn           = postgres.StringColumn("name")
		DescriptionColumn    = postgres.StringColumn("description")
		SlugColumn           = postgres.StringColumn("slug")
		OrganizationIDColumn = postgres.StringColumn("organization_id")
		allColumns           = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, NameColumn, DescriptionColumn, SlugColumn, OrganizationIDColumn}
		mutableColumns       = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, NameColumn, DescriptionColumn, SlugColumn, OrganizationIDColumn}
	)

	return teamsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		Name:           NameColumn,
		Description:    DescriptionColumn,
		Slug:           SlugColumn,
		OrganizationID: OrganizationIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}