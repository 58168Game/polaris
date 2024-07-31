// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"polaris/ent/storage"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Storage is the model entity for the Storage schema.
type Storage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Implementation holds the value of the "implementation" field.
	Implementation storage.Implementation `json:"implementation,omitempty"`
	// TvPath holds the value of the "tv_path" field.
	TvPath string `json:"tv_path,omitempty"`
	// MoviePath holds the value of the "movie_path" field.
	MoviePath string `json:"movie_path,omitempty"`
	// Settings holds the value of the "settings" field.
	Settings string `json:"settings,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// Default holds the value of the "default" field.
	Default      bool `json:"default,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Storage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case storage.FieldDeleted, storage.FieldDefault:
			values[i] = new(sql.NullBool)
		case storage.FieldID:
			values[i] = new(sql.NullInt64)
		case storage.FieldName, storage.FieldImplementation, storage.FieldTvPath, storage.FieldMoviePath, storage.FieldSettings:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Storage fields.
func (s *Storage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case storage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case storage.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case storage.FieldImplementation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field implementation", values[i])
			} else if value.Valid {
				s.Implementation = storage.Implementation(value.String)
			}
		case storage.FieldTvPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tv_path", values[i])
			} else if value.Valid {
				s.TvPath = value.String
			}
		case storage.FieldMoviePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field movie_path", values[i])
			} else if value.Valid {
				s.MoviePath = value.String
			}
		case storage.FieldSettings:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field settings", values[i])
			} else if value.Valid {
				s.Settings = value.String
			}
		case storage.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				s.Deleted = value.Bool
			}
		case storage.FieldDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field default", values[i])
			} else if value.Valid {
				s.Default = value.Bool
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Storage.
// This includes values selected through modifiers, order, etc.
func (s *Storage) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Storage.
// Note that you need to call Storage.Unwrap() before calling this method if this Storage
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Storage) Update() *StorageUpdateOne {
	return NewStorageClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Storage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Storage) Unwrap() *Storage {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Storage is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Storage) String() string {
	var builder strings.Builder
	builder.WriteString("Storage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("implementation=")
	builder.WriteString(fmt.Sprintf("%v", s.Implementation))
	builder.WriteString(", ")
	builder.WriteString("tv_path=")
	builder.WriteString(s.TvPath)
	builder.WriteString(", ")
	builder.WriteString("movie_path=")
	builder.WriteString(s.MoviePath)
	builder.WriteString(", ")
	builder.WriteString("settings=")
	builder.WriteString(s.Settings)
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", s.Deleted))
	builder.WriteString(", ")
	builder.WriteString("default=")
	builder.WriteString(fmt.Sprintf("%v", s.Default))
	builder.WriteByte(')')
	return builder.String()
}

// Storages is a parsable slice of Storage.
type Storages []*Storage
