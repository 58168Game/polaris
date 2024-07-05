// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/predicate"
	"polaris/ent/series"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SeriesUpdate is the builder for updating Series entities.
type SeriesUpdate struct {
	config
	hooks    []Hook
	mutation *SeriesMutation
}

// Where appends a list predicates to the SeriesUpdate builder.
func (su *SeriesUpdate) Where(ps ...predicate.Series) *SeriesUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetTmdbID sets the "tmdb_id" field.
func (su *SeriesUpdate) SetTmdbID(i int) *SeriesUpdate {
	su.mutation.ResetTmdbID()
	su.mutation.SetTmdbID(i)
	return su
}

// SetNillableTmdbID sets the "tmdb_id" field if the given value is not nil.
func (su *SeriesUpdate) SetNillableTmdbID(i *int) *SeriesUpdate {
	if i != nil {
		su.SetTmdbID(*i)
	}
	return su
}

// AddTmdbID adds i to the "tmdb_id" field.
func (su *SeriesUpdate) AddTmdbID(i int) *SeriesUpdate {
	su.mutation.AddTmdbID(i)
	return su
}

// SetImdbID sets the "imdb_id" field.
func (su *SeriesUpdate) SetImdbID(s string) *SeriesUpdate {
	su.mutation.SetImdbID(s)
	return su
}

// SetNillableImdbID sets the "imdb_id" field if the given value is not nil.
func (su *SeriesUpdate) SetNillableImdbID(s *string) *SeriesUpdate {
	if s != nil {
		su.SetImdbID(*s)
	}
	return su
}

// ClearImdbID clears the value of the "imdb_id" field.
func (su *SeriesUpdate) ClearImdbID() *SeriesUpdate {
	su.mutation.ClearImdbID()
	return su
}

// SetTitle sets the "title" field.
func (su *SeriesUpdate) SetTitle(s string) *SeriesUpdate {
	su.mutation.SetTitle(s)
	return su
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (su *SeriesUpdate) SetNillableTitle(s *string) *SeriesUpdate {
	if s != nil {
		su.SetTitle(*s)
	}
	return su
}

// SetOriginalName sets the "original_name" field.
func (su *SeriesUpdate) SetOriginalName(s string) *SeriesUpdate {
	su.mutation.SetOriginalName(s)
	return su
}

// SetNillableOriginalName sets the "original_name" field if the given value is not nil.
func (su *SeriesUpdate) SetNillableOriginalName(s *string) *SeriesUpdate {
	if s != nil {
		su.SetOriginalName(*s)
	}
	return su
}

// SetOverview sets the "overview" field.
func (su *SeriesUpdate) SetOverview(s string) *SeriesUpdate {
	su.mutation.SetOverview(s)
	return su
}

// SetNillableOverview sets the "overview" field if the given value is not nil.
func (su *SeriesUpdate) SetNillableOverview(s *string) *SeriesUpdate {
	if s != nil {
		su.SetOverview(*s)
	}
	return su
}

// SetPath sets the "path" field.
func (su *SeriesUpdate) SetPath(s string) *SeriesUpdate {
	su.mutation.SetPath(s)
	return su
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (su *SeriesUpdate) SetNillablePath(s *string) *SeriesUpdate {
	if s != nil {
		su.SetPath(*s)
	}
	return su
}

// SetPosterPath sets the "poster_path" field.
func (su *SeriesUpdate) SetPosterPath(s string) *SeriesUpdate {
	su.mutation.SetPosterPath(s)
	return su
}

// SetNillablePosterPath sets the "poster_path" field if the given value is not nil.
func (su *SeriesUpdate) SetNillablePosterPath(s *string) *SeriesUpdate {
	if s != nil {
		su.SetPosterPath(*s)
	}
	return su
}

// ClearPosterPath clears the value of the "poster_path" field.
func (su *SeriesUpdate) ClearPosterPath() *SeriesUpdate {
	su.mutation.ClearPosterPath()
	return su
}

// Mutation returns the SeriesMutation object of the builder.
func (su *SeriesUpdate) Mutation() *SeriesMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SeriesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SeriesUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SeriesUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SeriesUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SeriesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(series.Table, series.Columns, sqlgraph.NewFieldSpec(series.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.TmdbID(); ok {
		_spec.SetField(series.FieldTmdbID, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedTmdbID(); ok {
		_spec.AddField(series.FieldTmdbID, field.TypeInt, value)
	}
	if value, ok := su.mutation.ImdbID(); ok {
		_spec.SetField(series.FieldImdbID, field.TypeString, value)
	}
	if su.mutation.ImdbIDCleared() {
		_spec.ClearField(series.FieldImdbID, field.TypeString)
	}
	if value, ok := su.mutation.Title(); ok {
		_spec.SetField(series.FieldTitle, field.TypeString, value)
	}
	if value, ok := su.mutation.OriginalName(); ok {
		_spec.SetField(series.FieldOriginalName, field.TypeString, value)
	}
	if value, ok := su.mutation.Overview(); ok {
		_spec.SetField(series.FieldOverview, field.TypeString, value)
	}
	if value, ok := su.mutation.Path(); ok {
		_spec.SetField(series.FieldPath, field.TypeString, value)
	}
	if value, ok := su.mutation.PosterPath(); ok {
		_spec.SetField(series.FieldPosterPath, field.TypeString, value)
	}
	if su.mutation.PosterPathCleared() {
		_spec.ClearField(series.FieldPosterPath, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{series.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SeriesUpdateOne is the builder for updating a single Series entity.
type SeriesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SeriesMutation
}

// SetTmdbID sets the "tmdb_id" field.
func (suo *SeriesUpdateOne) SetTmdbID(i int) *SeriesUpdateOne {
	suo.mutation.ResetTmdbID()
	suo.mutation.SetTmdbID(i)
	return suo
}

// SetNillableTmdbID sets the "tmdb_id" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillableTmdbID(i *int) *SeriesUpdateOne {
	if i != nil {
		suo.SetTmdbID(*i)
	}
	return suo
}

// AddTmdbID adds i to the "tmdb_id" field.
func (suo *SeriesUpdateOne) AddTmdbID(i int) *SeriesUpdateOne {
	suo.mutation.AddTmdbID(i)
	return suo
}

// SetImdbID sets the "imdb_id" field.
func (suo *SeriesUpdateOne) SetImdbID(s string) *SeriesUpdateOne {
	suo.mutation.SetImdbID(s)
	return suo
}

// SetNillableImdbID sets the "imdb_id" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillableImdbID(s *string) *SeriesUpdateOne {
	if s != nil {
		suo.SetImdbID(*s)
	}
	return suo
}

// ClearImdbID clears the value of the "imdb_id" field.
func (suo *SeriesUpdateOne) ClearImdbID() *SeriesUpdateOne {
	suo.mutation.ClearImdbID()
	return suo
}

// SetTitle sets the "title" field.
func (suo *SeriesUpdateOne) SetTitle(s string) *SeriesUpdateOne {
	suo.mutation.SetTitle(s)
	return suo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillableTitle(s *string) *SeriesUpdateOne {
	if s != nil {
		suo.SetTitle(*s)
	}
	return suo
}

// SetOriginalName sets the "original_name" field.
func (suo *SeriesUpdateOne) SetOriginalName(s string) *SeriesUpdateOne {
	suo.mutation.SetOriginalName(s)
	return suo
}

// SetNillableOriginalName sets the "original_name" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillableOriginalName(s *string) *SeriesUpdateOne {
	if s != nil {
		suo.SetOriginalName(*s)
	}
	return suo
}

// SetOverview sets the "overview" field.
func (suo *SeriesUpdateOne) SetOverview(s string) *SeriesUpdateOne {
	suo.mutation.SetOverview(s)
	return suo
}

// SetNillableOverview sets the "overview" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillableOverview(s *string) *SeriesUpdateOne {
	if s != nil {
		suo.SetOverview(*s)
	}
	return suo
}

// SetPath sets the "path" field.
func (suo *SeriesUpdateOne) SetPath(s string) *SeriesUpdateOne {
	suo.mutation.SetPath(s)
	return suo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillablePath(s *string) *SeriesUpdateOne {
	if s != nil {
		suo.SetPath(*s)
	}
	return suo
}

// SetPosterPath sets the "poster_path" field.
func (suo *SeriesUpdateOne) SetPosterPath(s string) *SeriesUpdateOne {
	suo.mutation.SetPosterPath(s)
	return suo
}

// SetNillablePosterPath sets the "poster_path" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillablePosterPath(s *string) *SeriesUpdateOne {
	if s != nil {
		suo.SetPosterPath(*s)
	}
	return suo
}

// ClearPosterPath clears the value of the "poster_path" field.
func (suo *SeriesUpdateOne) ClearPosterPath() *SeriesUpdateOne {
	suo.mutation.ClearPosterPath()
	return suo
}

// Mutation returns the SeriesMutation object of the builder.
func (suo *SeriesUpdateOne) Mutation() *SeriesMutation {
	return suo.mutation
}

// Where appends a list predicates to the SeriesUpdate builder.
func (suo *SeriesUpdateOne) Where(ps ...predicate.Series) *SeriesUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SeriesUpdateOne) Select(field string, fields ...string) *SeriesUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Series entity.
func (suo *SeriesUpdateOne) Save(ctx context.Context) (*Series, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SeriesUpdateOne) SaveX(ctx context.Context) *Series {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SeriesUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SeriesUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SeriesUpdateOne) sqlSave(ctx context.Context) (_node *Series, err error) {
	_spec := sqlgraph.NewUpdateSpec(series.Table, series.Columns, sqlgraph.NewFieldSpec(series.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Series.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, series.FieldID)
		for _, f := range fields {
			if !series.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != series.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.TmdbID(); ok {
		_spec.SetField(series.FieldTmdbID, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedTmdbID(); ok {
		_spec.AddField(series.FieldTmdbID, field.TypeInt, value)
	}
	if value, ok := suo.mutation.ImdbID(); ok {
		_spec.SetField(series.FieldImdbID, field.TypeString, value)
	}
	if suo.mutation.ImdbIDCleared() {
		_spec.ClearField(series.FieldImdbID, field.TypeString)
	}
	if value, ok := suo.mutation.Title(); ok {
		_spec.SetField(series.FieldTitle, field.TypeString, value)
	}
	if value, ok := suo.mutation.OriginalName(); ok {
		_spec.SetField(series.FieldOriginalName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Overview(); ok {
		_spec.SetField(series.FieldOverview, field.TypeString, value)
	}
	if value, ok := suo.mutation.Path(); ok {
		_spec.SetField(series.FieldPath, field.TypeString, value)
	}
	if value, ok := suo.mutation.PosterPath(); ok {
		_spec.SetField(series.FieldPosterPath, field.TypeString, value)
	}
	if suo.mutation.PosterPathCleared() {
		_spec.ClearField(series.FieldPosterPath, field.TypeString)
	}
	_node = &Series{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{series.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
