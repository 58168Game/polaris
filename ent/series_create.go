// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/episode"
	"polaris/ent/series"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SeriesCreate is the builder for creating a Series entity.
type SeriesCreate struct {
	config
	mutation *SeriesMutation
	hooks    []Hook
}

// SetTmdbID sets the "tmdb_id" field.
func (sc *SeriesCreate) SetTmdbID(i int) *SeriesCreate {
	sc.mutation.SetTmdbID(i)
	return sc
}

// SetImdbID sets the "imdb_id" field.
func (sc *SeriesCreate) SetImdbID(s string) *SeriesCreate {
	sc.mutation.SetImdbID(s)
	return sc
}

// SetNillableImdbID sets the "imdb_id" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableImdbID(s *string) *SeriesCreate {
	if s != nil {
		sc.SetImdbID(*s)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SeriesCreate) SetName(s string) *SeriesCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetOriginalName sets the "original_name" field.
func (sc *SeriesCreate) SetOriginalName(s string) *SeriesCreate {
	sc.mutation.SetOriginalName(s)
	return sc
}

// SetOverview sets the "overview" field.
func (sc *SeriesCreate) SetOverview(s string) *SeriesCreate {
	sc.mutation.SetOverview(s)
	return sc
}

// SetPath sets the "path" field.
func (sc *SeriesCreate) SetPath(s string) *SeriesCreate {
	sc.mutation.SetPath(s)
	return sc
}

// SetPosterPath sets the "poster_path" field.
func (sc *SeriesCreate) SetPosterPath(s string) *SeriesCreate {
	sc.mutation.SetPosterPath(s)
	return sc
}

// SetNillablePosterPath sets the "poster_path" field if the given value is not nil.
func (sc *SeriesCreate) SetNillablePosterPath(s *string) *SeriesCreate {
	if s != nil {
		sc.SetPosterPath(*s)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SeriesCreate) SetCreatedAt(t time.Time) *SeriesCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableCreatedAt(t *time.Time) *SeriesCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetAirDate sets the "air_date" field.
func (sc *SeriesCreate) SetAirDate(s string) *SeriesCreate {
	sc.mutation.SetAirDate(s)
	return sc
}

// SetNillableAirDate sets the "air_date" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableAirDate(s *string) *SeriesCreate {
	if s != nil {
		sc.SetAirDate(*s)
	}
	return sc
}

// AddEpisodeIDs adds the "episodes" edge to the Episode entity by IDs.
func (sc *SeriesCreate) AddEpisodeIDs(ids ...int) *SeriesCreate {
	sc.mutation.AddEpisodeIDs(ids...)
	return sc
}

// AddEpisodes adds the "episodes" edges to the Episode entity.
func (sc *SeriesCreate) AddEpisodes(e ...*Episode) *SeriesCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return sc.AddEpisodeIDs(ids...)
}

// Mutation returns the SeriesMutation object of the builder.
func (sc *SeriesCreate) Mutation() *SeriesMutation {
	return sc.mutation
}

// Save creates the Series in the database.
func (sc *SeriesCreate) Save(ctx context.Context) (*Series, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SeriesCreate) SaveX(ctx context.Context) *Series {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SeriesCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SeriesCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SeriesCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := series.DefaultCreatedAt
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.AirDate(); !ok {
		v := series.DefaultAirDate
		sc.mutation.SetAirDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SeriesCreate) check() error {
	if _, ok := sc.mutation.TmdbID(); !ok {
		return &ValidationError{Name: "tmdb_id", err: errors.New(`ent: missing required field "Series.tmdb_id"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Series.name"`)}
	}
	if _, ok := sc.mutation.OriginalName(); !ok {
		return &ValidationError{Name: "original_name", err: errors.New(`ent: missing required field "Series.original_name"`)}
	}
	if _, ok := sc.mutation.Overview(); !ok {
		return &ValidationError{Name: "overview", err: errors.New(`ent: missing required field "Series.overview"`)}
	}
	if _, ok := sc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Series.path"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Series.created_at"`)}
	}
	if _, ok := sc.mutation.AirDate(); !ok {
		return &ValidationError{Name: "air_date", err: errors.New(`ent: missing required field "Series.air_date"`)}
	}
	return nil
}

func (sc *SeriesCreate) sqlSave(ctx context.Context) (*Series, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SeriesCreate) createSpec() (*Series, *sqlgraph.CreateSpec) {
	var (
		_node = &Series{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(series.Table, sqlgraph.NewFieldSpec(series.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.TmdbID(); ok {
		_spec.SetField(series.FieldTmdbID, field.TypeInt, value)
		_node.TmdbID = value
	}
	if value, ok := sc.mutation.ImdbID(); ok {
		_spec.SetField(series.FieldImdbID, field.TypeString, value)
		_node.ImdbID = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(series.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.OriginalName(); ok {
		_spec.SetField(series.FieldOriginalName, field.TypeString, value)
		_node.OriginalName = value
	}
	if value, ok := sc.mutation.Overview(); ok {
		_spec.SetField(series.FieldOverview, field.TypeString, value)
		_node.Overview = value
	}
	if value, ok := sc.mutation.Path(); ok {
		_spec.SetField(series.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := sc.mutation.PosterPath(); ok {
		_spec.SetField(series.FieldPosterPath, field.TypeString, value)
		_node.PosterPath = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(series.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.AirDate(); ok {
		_spec.SetField(series.FieldAirDate, field.TypeString, value)
		_node.AirDate = value
	}
	if nodes := sc.mutation.EpisodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   series.EpisodesTable,
			Columns: []string{series.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SeriesCreateBulk is the builder for creating many Series entities in bulk.
type SeriesCreateBulk struct {
	config
	err      error
	builders []*SeriesCreate
}

// Save creates the Series entities in the database.
func (scb *SeriesCreateBulk) Save(ctx context.Context) ([]*Series, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Series, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SeriesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SeriesCreateBulk) SaveX(ctx context.Context) []*Series {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SeriesCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SeriesCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
