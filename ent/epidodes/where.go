// Code generated by ent, DO NOT EDIT.

package epidodes

import (
	"polaris/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLTE(FieldID, id))
}

// SeriesID applies equality check predicate on the "series_id" field. It's identical to SeriesIDEQ.
func SeriesID(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldSeriesID, v))
}

// SeasonNumber applies equality check predicate on the "season_number" field. It's identical to SeasonNumberEQ.
func SeasonNumber(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldSeasonNumber, v))
}

// EpisodeNumber applies equality check predicate on the "episode_number" field. It's identical to EpisodeNumberEQ.
func EpisodeNumber(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldEpisodeNumber, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldTitle, v))
}

// Overview applies equality check predicate on the "overview" field. It's identical to OverviewEQ.
func Overview(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldOverview, v))
}

// AirDate applies equality check predicate on the "air_date" field. It's identical to AirDateEQ.
func AirDate(v time.Time) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldAirDate, v))
}

// SeriesIDEQ applies the EQ predicate on the "series_id" field.
func SeriesIDEQ(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldSeriesID, v))
}

// SeriesIDNEQ applies the NEQ predicate on the "series_id" field.
func SeriesIDNEQ(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNEQ(FieldSeriesID, v))
}

// SeriesIDIn applies the In predicate on the "series_id" field.
func SeriesIDIn(vs ...int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldIn(FieldSeriesID, vs...))
}

// SeriesIDNotIn applies the NotIn predicate on the "series_id" field.
func SeriesIDNotIn(vs ...int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNotIn(FieldSeriesID, vs...))
}

// SeriesIDGT applies the GT predicate on the "series_id" field.
func SeriesIDGT(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGT(FieldSeriesID, v))
}

// SeriesIDGTE applies the GTE predicate on the "series_id" field.
func SeriesIDGTE(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGTE(FieldSeriesID, v))
}

// SeriesIDLT applies the LT predicate on the "series_id" field.
func SeriesIDLT(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLT(FieldSeriesID, v))
}

// SeriesIDLTE applies the LTE predicate on the "series_id" field.
func SeriesIDLTE(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLTE(FieldSeriesID, v))
}

// SeasonNumberEQ applies the EQ predicate on the "season_number" field.
func SeasonNumberEQ(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldSeasonNumber, v))
}

// SeasonNumberNEQ applies the NEQ predicate on the "season_number" field.
func SeasonNumberNEQ(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNEQ(FieldSeasonNumber, v))
}

// SeasonNumberIn applies the In predicate on the "season_number" field.
func SeasonNumberIn(vs ...int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldIn(FieldSeasonNumber, vs...))
}

// SeasonNumberNotIn applies the NotIn predicate on the "season_number" field.
func SeasonNumberNotIn(vs ...int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNotIn(FieldSeasonNumber, vs...))
}

// SeasonNumberGT applies the GT predicate on the "season_number" field.
func SeasonNumberGT(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGT(FieldSeasonNumber, v))
}

// SeasonNumberGTE applies the GTE predicate on the "season_number" field.
func SeasonNumberGTE(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGTE(FieldSeasonNumber, v))
}

// SeasonNumberLT applies the LT predicate on the "season_number" field.
func SeasonNumberLT(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLT(FieldSeasonNumber, v))
}

// SeasonNumberLTE applies the LTE predicate on the "season_number" field.
func SeasonNumberLTE(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLTE(FieldSeasonNumber, v))
}

// EpisodeNumberEQ applies the EQ predicate on the "episode_number" field.
func EpisodeNumberEQ(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldEpisodeNumber, v))
}

// EpisodeNumberNEQ applies the NEQ predicate on the "episode_number" field.
func EpisodeNumberNEQ(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNEQ(FieldEpisodeNumber, v))
}

// EpisodeNumberIn applies the In predicate on the "episode_number" field.
func EpisodeNumberIn(vs ...int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldIn(FieldEpisodeNumber, vs...))
}

// EpisodeNumberNotIn applies the NotIn predicate on the "episode_number" field.
func EpisodeNumberNotIn(vs ...int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNotIn(FieldEpisodeNumber, vs...))
}

// EpisodeNumberGT applies the GT predicate on the "episode_number" field.
func EpisodeNumberGT(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGT(FieldEpisodeNumber, v))
}

// EpisodeNumberGTE applies the GTE predicate on the "episode_number" field.
func EpisodeNumberGTE(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGTE(FieldEpisodeNumber, v))
}

// EpisodeNumberLT applies the LT predicate on the "episode_number" field.
func EpisodeNumberLT(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLT(FieldEpisodeNumber, v))
}

// EpisodeNumberLTE applies the LTE predicate on the "episode_number" field.
func EpisodeNumberLTE(v int) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLTE(FieldEpisodeNumber, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldContainsFold(FieldTitle, v))
}

// OverviewEQ applies the EQ predicate on the "overview" field.
func OverviewEQ(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldOverview, v))
}

// OverviewNEQ applies the NEQ predicate on the "overview" field.
func OverviewNEQ(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNEQ(FieldOverview, v))
}

// OverviewIn applies the In predicate on the "overview" field.
func OverviewIn(vs ...string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldIn(FieldOverview, vs...))
}

// OverviewNotIn applies the NotIn predicate on the "overview" field.
func OverviewNotIn(vs ...string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNotIn(FieldOverview, vs...))
}

// OverviewGT applies the GT predicate on the "overview" field.
func OverviewGT(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGT(FieldOverview, v))
}

// OverviewGTE applies the GTE predicate on the "overview" field.
func OverviewGTE(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGTE(FieldOverview, v))
}

// OverviewLT applies the LT predicate on the "overview" field.
func OverviewLT(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLT(FieldOverview, v))
}

// OverviewLTE applies the LTE predicate on the "overview" field.
func OverviewLTE(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLTE(FieldOverview, v))
}

// OverviewContains applies the Contains predicate on the "overview" field.
func OverviewContains(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldContains(FieldOverview, v))
}

// OverviewHasPrefix applies the HasPrefix predicate on the "overview" field.
func OverviewHasPrefix(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldHasPrefix(FieldOverview, v))
}

// OverviewHasSuffix applies the HasSuffix predicate on the "overview" field.
func OverviewHasSuffix(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldHasSuffix(FieldOverview, v))
}

// OverviewEqualFold applies the EqualFold predicate on the "overview" field.
func OverviewEqualFold(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEqualFold(FieldOverview, v))
}

// OverviewContainsFold applies the ContainsFold predicate on the "overview" field.
func OverviewContainsFold(v string) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldContainsFold(FieldOverview, v))
}

// AirDateEQ applies the EQ predicate on the "air_date" field.
func AirDateEQ(v time.Time) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldEQ(FieldAirDate, v))
}

// AirDateNEQ applies the NEQ predicate on the "air_date" field.
func AirDateNEQ(v time.Time) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNEQ(FieldAirDate, v))
}

// AirDateIn applies the In predicate on the "air_date" field.
func AirDateIn(vs ...time.Time) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldIn(FieldAirDate, vs...))
}

// AirDateNotIn applies the NotIn predicate on the "air_date" field.
func AirDateNotIn(vs ...time.Time) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldNotIn(FieldAirDate, vs...))
}

// AirDateGT applies the GT predicate on the "air_date" field.
func AirDateGT(v time.Time) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGT(FieldAirDate, v))
}

// AirDateGTE applies the GTE predicate on the "air_date" field.
func AirDateGTE(v time.Time) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldGTE(FieldAirDate, v))
}

// AirDateLT applies the LT predicate on the "air_date" field.
func AirDateLT(v time.Time) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLT(FieldAirDate, v))
}

// AirDateLTE applies the LTE predicate on the "air_date" field.
func AirDateLTE(v time.Time) predicate.Epidodes {
	return predicate.Epidodes(sql.FieldLTE(FieldAirDate, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Epidodes) predicate.Epidodes {
	return predicate.Epidodes(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Epidodes) predicate.Epidodes {
	return predicate.Epidodes(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Epidodes) predicate.Epidodes {
	return predicate.Epidodes(sql.NotPredicates(p))
}
