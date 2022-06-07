// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ConfidenceScoreLookup is an object representing the database table.
type ConfidenceScoreLookup struct {
	ConfidenceScore int16  `boil:"confidence_score" json:"confidence_score" toml:"confidence_score" yaml:"confidence_score"`
	Description     string `boil:"description" json:"description" toml:"description" yaml:"description"`
	TargetMapping   string `boil:"target_mapping" json:"target_mapping" toml:"target_mapping" yaml:"target_mapping"`

	R *confidenceScoreLookupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L confidenceScoreLookupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ConfidenceScoreLookupColumns = struct {
	ConfidenceScore string
	Description     string
	TargetMapping   string
}{
	ConfidenceScore: "confidence_score",
	Description:     "description",
	TargetMapping:   "target_mapping",
}

var ConfidenceScoreLookupTableColumns = struct {
	ConfidenceScore string
	Description     string
	TargetMapping   string
}{
	ConfidenceScore: "confidence_score_lookup.confidence_score",
	Description:     "confidence_score_lookup.description",
	TargetMapping:   "confidence_score_lookup.target_mapping",
}

// Generated where

var ConfidenceScoreLookupWhere = struct {
	ConfidenceScore whereHelperint16
	Description     whereHelperstring
	TargetMapping   whereHelperstring
}{
	ConfidenceScore: whereHelperint16{field: "\"confidence_score_lookup\".\"confidence_score\""},
	Description:     whereHelperstring{field: "\"confidence_score_lookup\".\"description\""},
	TargetMapping:   whereHelperstring{field: "\"confidence_score_lookup\".\"target_mapping\""},
}

// ConfidenceScoreLookupRels is where relationship names are stored.
var ConfidenceScoreLookupRels = struct {
	ConfidenceScoreAssays string
}{
	ConfidenceScoreAssays: "ConfidenceScoreAssays",
}

// confidenceScoreLookupR is where relationships are stored.
type confidenceScoreLookupR struct {
	ConfidenceScoreAssays AssaySlice `boil:"ConfidenceScoreAssays" json:"ConfidenceScoreAssays" toml:"ConfidenceScoreAssays" yaml:"ConfidenceScoreAssays"`
}

// NewStruct creates a new relationship struct
func (*confidenceScoreLookupR) NewStruct() *confidenceScoreLookupR {
	return &confidenceScoreLookupR{}
}

func (r *confidenceScoreLookupR) GetConfidenceScoreAssays() AssaySlice {
	if r == nil {
		return nil
	}
	return r.ConfidenceScoreAssays
}

// confidenceScoreLookupL is where Load methods for each relationship are stored.
type confidenceScoreLookupL struct{}

var (
	confidenceScoreLookupAllColumns            = []string{"confidence_score", "description", "target_mapping"}
	confidenceScoreLookupColumnsWithoutDefault = []string{"confidence_score", "description", "target_mapping"}
	confidenceScoreLookupColumnsWithDefault    = []string{}
	confidenceScoreLookupPrimaryKeyColumns     = []string{"confidence_score"}
	confidenceScoreLookupGeneratedColumns      = []string{}
)

type (
	// ConfidenceScoreLookupSlice is an alias for a slice of pointers to ConfidenceScoreLookup.
	// This should almost always be used instead of []ConfidenceScoreLookup.
	ConfidenceScoreLookupSlice []*ConfidenceScoreLookup
	// ConfidenceScoreLookupHook is the signature for custom ConfidenceScoreLookup hook methods
	ConfidenceScoreLookupHook func(context.Context, boil.ContextExecutor, *ConfidenceScoreLookup) error

	confidenceScoreLookupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	confidenceScoreLookupType                 = reflect.TypeOf(&ConfidenceScoreLookup{})
	confidenceScoreLookupMapping              = queries.MakeStructMapping(confidenceScoreLookupType)
	confidenceScoreLookupPrimaryKeyMapping, _ = queries.BindMapping(confidenceScoreLookupType, confidenceScoreLookupMapping, confidenceScoreLookupPrimaryKeyColumns)
	confidenceScoreLookupInsertCacheMut       sync.RWMutex
	confidenceScoreLookupInsertCache          = make(map[string]insertCache)
	confidenceScoreLookupUpdateCacheMut       sync.RWMutex
	confidenceScoreLookupUpdateCache          = make(map[string]updateCache)
	confidenceScoreLookupUpsertCacheMut       sync.RWMutex
	confidenceScoreLookupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var confidenceScoreLookupAfterSelectHooks []ConfidenceScoreLookupHook

var confidenceScoreLookupBeforeInsertHooks []ConfidenceScoreLookupHook
var confidenceScoreLookupAfterInsertHooks []ConfidenceScoreLookupHook

var confidenceScoreLookupBeforeUpdateHooks []ConfidenceScoreLookupHook
var confidenceScoreLookupAfterUpdateHooks []ConfidenceScoreLookupHook

var confidenceScoreLookupBeforeDeleteHooks []ConfidenceScoreLookupHook
var confidenceScoreLookupAfterDeleteHooks []ConfidenceScoreLookupHook

var confidenceScoreLookupBeforeUpsertHooks []ConfidenceScoreLookupHook
var confidenceScoreLookupAfterUpsertHooks []ConfidenceScoreLookupHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ConfidenceScoreLookup) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range confidenceScoreLookupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ConfidenceScoreLookup) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range confidenceScoreLookupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ConfidenceScoreLookup) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range confidenceScoreLookupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ConfidenceScoreLookup) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range confidenceScoreLookupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ConfidenceScoreLookup) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range confidenceScoreLookupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ConfidenceScoreLookup) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range confidenceScoreLookupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ConfidenceScoreLookup) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range confidenceScoreLookupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ConfidenceScoreLookup) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range confidenceScoreLookupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ConfidenceScoreLookup) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range confidenceScoreLookupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddConfidenceScoreLookupHook registers your hook function for all future operations.
func AddConfidenceScoreLookupHook(hookPoint boil.HookPoint, confidenceScoreLookupHook ConfidenceScoreLookupHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		confidenceScoreLookupAfterSelectHooks = append(confidenceScoreLookupAfterSelectHooks, confidenceScoreLookupHook)
	case boil.BeforeInsertHook:
		confidenceScoreLookupBeforeInsertHooks = append(confidenceScoreLookupBeforeInsertHooks, confidenceScoreLookupHook)
	case boil.AfterInsertHook:
		confidenceScoreLookupAfterInsertHooks = append(confidenceScoreLookupAfterInsertHooks, confidenceScoreLookupHook)
	case boil.BeforeUpdateHook:
		confidenceScoreLookupBeforeUpdateHooks = append(confidenceScoreLookupBeforeUpdateHooks, confidenceScoreLookupHook)
	case boil.AfterUpdateHook:
		confidenceScoreLookupAfterUpdateHooks = append(confidenceScoreLookupAfterUpdateHooks, confidenceScoreLookupHook)
	case boil.BeforeDeleteHook:
		confidenceScoreLookupBeforeDeleteHooks = append(confidenceScoreLookupBeforeDeleteHooks, confidenceScoreLookupHook)
	case boil.AfterDeleteHook:
		confidenceScoreLookupAfterDeleteHooks = append(confidenceScoreLookupAfterDeleteHooks, confidenceScoreLookupHook)
	case boil.BeforeUpsertHook:
		confidenceScoreLookupBeforeUpsertHooks = append(confidenceScoreLookupBeforeUpsertHooks, confidenceScoreLookupHook)
	case boil.AfterUpsertHook:
		confidenceScoreLookupAfterUpsertHooks = append(confidenceScoreLookupAfterUpsertHooks, confidenceScoreLookupHook)
	}
}

// One returns a single confidenceScoreLookup record from the query.
func (q confidenceScoreLookupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ConfidenceScoreLookup, error) {
	o := &ConfidenceScoreLookup{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for confidence_score_lookup")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ConfidenceScoreLookup records from the query.
func (q confidenceScoreLookupQuery) All(ctx context.Context, exec boil.ContextExecutor) (ConfidenceScoreLookupSlice, error) {
	var o []*ConfidenceScoreLookup

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ConfidenceScoreLookup slice")
	}

	if len(confidenceScoreLookupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ConfidenceScoreLookup records in the query.
func (q confidenceScoreLookupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count confidence_score_lookup rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q confidenceScoreLookupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if confidence_score_lookup exists")
	}

	return count > 0, nil
}

// ConfidenceScoreAssays retrieves all the assay's Assays with an executor via confidence_score column.
func (o *ConfidenceScoreLookup) ConfidenceScoreAssays(mods ...qm.QueryMod) assayQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"assays\".\"confidence_score\"=?", o.ConfidenceScore),
	)

	return Assays(queryMods...)
}

// LoadConfidenceScoreAssays allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (confidenceScoreLookupL) LoadConfidenceScoreAssays(ctx context.Context, e boil.ContextExecutor, singular bool, maybeConfidenceScoreLookup interface{}, mods queries.Applicator) error {
	var slice []*ConfidenceScoreLookup
	var object *ConfidenceScoreLookup

	if singular {
		object = maybeConfidenceScoreLookup.(*ConfidenceScoreLookup)
	} else {
		slice = *maybeConfidenceScoreLookup.(*[]*ConfidenceScoreLookup)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &confidenceScoreLookupR{}
		}
		args = append(args, object.ConfidenceScore)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &confidenceScoreLookupR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ConfidenceScore) {
					continue Outer
				}
			}

			args = append(args, obj.ConfidenceScore)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`assays`),
		qm.WhereIn(`assays.confidence_score in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load assays")
	}

	var resultSlice []*Assay
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice assays")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on assays")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for assays")
	}

	if len(assayAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ConfidenceScoreAssays = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &assayR{}
			}
			foreign.R.ConfidenceScoreConfidenceScoreLookup = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ConfidenceScore, foreign.ConfidenceScore) {
				local.R.ConfidenceScoreAssays = append(local.R.ConfidenceScoreAssays, foreign)
				if foreign.R == nil {
					foreign.R = &assayR{}
				}
				foreign.R.ConfidenceScoreConfidenceScoreLookup = local
				break
			}
		}
	}

	return nil
}

// AddConfidenceScoreAssays adds the given related objects to the existing relationships
// of the confidence_score_lookup, optionally inserting them as new records.
// Appends related to o.R.ConfidenceScoreAssays.
// Sets related.R.ConfidenceScoreConfidenceScoreLookup appropriately.
func (o *ConfidenceScoreLookup) AddConfidenceScoreAssays(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Assay) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ConfidenceScore, o.ConfidenceScore)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"assays\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"confidence_score"}),
				strmangle.WhereClause("\"", "\"", 0, assayPrimaryKeyColumns),
			)
			values := []interface{}{o.ConfidenceScore, rel.AssayID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.ConfidenceScore, o.ConfidenceScore)
		}
	}

	if o.R == nil {
		o.R = &confidenceScoreLookupR{
			ConfidenceScoreAssays: related,
		}
	} else {
		o.R.ConfidenceScoreAssays = append(o.R.ConfidenceScoreAssays, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &assayR{
				ConfidenceScoreConfidenceScoreLookup: o,
			}
		} else {
			rel.R.ConfidenceScoreConfidenceScoreLookup = o
		}
	}
	return nil
}

// SetConfidenceScoreAssays removes all previously related items of the
// confidence_score_lookup replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.ConfidenceScoreConfidenceScoreLookup's ConfidenceScoreAssays accordingly.
// Replaces o.R.ConfidenceScoreAssays with related.
// Sets related.R.ConfidenceScoreConfidenceScoreLookup's ConfidenceScoreAssays accordingly.
func (o *ConfidenceScoreLookup) SetConfidenceScoreAssays(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Assay) error {
	query := "update \"assays\" set \"confidence_score\" = null where \"confidence_score\" = ?"
	values := []interface{}{o.ConfidenceScore}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.ConfidenceScoreAssays {
			queries.SetScanner(&rel.ConfidenceScore, nil)
			if rel.R == nil {
				continue
			}

			rel.R.ConfidenceScoreConfidenceScoreLookup = nil
		}
		o.R.ConfidenceScoreAssays = nil
	}

	return o.AddConfidenceScoreAssays(ctx, exec, insert, related...)
}

// RemoveConfidenceScoreAssays relationships from objects passed in.
// Removes related items from R.ConfidenceScoreAssays (uses pointer comparison, removal does not keep order)
// Sets related.R.ConfidenceScoreConfidenceScoreLookup.
func (o *ConfidenceScoreLookup) RemoveConfidenceScoreAssays(ctx context.Context, exec boil.ContextExecutor, related ...*Assay) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ConfidenceScore, nil)
		if rel.R != nil {
			rel.R.ConfidenceScoreConfidenceScoreLookup = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("confidence_score")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ConfidenceScoreAssays {
			if rel != ri {
				continue
			}

			ln := len(o.R.ConfidenceScoreAssays)
			if ln > 1 && i < ln-1 {
				o.R.ConfidenceScoreAssays[i] = o.R.ConfidenceScoreAssays[ln-1]
			}
			o.R.ConfidenceScoreAssays = o.R.ConfidenceScoreAssays[:ln-1]
			break
		}
	}

	return nil
}

// ConfidenceScoreLookups retrieves all the records using an executor.
func ConfidenceScoreLookups(mods ...qm.QueryMod) confidenceScoreLookupQuery {
	mods = append(mods, qm.From("\"confidence_score_lookup\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"confidence_score_lookup\".*"})
	}

	return confidenceScoreLookupQuery{q}
}

// FindConfidenceScoreLookup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindConfidenceScoreLookup(ctx context.Context, exec boil.ContextExecutor, confidenceScore int16, selectCols ...string) (*ConfidenceScoreLookup, error) {
	confidenceScoreLookupObj := &ConfidenceScoreLookup{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"confidence_score_lookup\" where \"confidence_score\"=?", sel,
	)

	q := queries.Raw(query, confidenceScore)

	err := q.Bind(ctx, exec, confidenceScoreLookupObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from confidence_score_lookup")
	}

	if err = confidenceScoreLookupObj.doAfterSelectHooks(ctx, exec); err != nil {
		return confidenceScoreLookupObj, err
	}

	return confidenceScoreLookupObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ConfidenceScoreLookup) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no confidence_score_lookup provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(confidenceScoreLookupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	confidenceScoreLookupInsertCacheMut.RLock()
	cache, cached := confidenceScoreLookupInsertCache[key]
	confidenceScoreLookupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			confidenceScoreLookupAllColumns,
			confidenceScoreLookupColumnsWithDefault,
			confidenceScoreLookupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(confidenceScoreLookupType, confidenceScoreLookupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(confidenceScoreLookupType, confidenceScoreLookupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"confidence_score_lookup\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"confidence_score_lookup\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into confidence_score_lookup")
	}

	if !cached {
		confidenceScoreLookupInsertCacheMut.Lock()
		confidenceScoreLookupInsertCache[key] = cache
		confidenceScoreLookupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ConfidenceScoreLookup.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ConfidenceScoreLookup) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	confidenceScoreLookupUpdateCacheMut.RLock()
	cache, cached := confidenceScoreLookupUpdateCache[key]
	confidenceScoreLookupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			confidenceScoreLookupAllColumns,
			confidenceScoreLookupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update confidence_score_lookup, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"confidence_score_lookup\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, confidenceScoreLookupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(confidenceScoreLookupType, confidenceScoreLookupMapping, append(wl, confidenceScoreLookupPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update confidence_score_lookup row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for confidence_score_lookup")
	}

	if !cached {
		confidenceScoreLookupUpdateCacheMut.Lock()
		confidenceScoreLookupUpdateCache[key] = cache
		confidenceScoreLookupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q confidenceScoreLookupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for confidence_score_lookup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for confidence_score_lookup")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ConfidenceScoreLookupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), confidenceScoreLookupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"confidence_score_lookup\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, confidenceScoreLookupPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in confidenceScoreLookup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all confidenceScoreLookup")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ConfidenceScoreLookup) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no confidence_score_lookup provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(confidenceScoreLookupColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	confidenceScoreLookupUpsertCacheMut.RLock()
	cache, cached := confidenceScoreLookupUpsertCache[key]
	confidenceScoreLookupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			confidenceScoreLookupAllColumns,
			confidenceScoreLookupColumnsWithDefault,
			confidenceScoreLookupColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			confidenceScoreLookupAllColumns,
			confidenceScoreLookupPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert confidence_score_lookup, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(confidenceScoreLookupPrimaryKeyColumns))
			copy(conflict, confidenceScoreLookupPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"confidence_score_lookup\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(confidenceScoreLookupType, confidenceScoreLookupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(confidenceScoreLookupType, confidenceScoreLookupMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert confidence_score_lookup")
	}

	if !cached {
		confidenceScoreLookupUpsertCacheMut.Lock()
		confidenceScoreLookupUpsertCache[key] = cache
		confidenceScoreLookupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ConfidenceScoreLookup record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ConfidenceScoreLookup) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ConfidenceScoreLookup provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), confidenceScoreLookupPrimaryKeyMapping)
	sql := "DELETE FROM \"confidence_score_lookup\" WHERE \"confidence_score\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from confidence_score_lookup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for confidence_score_lookup")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q confidenceScoreLookupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no confidenceScoreLookupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from confidence_score_lookup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for confidence_score_lookup")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ConfidenceScoreLookupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(confidenceScoreLookupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), confidenceScoreLookupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"confidence_score_lookup\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, confidenceScoreLookupPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from confidenceScoreLookup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for confidence_score_lookup")
	}

	if len(confidenceScoreLookupAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ConfidenceScoreLookup) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindConfidenceScoreLookup(ctx, exec, o.ConfidenceScore)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ConfidenceScoreLookupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ConfidenceScoreLookupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), confidenceScoreLookupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"confidence_score_lookup\".* FROM \"confidence_score_lookup\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, confidenceScoreLookupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ConfidenceScoreLookupSlice")
	}

	*o = slice

	return nil
}

// ConfidenceScoreLookupExists checks if the ConfidenceScoreLookup row exists.
func ConfidenceScoreLookupExists(ctx context.Context, exec boil.ContextExecutor, confidenceScore int16) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"confidence_score_lookup\" where \"confidence_score\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, confidenceScore)
	}
	row := exec.QueryRowContext(ctx, sql, confidenceScore)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if confidence_score_lookup exists")
	}

	return exists, nil
}