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

// CurationLookup is an object representing the database table.
type CurationLookup struct {
	CuratedBy   string `boil:"curated_by" json:"curated_by" toml:"curated_by" yaml:"curated_by"`
	Description string `boil:"description" json:"description" toml:"description" yaml:"description"`

	R *curationLookupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L curationLookupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CurationLookupColumns = struct {
	CuratedBy   string
	Description string
}{
	CuratedBy:   "curated_by",
	Description: "description",
}

var CurationLookupTableColumns = struct {
	CuratedBy   string
	Description string
}{
	CuratedBy:   "curation_lookup.curated_by",
	Description: "curation_lookup.description",
}

// Generated where

var CurationLookupWhere = struct {
	CuratedBy   whereHelperstring
	Description whereHelperstring
}{
	CuratedBy:   whereHelperstring{field: "\"curation_lookup\".\"curated_by\""},
	Description: whereHelperstring{field: "\"curation_lookup\".\"description\""},
}

// CurationLookupRels is where relationship names are stored.
var CurationLookupRels = struct {
	CuratedByAssays string
}{
	CuratedByAssays: "CuratedByAssays",
}

// curationLookupR is where relationships are stored.
type curationLookupR struct {
	CuratedByAssays AssaySlice `boil:"CuratedByAssays" json:"CuratedByAssays" toml:"CuratedByAssays" yaml:"CuratedByAssays"`
}

// NewStruct creates a new relationship struct
func (*curationLookupR) NewStruct() *curationLookupR {
	return &curationLookupR{}
}

func (r *curationLookupR) GetCuratedByAssays() AssaySlice {
	if r == nil {
		return nil
	}
	return r.CuratedByAssays
}

// curationLookupL is where Load methods for each relationship are stored.
type curationLookupL struct{}

var (
	curationLookupAllColumns            = []string{"curated_by", "description"}
	curationLookupColumnsWithoutDefault = []string{"curated_by", "description"}
	curationLookupColumnsWithDefault    = []string{}
	curationLookupPrimaryKeyColumns     = []string{"curated_by"}
	curationLookupGeneratedColumns      = []string{}
)

type (
	// CurationLookupSlice is an alias for a slice of pointers to CurationLookup.
	// This should almost always be used instead of []CurationLookup.
	CurationLookupSlice []*CurationLookup
	// CurationLookupHook is the signature for custom CurationLookup hook methods
	CurationLookupHook func(context.Context, boil.ContextExecutor, *CurationLookup) error

	curationLookupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	curationLookupType                 = reflect.TypeOf(&CurationLookup{})
	curationLookupMapping              = queries.MakeStructMapping(curationLookupType)
	curationLookupPrimaryKeyMapping, _ = queries.BindMapping(curationLookupType, curationLookupMapping, curationLookupPrimaryKeyColumns)
	curationLookupInsertCacheMut       sync.RWMutex
	curationLookupInsertCache          = make(map[string]insertCache)
	curationLookupUpdateCacheMut       sync.RWMutex
	curationLookupUpdateCache          = make(map[string]updateCache)
	curationLookupUpsertCacheMut       sync.RWMutex
	curationLookupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var curationLookupAfterSelectHooks []CurationLookupHook

var curationLookupBeforeInsertHooks []CurationLookupHook
var curationLookupAfterInsertHooks []CurationLookupHook

var curationLookupBeforeUpdateHooks []CurationLookupHook
var curationLookupAfterUpdateHooks []CurationLookupHook

var curationLookupBeforeDeleteHooks []CurationLookupHook
var curationLookupAfterDeleteHooks []CurationLookupHook

var curationLookupBeforeUpsertHooks []CurationLookupHook
var curationLookupAfterUpsertHooks []CurationLookupHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CurationLookup) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range curationLookupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CurationLookup) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range curationLookupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CurationLookup) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range curationLookupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CurationLookup) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range curationLookupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CurationLookup) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range curationLookupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CurationLookup) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range curationLookupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CurationLookup) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range curationLookupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CurationLookup) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range curationLookupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CurationLookup) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range curationLookupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCurationLookupHook registers your hook function for all future operations.
func AddCurationLookupHook(hookPoint boil.HookPoint, curationLookupHook CurationLookupHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		curationLookupAfterSelectHooks = append(curationLookupAfterSelectHooks, curationLookupHook)
	case boil.BeforeInsertHook:
		curationLookupBeforeInsertHooks = append(curationLookupBeforeInsertHooks, curationLookupHook)
	case boil.AfterInsertHook:
		curationLookupAfterInsertHooks = append(curationLookupAfterInsertHooks, curationLookupHook)
	case boil.BeforeUpdateHook:
		curationLookupBeforeUpdateHooks = append(curationLookupBeforeUpdateHooks, curationLookupHook)
	case boil.AfterUpdateHook:
		curationLookupAfterUpdateHooks = append(curationLookupAfterUpdateHooks, curationLookupHook)
	case boil.BeforeDeleteHook:
		curationLookupBeforeDeleteHooks = append(curationLookupBeforeDeleteHooks, curationLookupHook)
	case boil.AfterDeleteHook:
		curationLookupAfterDeleteHooks = append(curationLookupAfterDeleteHooks, curationLookupHook)
	case boil.BeforeUpsertHook:
		curationLookupBeforeUpsertHooks = append(curationLookupBeforeUpsertHooks, curationLookupHook)
	case boil.AfterUpsertHook:
		curationLookupAfterUpsertHooks = append(curationLookupAfterUpsertHooks, curationLookupHook)
	}
}

// One returns a single curationLookup record from the query.
func (q curationLookupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CurationLookup, error) {
	o := &CurationLookup{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for curation_lookup")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CurationLookup records from the query.
func (q curationLookupQuery) All(ctx context.Context, exec boil.ContextExecutor) (CurationLookupSlice, error) {
	var o []*CurationLookup

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CurationLookup slice")
	}

	if len(curationLookupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CurationLookup records in the query.
func (q curationLookupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count curation_lookup rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q curationLookupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if curation_lookup exists")
	}

	return count > 0, nil
}

// CuratedByAssays retrieves all the assay's Assays with an executor via curated_by column.
func (o *CurationLookup) CuratedByAssays(mods ...qm.QueryMod) assayQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"assays\".\"curated_by\"=?", o.CuratedBy),
	)

	return Assays(queryMods...)
}

// LoadCuratedByAssays allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (curationLookupL) LoadCuratedByAssays(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCurationLookup interface{}, mods queries.Applicator) error {
	var slice []*CurationLookup
	var object *CurationLookup

	if singular {
		object = maybeCurationLookup.(*CurationLookup)
	} else {
		slice = *maybeCurationLookup.(*[]*CurationLookup)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &curationLookupR{}
		}
		args = append(args, object.CuratedBy)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &curationLookupR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.CuratedBy) {
					continue Outer
				}
			}

			args = append(args, obj.CuratedBy)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`assays`),
		qm.WhereIn(`assays.curated_by in ?`, args...),
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
		object.R.CuratedByAssays = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &assayR{}
			}
			foreign.R.CuratedByCurationLookup = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.CuratedBy, foreign.CuratedBy) {
				local.R.CuratedByAssays = append(local.R.CuratedByAssays, foreign)
				if foreign.R == nil {
					foreign.R = &assayR{}
				}
				foreign.R.CuratedByCurationLookup = local
				break
			}
		}
	}

	return nil
}

// AddCuratedByAssays adds the given related objects to the existing relationships
// of the curation_lookup, optionally inserting them as new records.
// Appends related to o.R.CuratedByAssays.
// Sets related.R.CuratedByCurationLookup appropriately.
func (o *CurationLookup) AddCuratedByAssays(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Assay) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.CuratedBy, o.CuratedBy)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"assays\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"curated_by"}),
				strmangle.WhereClause("\"", "\"", 0, assayPrimaryKeyColumns),
			)
			values := []interface{}{o.CuratedBy, rel.AssayID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.CuratedBy, o.CuratedBy)
		}
	}

	if o.R == nil {
		o.R = &curationLookupR{
			CuratedByAssays: related,
		}
	} else {
		o.R.CuratedByAssays = append(o.R.CuratedByAssays, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &assayR{
				CuratedByCurationLookup: o,
			}
		} else {
			rel.R.CuratedByCurationLookup = o
		}
	}
	return nil
}

// SetCuratedByAssays removes all previously related items of the
// curation_lookup replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.CuratedByCurationLookup's CuratedByAssays accordingly.
// Replaces o.R.CuratedByAssays with related.
// Sets related.R.CuratedByCurationLookup's CuratedByAssays accordingly.
func (o *CurationLookup) SetCuratedByAssays(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Assay) error {
	query := "update \"assays\" set \"curated_by\" = null where \"curated_by\" = ?"
	values := []interface{}{o.CuratedBy}
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
		for _, rel := range o.R.CuratedByAssays {
			queries.SetScanner(&rel.CuratedBy, nil)
			if rel.R == nil {
				continue
			}

			rel.R.CuratedByCurationLookup = nil
		}
		o.R.CuratedByAssays = nil
	}

	return o.AddCuratedByAssays(ctx, exec, insert, related...)
}

// RemoveCuratedByAssays relationships from objects passed in.
// Removes related items from R.CuratedByAssays (uses pointer comparison, removal does not keep order)
// Sets related.R.CuratedByCurationLookup.
func (o *CurationLookup) RemoveCuratedByAssays(ctx context.Context, exec boil.ContextExecutor, related ...*Assay) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.CuratedBy, nil)
		if rel.R != nil {
			rel.R.CuratedByCurationLookup = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("curated_by")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.CuratedByAssays {
			if rel != ri {
				continue
			}

			ln := len(o.R.CuratedByAssays)
			if ln > 1 && i < ln-1 {
				o.R.CuratedByAssays[i] = o.R.CuratedByAssays[ln-1]
			}
			o.R.CuratedByAssays = o.R.CuratedByAssays[:ln-1]
			break
		}
	}

	return nil
}

// CurationLookups retrieves all the records using an executor.
func CurationLookups(mods ...qm.QueryMod) curationLookupQuery {
	mods = append(mods, qm.From("\"curation_lookup\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"curation_lookup\".*"})
	}

	return curationLookupQuery{q}
}

// FindCurationLookup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCurationLookup(ctx context.Context, exec boil.ContextExecutor, curatedBy string, selectCols ...string) (*CurationLookup, error) {
	curationLookupObj := &CurationLookup{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"curation_lookup\" where \"curated_by\"=?", sel,
	)

	q := queries.Raw(query, curatedBy)

	err := q.Bind(ctx, exec, curationLookupObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from curation_lookup")
	}

	if err = curationLookupObj.doAfterSelectHooks(ctx, exec); err != nil {
		return curationLookupObj, err
	}

	return curationLookupObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CurationLookup) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no curation_lookup provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(curationLookupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	curationLookupInsertCacheMut.RLock()
	cache, cached := curationLookupInsertCache[key]
	curationLookupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			curationLookupAllColumns,
			curationLookupColumnsWithDefault,
			curationLookupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(curationLookupType, curationLookupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(curationLookupType, curationLookupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"curation_lookup\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"curation_lookup\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into curation_lookup")
	}

	if !cached {
		curationLookupInsertCacheMut.Lock()
		curationLookupInsertCache[key] = cache
		curationLookupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CurationLookup.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CurationLookup) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	curationLookupUpdateCacheMut.RLock()
	cache, cached := curationLookupUpdateCache[key]
	curationLookupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			curationLookupAllColumns,
			curationLookupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update curation_lookup, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"curation_lookup\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, curationLookupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(curationLookupType, curationLookupMapping, append(wl, curationLookupPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update curation_lookup row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for curation_lookup")
	}

	if !cached {
		curationLookupUpdateCacheMut.Lock()
		curationLookupUpdateCache[key] = cache
		curationLookupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q curationLookupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for curation_lookup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for curation_lookup")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CurationLookupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), curationLookupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"curation_lookup\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, curationLookupPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in curationLookup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all curationLookup")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CurationLookup) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no curation_lookup provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(curationLookupColumnsWithDefault, o)

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

	curationLookupUpsertCacheMut.RLock()
	cache, cached := curationLookupUpsertCache[key]
	curationLookupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			curationLookupAllColumns,
			curationLookupColumnsWithDefault,
			curationLookupColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			curationLookupAllColumns,
			curationLookupPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert curation_lookup, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(curationLookupPrimaryKeyColumns))
			copy(conflict, curationLookupPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"curation_lookup\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(curationLookupType, curationLookupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(curationLookupType, curationLookupMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert curation_lookup")
	}

	if !cached {
		curationLookupUpsertCacheMut.Lock()
		curationLookupUpsertCache[key] = cache
		curationLookupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CurationLookup record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CurationLookup) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CurationLookup provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), curationLookupPrimaryKeyMapping)
	sql := "DELETE FROM \"curation_lookup\" WHERE \"curated_by\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from curation_lookup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for curation_lookup")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q curationLookupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no curationLookupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from curation_lookup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for curation_lookup")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CurationLookupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(curationLookupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), curationLookupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"curation_lookup\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, curationLookupPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from curationLookup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for curation_lookup")
	}

	if len(curationLookupAfterDeleteHooks) != 0 {
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
func (o *CurationLookup) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCurationLookup(ctx, exec, o.CuratedBy)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CurationLookupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CurationLookupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), curationLookupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"curation_lookup\".* FROM \"curation_lookup\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, curationLookupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CurationLookupSlice")
	}

	*o = slice

	return nil
}

// CurationLookupExists checks if the CurationLookup row exists.
func CurationLookupExists(ctx context.Context, exec boil.ContextExecutor, curatedBy string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"curation_lookup\" where \"curated_by\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, curatedBy)
	}
	row := exec.QueryRowContext(ctx, sql, curatedBy)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if curation_lookup exists")
	}

	return exists, nil
}