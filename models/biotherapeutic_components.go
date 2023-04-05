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

// BiotherapeuticComponent is an object representing the database table.
type BiotherapeuticComponent struct {
	BiocompID   int64 `boil:"biocomp_id" json:"biocomp_id" toml:"biocomp_id" yaml:"biocomp_id"`
	Molregno    int64 `boil:"molregno" json:"molregno" toml:"molregno" yaml:"molregno"`
	ComponentID int64 `boil:"component_id" json:"component_id" toml:"component_id" yaml:"component_id"`

	R *biotherapeuticComponentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L biotherapeuticComponentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BiotherapeuticComponentColumns = struct {
	BiocompID   string
	Molregno    string
	ComponentID string
}{
	BiocompID:   "biocomp_id",
	Molregno:    "molregno",
	ComponentID: "component_id",
}

var BiotherapeuticComponentTableColumns = struct {
	BiocompID   string
	Molregno    string
	ComponentID string
}{
	BiocompID:   "biotherapeutic_components.biocomp_id",
	Molregno:    "biotherapeutic_components.molregno",
	ComponentID: "biotherapeutic_components.component_id",
}

// Generated where

var BiotherapeuticComponentWhere = struct {
	BiocompID   whereHelperint64
	Molregno    whereHelperint64
	ComponentID whereHelperint64
}{
	BiocompID:   whereHelperint64{field: "\"biotherapeutic_components\".\"biocomp_id\""},
	Molregno:    whereHelperint64{field: "\"biotherapeutic_components\".\"molregno\""},
	ComponentID: whereHelperint64{field: "\"biotherapeutic_components\".\"component_id\""},
}

// BiotherapeuticComponentRels is where relationship names are stored.
var BiotherapeuticComponentRels = struct {
	MolregnoBiotherapeutic string
	Component              string
}{
	MolregnoBiotherapeutic: "MolregnoBiotherapeutic",
	Component:              "Component",
}

// biotherapeuticComponentR is where relationships are stored.
type biotherapeuticComponentR struct {
	MolregnoBiotherapeutic *Biotherapeutic       `boil:"MolregnoBiotherapeutic" json:"MolregnoBiotherapeutic" toml:"MolregnoBiotherapeutic" yaml:"MolregnoBiotherapeutic"`
	Component              *BioComponentSequence `boil:"Component" json:"Component" toml:"Component" yaml:"Component"`
}

// NewStruct creates a new relationship struct
func (*biotherapeuticComponentR) NewStruct() *biotherapeuticComponentR {
	return &biotherapeuticComponentR{}
}

func (r *biotherapeuticComponentR) GetMolregnoBiotherapeutic() *Biotherapeutic {
	if r == nil {
		return nil
	}
	return r.MolregnoBiotherapeutic
}

func (r *biotherapeuticComponentR) GetComponent() *BioComponentSequence {
	if r == nil {
		return nil
	}
	return r.Component
}

// biotherapeuticComponentL is where Load methods for each relationship are stored.
type biotherapeuticComponentL struct{}

var (
	biotherapeuticComponentAllColumns            = []string{"biocomp_id", "molregno", "component_id"}
	biotherapeuticComponentColumnsWithoutDefault = []string{"biocomp_id", "molregno", "component_id"}
	biotherapeuticComponentColumnsWithDefault    = []string{}
	biotherapeuticComponentPrimaryKeyColumns     = []string{"biocomp_id"}
	biotherapeuticComponentGeneratedColumns      = []string{}
)

type (
	// BiotherapeuticComponentSlice is an alias for a slice of pointers to BiotherapeuticComponent.
	// This should almost always be used instead of []BiotherapeuticComponent.
	BiotherapeuticComponentSlice []*BiotherapeuticComponent
	// BiotherapeuticComponentHook is the signature for custom BiotherapeuticComponent hook methods
	BiotherapeuticComponentHook func(context.Context, boil.ContextExecutor, *BiotherapeuticComponent) error

	biotherapeuticComponentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	biotherapeuticComponentType                 = reflect.TypeOf(&BiotherapeuticComponent{})
	biotherapeuticComponentMapping              = queries.MakeStructMapping(biotherapeuticComponentType)
	biotherapeuticComponentPrimaryKeyMapping, _ = queries.BindMapping(biotherapeuticComponentType, biotherapeuticComponentMapping, biotherapeuticComponentPrimaryKeyColumns)
	biotherapeuticComponentInsertCacheMut       sync.RWMutex
	biotherapeuticComponentInsertCache          = make(map[string]insertCache)
	biotherapeuticComponentUpdateCacheMut       sync.RWMutex
	biotherapeuticComponentUpdateCache          = make(map[string]updateCache)
	biotherapeuticComponentUpsertCacheMut       sync.RWMutex
	biotherapeuticComponentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var biotherapeuticComponentAfterSelectHooks []BiotherapeuticComponentHook

var biotherapeuticComponentBeforeInsertHooks []BiotherapeuticComponentHook
var biotherapeuticComponentAfterInsertHooks []BiotherapeuticComponentHook

var biotherapeuticComponentBeforeUpdateHooks []BiotherapeuticComponentHook
var biotherapeuticComponentAfterUpdateHooks []BiotherapeuticComponentHook

var biotherapeuticComponentBeforeDeleteHooks []BiotherapeuticComponentHook
var biotherapeuticComponentAfterDeleteHooks []BiotherapeuticComponentHook

var biotherapeuticComponentBeforeUpsertHooks []BiotherapeuticComponentHook
var biotherapeuticComponentAfterUpsertHooks []BiotherapeuticComponentHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BiotherapeuticComponent) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range biotherapeuticComponentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BiotherapeuticComponent) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range biotherapeuticComponentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BiotherapeuticComponent) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range biotherapeuticComponentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BiotherapeuticComponent) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range biotherapeuticComponentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BiotherapeuticComponent) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range biotherapeuticComponentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BiotherapeuticComponent) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range biotherapeuticComponentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BiotherapeuticComponent) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range biotherapeuticComponentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BiotherapeuticComponent) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range biotherapeuticComponentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BiotherapeuticComponent) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range biotherapeuticComponentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBiotherapeuticComponentHook registers your hook function for all future operations.
func AddBiotherapeuticComponentHook(hookPoint boil.HookPoint, biotherapeuticComponentHook BiotherapeuticComponentHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		biotherapeuticComponentAfterSelectHooks = append(biotherapeuticComponentAfterSelectHooks, biotherapeuticComponentHook)
	case boil.BeforeInsertHook:
		biotherapeuticComponentBeforeInsertHooks = append(biotherapeuticComponentBeforeInsertHooks, biotherapeuticComponentHook)
	case boil.AfterInsertHook:
		biotherapeuticComponentAfterInsertHooks = append(biotherapeuticComponentAfterInsertHooks, biotherapeuticComponentHook)
	case boil.BeforeUpdateHook:
		biotherapeuticComponentBeforeUpdateHooks = append(biotherapeuticComponentBeforeUpdateHooks, biotherapeuticComponentHook)
	case boil.AfterUpdateHook:
		biotherapeuticComponentAfterUpdateHooks = append(biotherapeuticComponentAfterUpdateHooks, biotherapeuticComponentHook)
	case boil.BeforeDeleteHook:
		biotherapeuticComponentBeforeDeleteHooks = append(biotherapeuticComponentBeforeDeleteHooks, biotherapeuticComponentHook)
	case boil.AfterDeleteHook:
		biotherapeuticComponentAfterDeleteHooks = append(biotherapeuticComponentAfterDeleteHooks, biotherapeuticComponentHook)
	case boil.BeforeUpsertHook:
		biotherapeuticComponentBeforeUpsertHooks = append(biotherapeuticComponentBeforeUpsertHooks, biotherapeuticComponentHook)
	case boil.AfterUpsertHook:
		biotherapeuticComponentAfterUpsertHooks = append(biotherapeuticComponentAfterUpsertHooks, biotherapeuticComponentHook)
	}
}

// One returns a single biotherapeuticComponent record from the query.
func (q biotherapeuticComponentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BiotherapeuticComponent, error) {
	o := &BiotherapeuticComponent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for biotherapeutic_components")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BiotherapeuticComponent records from the query.
func (q biotherapeuticComponentQuery) All(ctx context.Context, exec boil.ContextExecutor) (BiotherapeuticComponentSlice, error) {
	var o []*BiotherapeuticComponent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BiotherapeuticComponent slice")
	}

	if len(biotherapeuticComponentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BiotherapeuticComponent records in the query.
func (q biotherapeuticComponentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count biotherapeutic_components rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q biotherapeuticComponentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if biotherapeutic_components exists")
	}

	return count > 0, nil
}

// MolregnoBiotherapeutic pointed to by the foreign key.
func (o *BiotherapeuticComponent) MolregnoBiotherapeutic(mods ...qm.QueryMod) biotherapeuticQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"molregno\" = ?", o.Molregno),
	}

	queryMods = append(queryMods, mods...)

	return Biotherapeutics(queryMods...)
}

// Component pointed to by the foreign key.
func (o *BiotherapeuticComponent) Component(mods ...qm.QueryMod) bioComponentSequenceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"component_id\" = ?", o.ComponentID),
	}

	queryMods = append(queryMods, mods...)

	return BioComponentSequences(queryMods...)
}

// LoadMolregnoBiotherapeutic allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (biotherapeuticComponentL) LoadMolregnoBiotherapeutic(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBiotherapeuticComponent interface{}, mods queries.Applicator) error {
	var slice []*BiotherapeuticComponent
	var object *BiotherapeuticComponent

	if singular {
		object = maybeBiotherapeuticComponent.(*BiotherapeuticComponent)
	} else {
		slice = *maybeBiotherapeuticComponent.(*[]*BiotherapeuticComponent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &biotherapeuticComponentR{}
		}
		args = append(args, object.Molregno)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &biotherapeuticComponentR{}
			}

			for _, a := range args {
				if a == obj.Molregno {
					continue Outer
				}
			}

			args = append(args, obj.Molregno)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`biotherapeutics`),
		qm.WhereIn(`biotherapeutics.molregno in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Biotherapeutic")
	}

	var resultSlice []*Biotherapeutic
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Biotherapeutic")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for biotherapeutics")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for biotherapeutics")
	}

	if len(biotherapeuticComponentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.MolregnoBiotherapeutic = foreign
		if foreign.R == nil {
			foreign.R = &biotherapeuticR{}
		}
		foreign.R.MolregnoBiotherapeuticComponents = append(foreign.R.MolregnoBiotherapeuticComponents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Molregno == foreign.Molregno {
				local.R.MolregnoBiotherapeutic = foreign
				if foreign.R == nil {
					foreign.R = &biotherapeuticR{}
				}
				foreign.R.MolregnoBiotherapeuticComponents = append(foreign.R.MolregnoBiotherapeuticComponents, local)
				break
			}
		}
	}

	return nil
}

// LoadComponent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (biotherapeuticComponentL) LoadComponent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBiotherapeuticComponent interface{}, mods queries.Applicator) error {
	var slice []*BiotherapeuticComponent
	var object *BiotherapeuticComponent

	if singular {
		object = maybeBiotherapeuticComponent.(*BiotherapeuticComponent)
	} else {
		slice = *maybeBiotherapeuticComponent.(*[]*BiotherapeuticComponent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &biotherapeuticComponentR{}
		}
		args = append(args, object.ComponentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &biotherapeuticComponentR{}
			}

			for _, a := range args {
				if a == obj.ComponentID {
					continue Outer
				}
			}

			args = append(args, obj.ComponentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`bio_component_sequences`),
		qm.WhereIn(`bio_component_sequences.component_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load BioComponentSequence")
	}

	var resultSlice []*BioComponentSequence
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice BioComponentSequence")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for bio_component_sequences")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for bio_component_sequences")
	}

	if len(biotherapeuticComponentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Component = foreign
		if foreign.R == nil {
			foreign.R = &bioComponentSequenceR{}
		}
		foreign.R.ComponentBiotherapeuticComponents = append(foreign.R.ComponentBiotherapeuticComponents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ComponentID == foreign.ComponentID {
				local.R.Component = foreign
				if foreign.R == nil {
					foreign.R = &bioComponentSequenceR{}
				}
				foreign.R.ComponentBiotherapeuticComponents = append(foreign.R.ComponentBiotherapeuticComponents, local)
				break
			}
		}
	}

	return nil
}

// SetMolregnoBiotherapeutic of the biotherapeuticComponent to the related item.
// Sets o.R.MolregnoBiotherapeutic to related.
// Adds o to related.R.MolregnoBiotherapeuticComponents.
func (o *BiotherapeuticComponent) SetMolregnoBiotherapeutic(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Biotherapeutic) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"biotherapeutic_components\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"molregno"}),
		strmangle.WhereClause("\"", "\"", 0, biotherapeuticComponentPrimaryKeyColumns),
	)
	values := []interface{}{related.Molregno, o.BiocompID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Molregno = related.Molregno
	if o.R == nil {
		o.R = &biotherapeuticComponentR{
			MolregnoBiotherapeutic: related,
		}
	} else {
		o.R.MolregnoBiotherapeutic = related
	}

	if related.R == nil {
		related.R = &biotherapeuticR{
			MolregnoBiotherapeuticComponents: BiotherapeuticComponentSlice{o},
		}
	} else {
		related.R.MolregnoBiotherapeuticComponents = append(related.R.MolregnoBiotherapeuticComponents, o)
	}

	return nil
}

// SetComponent of the biotherapeuticComponent to the related item.
// Sets o.R.Component to related.
// Adds o to related.R.ComponentBiotherapeuticComponents.
func (o *BiotherapeuticComponent) SetComponent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *BioComponentSequence) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"biotherapeutic_components\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"component_id"}),
		strmangle.WhereClause("\"", "\"", 0, biotherapeuticComponentPrimaryKeyColumns),
	)
	values := []interface{}{related.ComponentID, o.BiocompID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ComponentID = related.ComponentID
	if o.R == nil {
		o.R = &biotherapeuticComponentR{
			Component: related,
		}
	} else {
		o.R.Component = related
	}

	if related.R == nil {
		related.R = &bioComponentSequenceR{
			ComponentBiotherapeuticComponents: BiotherapeuticComponentSlice{o},
		}
	} else {
		related.R.ComponentBiotherapeuticComponents = append(related.R.ComponentBiotherapeuticComponents, o)
	}

	return nil
}

// BiotherapeuticComponents retrieves all the records using an executor.
func BiotherapeuticComponents(mods ...qm.QueryMod) biotherapeuticComponentQuery {
	mods = append(mods, qm.From("\"biotherapeutic_components\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"biotherapeutic_components\".*"})
	}

	return biotherapeuticComponentQuery{q}
}

// FindBiotherapeuticComponent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBiotherapeuticComponent(ctx context.Context, exec boil.ContextExecutor, biocompID int64, selectCols ...string) (*BiotherapeuticComponent, error) {
	biotherapeuticComponentObj := &BiotherapeuticComponent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"biotherapeutic_components\" where \"biocomp_id\"=?", sel,
	)

	q := queries.Raw(query, biocompID)

	err := q.Bind(ctx, exec, biotherapeuticComponentObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from biotherapeutic_components")
	}

	if err = biotherapeuticComponentObj.doAfterSelectHooks(ctx, exec); err != nil {
		return biotherapeuticComponentObj, err
	}

	return biotherapeuticComponentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BiotherapeuticComponent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no biotherapeutic_components provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(biotherapeuticComponentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	biotherapeuticComponentInsertCacheMut.RLock()
	cache, cached := biotherapeuticComponentInsertCache[key]
	biotherapeuticComponentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			biotherapeuticComponentAllColumns,
			biotherapeuticComponentColumnsWithDefault,
			biotherapeuticComponentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(biotherapeuticComponentType, biotherapeuticComponentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(biotherapeuticComponentType, biotherapeuticComponentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"biotherapeutic_components\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"biotherapeutic_components\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into biotherapeutic_components")
	}

	if !cached {
		biotherapeuticComponentInsertCacheMut.Lock()
		biotherapeuticComponentInsertCache[key] = cache
		biotherapeuticComponentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BiotherapeuticComponent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BiotherapeuticComponent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	biotherapeuticComponentUpdateCacheMut.RLock()
	cache, cached := biotherapeuticComponentUpdateCache[key]
	biotherapeuticComponentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			biotherapeuticComponentAllColumns,
			biotherapeuticComponentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update biotherapeutic_components, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"biotherapeutic_components\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, biotherapeuticComponentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(biotherapeuticComponentType, biotherapeuticComponentMapping, append(wl, biotherapeuticComponentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update biotherapeutic_components row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for biotherapeutic_components")
	}

	if !cached {
		biotherapeuticComponentUpdateCacheMut.Lock()
		biotherapeuticComponentUpdateCache[key] = cache
		biotherapeuticComponentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q biotherapeuticComponentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for biotherapeutic_components")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for biotherapeutic_components")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BiotherapeuticComponentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), biotherapeuticComponentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"biotherapeutic_components\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, biotherapeuticComponentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in biotherapeuticComponent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all biotherapeuticComponent")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BiotherapeuticComponent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no biotherapeutic_components provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(biotherapeuticComponentColumnsWithDefault, o)

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

	biotherapeuticComponentUpsertCacheMut.RLock()
	cache, cached := biotherapeuticComponentUpsertCache[key]
	biotherapeuticComponentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			biotherapeuticComponentAllColumns,
			biotherapeuticComponentColumnsWithDefault,
			biotherapeuticComponentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			biotherapeuticComponentAllColumns,
			biotherapeuticComponentPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert biotherapeutic_components, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(biotherapeuticComponentPrimaryKeyColumns))
			copy(conflict, biotherapeuticComponentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"biotherapeutic_components\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(biotherapeuticComponentType, biotherapeuticComponentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(biotherapeuticComponentType, biotherapeuticComponentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert biotherapeutic_components")
	}

	if !cached {
		biotherapeuticComponentUpsertCacheMut.Lock()
		biotherapeuticComponentUpsertCache[key] = cache
		biotherapeuticComponentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BiotherapeuticComponent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BiotherapeuticComponent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BiotherapeuticComponent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), biotherapeuticComponentPrimaryKeyMapping)
	sql := "DELETE FROM \"biotherapeutic_components\" WHERE \"biocomp_id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from biotherapeutic_components")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for biotherapeutic_components")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q biotherapeuticComponentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no biotherapeuticComponentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from biotherapeutic_components")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for biotherapeutic_components")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BiotherapeuticComponentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(biotherapeuticComponentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), biotherapeuticComponentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"biotherapeutic_components\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, biotherapeuticComponentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from biotherapeuticComponent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for biotherapeutic_components")
	}

	if len(biotherapeuticComponentAfterDeleteHooks) != 0 {
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
func (o *BiotherapeuticComponent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBiotherapeuticComponent(ctx, exec, o.BiocompID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BiotherapeuticComponentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BiotherapeuticComponentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), biotherapeuticComponentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"biotherapeutic_components\".* FROM \"biotherapeutic_components\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, biotherapeuticComponentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BiotherapeuticComponentSlice")
	}

	*o = slice

	return nil
}

// BiotherapeuticComponentExists checks if the BiotherapeuticComponent row exists.
func BiotherapeuticComponentExists(ctx context.Context, exec boil.ContextExecutor, biocompID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"biotherapeutic_components\" where \"biocomp_id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, biocompID)
	}
	row := exec.QueryRowContext(ctx, sql, biocompID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if biotherapeutic_components exists")
	}

	return exists, nil
}