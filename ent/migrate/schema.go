// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DownloadClientsColumns holds the columns for the "download_clients" table.
	DownloadClientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "enable", Type: field.TypeBool},
		{Name: "name", Type: field.TypeString},
		{Name: "implementation", Type: field.TypeString},
		{Name: "settings", Type: field.TypeString},
		{Name: "priority", Type: field.TypeString},
		{Name: "remove_completed_downloads", Type: field.TypeBool},
		{Name: "remove_failed_downloads", Type: field.TypeBool},
		{Name: "tags", Type: field.TypeString},
	}
	// DownloadClientsTable holds the schema information for the "download_clients" table.
	DownloadClientsTable = &schema.Table{
		Name:       "download_clients",
		Columns:    DownloadClientsColumns,
		PrimaryKey: []*schema.Column{DownloadClientsColumns[0]},
	}
	// EpidodesColumns holds the columns for the "epidodes" table.
	EpidodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "series_id", Type: field.TypeInt},
		{Name: "season_number", Type: field.TypeInt},
		{Name: "episode_number", Type: field.TypeInt},
		{Name: "title", Type: field.TypeString},
		{Name: "overview", Type: field.TypeString},
		{Name: "air_date", Type: field.TypeTime},
	}
	// EpidodesTable holds the schema information for the "epidodes" table.
	EpidodesTable = &schema.Table{
		Name:       "epidodes",
		Columns:    EpidodesColumns,
		PrimaryKey: []*schema.Column{EpidodesColumns[0]},
	}
	// IndexersColumns holds the columns for the "indexers" table.
	IndexersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "implementation", Type: field.TypeString},
		{Name: "settings", Type: field.TypeString},
		{Name: "enable_rss", Type: field.TypeBool},
		{Name: "priority", Type: field.TypeInt},
	}
	// IndexersTable holds the schema information for the "indexers" table.
	IndexersTable = &schema.Table{
		Name:       "indexers",
		Columns:    IndexersColumns,
		PrimaryKey: []*schema.Column{IndexersColumns[0]},
	}
	// SeriesColumns holds the columns for the "series" table.
	SeriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tmdb_id", Type: field.TypeInt},
		{Name: "imdb_id", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "original_name", Type: field.TypeString},
		{Name: "overview", Type: field.TypeString},
		{Name: "path", Type: field.TypeString},
	}
	// SeriesTable holds the schema information for the "series" table.
	SeriesTable = &schema.Table{
		Name:       "series",
		Columns:    SeriesColumns,
		PrimaryKey: []*schema.Column{SeriesColumns[0]},
	}
	// SettingsColumns holds the columns for the "settings" table.
	SettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
	}
	// SettingsTable holds the schema information for the "settings" table.
	SettingsTable = &schema.Table{
		Name:       "settings",
		Columns:    SettingsColumns,
		PrimaryKey: []*schema.Column{SettingsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DownloadClientsTable,
		EpidodesTable,
		IndexersTable,
		SeriesTable,
		SettingsTable,
	}
)

func init() {
}
