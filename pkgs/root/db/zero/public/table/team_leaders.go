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

var TeamLeaders = newTeamLeadersTable("public", "team_leaders", "")

type teamLeadersTable struct {
	postgres.Table

	// Columns
	ID           postgres.ColumnString
	CreatedAt    postgres.ColumnTimestampz
	TeamID       postgres.ColumnString
	TeamMemberID postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TeamLeadersTable struct {
	teamLeadersTable

	EXCLUDED teamLeadersTable
}

// AS creates new TeamLeadersTable with assigned alias
func (a TeamLeadersTable) AS(alias string) *TeamLeadersTable {
	return newTeamLeadersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TeamLeadersTable with assigned schema name
func (a TeamLeadersTable) FromSchema(schemaName string) *TeamLeadersTable {
	return newTeamLeadersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TeamLeadersTable with assigned table prefix
func (a TeamLeadersTable) WithPrefix(prefix string) *TeamLeadersTable {
	return newTeamLeadersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TeamLeadersTable with assigned table suffix
func (a TeamLeadersTable) WithSuffix(suffix string) *TeamLeadersTable {
	return newTeamLeadersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTeamLeadersTable(schemaName, tableName, alias string) *TeamLeadersTable {
	return &TeamLeadersTable{
		teamLeadersTable: newTeamLeadersTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newTeamLeadersTableImpl("", "excluded", ""),
	}
}

func newTeamLeadersTableImpl(schemaName, tableName, alias string) teamLeadersTable {
	var (
		IDColumn           = postgres.StringColumn("id")
		CreatedAtColumn    = postgres.TimestampzColumn("created_at")
		TeamIDColumn       = postgres.StringColumn("team_id")
		TeamMemberIDColumn = postgres.StringColumn("team_member_id")
		allColumns         = postgres.ColumnList{IDColumn, CreatedAtColumn, TeamIDColumn, TeamMemberIDColumn}
		mutableColumns     = postgres.ColumnList{CreatedAtColumn, TeamIDColumn, TeamMemberIDColumn}
	)

	return teamLeadersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		CreatedAt:    CreatedAtColumn,
		TeamID:       TeamIDColumn,
		TeamMemberID: TeamMemberIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
