// Code generated by ent, DO NOT EDIT.

package downloadclients

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the downloadclients type in the database.
	Label = "download_clients"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEnable holds the string denoting the enable field in the database.
	FieldEnable = "enable"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldImplementation holds the string denoting the implementation field in the database.
	FieldImplementation = "implementation"
	// FieldSettings holds the string denoting the settings field in the database.
	FieldSettings = "settings"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldRemoveCompletedDownloads holds the string denoting the remove_completed_downloads field in the database.
	FieldRemoveCompletedDownloads = "remove_completed_downloads"
	// FieldRemoveFailedDownloads holds the string denoting the remove_failed_downloads field in the database.
	FieldRemoveFailedDownloads = "remove_failed_downloads"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// Table holds the table name of the downloadclients in the database.
	Table = "download_clients"
)

// Columns holds all SQL columns for downloadclients fields.
var Columns = []string{
	FieldID,
	FieldEnable,
	FieldName,
	FieldImplementation,
	FieldSettings,
	FieldPriority,
	FieldRemoveCompletedDownloads,
	FieldRemoveFailedDownloads,
	FieldTags,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the DownloadClients queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEnable orders the results by the enable field.
func ByEnable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnable, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByImplementation orders the results by the implementation field.
func ByImplementation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImplementation, opts...).ToFunc()
}

// BySettings orders the results by the settings field.
func BySettings(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSettings, opts...).ToFunc()
}

// ByPriority orders the results by the priority field.
func ByPriority(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPriority, opts...).ToFunc()
}

// ByRemoveCompletedDownloads orders the results by the remove_completed_downloads field.
func ByRemoveCompletedDownloads(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemoveCompletedDownloads, opts...).ToFunc()
}

// ByRemoveFailedDownloads orders the results by the remove_failed_downloads field.
func ByRemoveFailedDownloads(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemoveFailedDownloads, opts...).ToFunc()
}

// ByTags orders the results by the tags field.
func ByTags(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTags, opts...).ToFunc()
}
