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

// MoleculeHracClassification is an object representing the database table.
type MoleculeHracClassification struct {
	MolHracID   int64 `boil:"mol_hrac_id" json:"mol_hrac_id" toml:"mol_hrac_id" yaml:"mol_hrac_id"`
	HracClassID int64 `boil:"hrac_class_id" json:"hrac_class_id" toml:"hrac_class_id" yaml:"hrac_class_id"`
	Molregno    int64 `boil:"molregno" json:"molregno" toml:"molregno" yaml:"molregno"`

	R *moleculeHracClassificationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L moleculeHracClassificationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MoleculeHracClassificationColumns = struct {
	MolHracID   string
	HracClassID string
	Molregno    string
}{
	MolHracID:   "mol_hrac_id",
	HracClassID: "hrac_class_id",
	Molregno:    "molregno",
}

var MoleculeHracClassificationTableColumns = struct {
	MolHracID   string
	HracClassID string
	Molregno    string
}{
	MolHracID:   "molecule_hrac_classification.mol_hrac_id",
	HracClassID: "molecule_hrac_classification.hrac_class_id",
	Molregno:    "molecule_hrac_classification.molregno",
}

// Generated where

var MoleculeHracClassificationWhere = struct {
	MolHracID   whereHelperint64
	HracClassID whereHelperint64
	Molregno    whereHelperint64
}{
	MolHracID:   whereHelperint64{field: "\"molecule_hrac_classification\".\"mol_hrac_id\""},
	HracClassID: whereHelperint64{field: "\"molecule_hrac_classification\".\"hrac_class_id\""},
	Molregno:    whereHelperint64{field: "\"molecule_hrac_classification\".\"molregno\""},
}

// MoleculeHracClassificationRels is where relationship names are stored.
var MoleculeHracClassificationRels = struct {
	MolregnoMoleculeDictionary string
	HracClass                  string
}{
	MolregnoMoleculeDictionary: "MolregnoMoleculeDictionary",
	HracClass:                  "HracClass",
}

// moleculeHracClassificationR is where relationships are stored.
type moleculeHracClassificationR struct {
	MolregnoMoleculeDictionary *MoleculeDictionary `boil:"MolregnoMoleculeDictionary" json:"MolregnoMoleculeDictionary" toml:"MolregnoMoleculeDictionary" yaml:"MolregnoMoleculeDictionary"`
	HracClass                  *HracClassification `boil:"HracClass" json:"HracClass" toml:"HracClass" yaml:"HracClass"`
}

// NewStruct creates a new relationship struct
func (*moleculeHracClassificationR) NewStruct() *moleculeHracClassificationR {
	return &moleculeHracClassificationR{}
}

func (r *moleculeHracClassificationR) GetMolregnoMoleculeDictionary() *MoleculeDictionary {
	if r == nil {
		return nil
	}
	return r.MolregnoMoleculeDictionary
}

func (r *moleculeHracClassificationR) GetHracClass() *HracClassification {
	if r == nil {
		return nil
	}
	return r.HracClass
}

// moleculeHracClassificationL is where Load methods for each relationship are stored.
type moleculeHracClassificationL struct{}

var (
	moleculeHracClassificationAllColumns            = []string{"mol_hrac_id", "hrac_class_id", "molregno"}
	moleculeHracClassificationColumnsWithoutDefault = []string{"mol_hrac_id", "hrac_class_id", "molregno"}
	moleculeHracClassificationColumnsWithDefault    = []string{}
	moleculeHracClassificationPrimaryKeyColumns     = []string{"mol_hrac_id"}
	moleculeHracClassificationGeneratedColumns      = []string{}
)

type (
	// MoleculeHracClassificationSlice is an alias for a slice of pointers to MoleculeHracClassification.
	// This should almost always be used instead of []MoleculeHracClassification.
	MoleculeHracClassificationSlice []*MoleculeHracClassification
	// MoleculeHracClassificationHook is the signature for custom MoleculeHracClassification hook methods
	MoleculeHracClassificationHook func(context.Context, boil.ContextExecutor, *MoleculeHracClassification) error

	moleculeHracClassificationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	moleculeHracClassificationType                 = reflect.TypeOf(&MoleculeHracClassification{})
	moleculeHracClassificationMapping              = queries.MakeStructMapping(moleculeHracClassificationType)
	moleculeHracClassificationPrimaryKeyMapping, _ = queries.BindMapping(moleculeHracClassificationType, moleculeHracClassificationMapping, moleculeHracClassificationPrimaryKeyColumns)
	moleculeHracClassificationInsertCacheMut       sync.RWMutex
	moleculeHracClassificationInsertCache          = make(map[string]insertCache)
	moleculeHracClassificationUpdateCacheMut       sync.RWMutex
	moleculeHracClassificationUpdateCache          = make(map[string]updateCache)
	moleculeHracClassificationUpsertCacheMut       sync.RWMutex
	moleculeHracClassificationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var moleculeHracClassificationAfterSelectHooks []MoleculeHracClassificationHook

var moleculeHracClassificationBeforeInsertHooks []MoleculeHracClassificationHook
var moleculeHracClassificationAfterInsertHooks []MoleculeHracClassificationHook

var moleculeHracClassificationBeforeUpdateHooks []MoleculeHracClassificationHook
var moleculeHracClassificationAfterUpdateHooks []MoleculeHracClassificationHook

var moleculeHracClassificationBeforeDeleteHooks []MoleculeHracClassificationHook
var moleculeHracClassificationAfterDeleteHooks []MoleculeHracClassificationHook

var moleculeHracClassificationBeforeUpsertHooks []MoleculeHracClassificationHook
var moleculeHracClassificationAfterUpsertHooks []MoleculeHracClassificationHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MoleculeHracClassification) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moleculeHracClassificationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MoleculeHracClassification) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moleculeHracClassificationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MoleculeHracClassification) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moleculeHracClassificationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MoleculeHracClassification) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moleculeHracClassificationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MoleculeHracClassification) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moleculeHracClassificationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MoleculeHracClassification) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moleculeHracClassificationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MoleculeHracClassification) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moleculeHracClassificationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MoleculeHracClassification) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moleculeHracClassificationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MoleculeHracClassification) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range moleculeHracClassificationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMoleculeHracClassificationHook registers your hook function for all future operations.
func AddMoleculeHracClassificationHook(hookPoint boil.HookPoint, moleculeHracClassificationHook MoleculeHracClassificationHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		moleculeHracClassificationAfterSelectHooks = append(moleculeHracClassificationAfterSelectHooks, moleculeHracClassificationHook)
	case boil.BeforeInsertHook:
		moleculeHracClassificationBeforeInsertHooks = append(moleculeHracClassificationBeforeInsertHooks, moleculeHracClassificationHook)
	case boil.AfterInsertHook:
		moleculeHracClassificationAfterInsertHooks = append(moleculeHracClassificationAfterInsertHooks, moleculeHracClassificationHook)
	case boil.BeforeUpdateHook:
		moleculeHracClassificationBeforeUpdateHooks = append(moleculeHracClassificationBeforeUpdateHooks, moleculeHracClassificationHook)
	case boil.AfterUpdateHook:
		moleculeHracClassificationAfterUpdateHooks = append(moleculeHracClassificationAfterUpdateHooks, moleculeHracClassificationHook)
	case boil.BeforeDeleteHook:
		moleculeHracClassificationBeforeDeleteHooks = append(moleculeHracClassificationBeforeDeleteHooks, moleculeHracClassificationHook)
	case boil.AfterDeleteHook:
		moleculeHracClassificationAfterDeleteHooks = append(moleculeHracClassificationAfterDeleteHooks, moleculeHracClassificationHook)
	case boil.BeforeUpsertHook:
		moleculeHracClassificationBeforeUpsertHooks = append(moleculeHracClassificationBeforeUpsertHooks, moleculeHracClassificationHook)
	case boil.AfterUpsertHook:
		moleculeHracClassificationAfterUpsertHooks = append(moleculeHracClassificationAfterUpsertHooks, moleculeHracClassificationHook)
	}
}

// One returns a single moleculeHracClassification record from the query.
func (q moleculeHracClassificationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MoleculeHracClassification, error) {
	o := &MoleculeHracClassification{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for molecule_hrac_classification")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MoleculeHracClassification records from the query.
func (q moleculeHracClassificationQuery) All(ctx context.Context, exec boil.ContextExecutor) (MoleculeHracClassificationSlice, error) {
	var o []*MoleculeHracClassification

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MoleculeHracClassification slice")
	}

	if len(moleculeHracClassificationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MoleculeHracClassification records in the query.
func (q moleculeHracClassificationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count molecule_hrac_classification rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q moleculeHracClassificationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if molecule_hrac_classification exists")
	}

	return count > 0, nil
}

// MolregnoMoleculeDictionary pointed to by the foreign key.
func (o *MoleculeHracClassification) MolregnoMoleculeDictionary(mods ...qm.QueryMod) moleculeDictionaryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"molregno\" = ?", o.Molregno),
	}

	queryMods = append(queryMods, mods...)

	return MoleculeDictionaries(queryMods...)
}

// HracClass pointed to by the foreign key.
func (o *MoleculeHracClassification) HracClass(mods ...qm.QueryMod) hracClassificationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"hrac_class_id\" = ?", o.HracClassID),
	}

	queryMods = append(queryMods, mods...)

	return HracClassifications(queryMods...)
}

// LoadMolregnoMoleculeDictionary allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (moleculeHracClassificationL) LoadMolregnoMoleculeDictionary(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMoleculeHracClassification interface{}, mods queries.Applicator) error {
	var slice []*MoleculeHracClassification
	var object *MoleculeHracClassification

	if singular {
		object = maybeMoleculeHracClassification.(*MoleculeHracClassification)
	} else {
		slice = *maybeMoleculeHracClassification.(*[]*MoleculeHracClassification)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &moleculeHracClassificationR{}
		}
		args = append(args, object.Molregno)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &moleculeHracClassificationR{}
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
		qm.From(`molecule_dictionary`),
		qm.WhereIn(`molecule_dictionary.molregno in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load MoleculeDictionary")
	}

	var resultSlice []*MoleculeDictionary
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice MoleculeDictionary")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for molecule_dictionary")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for molecule_dictionary")
	}

	if len(moleculeHracClassificationAfterSelectHooks) != 0 {
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
		object.R.MolregnoMoleculeDictionary = foreign
		if foreign.R == nil {
			foreign.R = &moleculeDictionaryR{}
		}
		foreign.R.MolregnoMoleculeHracClassifications = append(foreign.R.MolregnoMoleculeHracClassifications, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Molregno == foreign.Molregno {
				local.R.MolregnoMoleculeDictionary = foreign
				if foreign.R == nil {
					foreign.R = &moleculeDictionaryR{}
				}
				foreign.R.MolregnoMoleculeHracClassifications = append(foreign.R.MolregnoMoleculeHracClassifications, local)
				break
			}
		}
	}

	return nil
}

// LoadHracClass allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (moleculeHracClassificationL) LoadHracClass(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMoleculeHracClassification interface{}, mods queries.Applicator) error {
	var slice []*MoleculeHracClassification
	var object *MoleculeHracClassification

	if singular {
		object = maybeMoleculeHracClassification.(*MoleculeHracClassification)
	} else {
		slice = *maybeMoleculeHracClassification.(*[]*MoleculeHracClassification)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &moleculeHracClassificationR{}
		}
		args = append(args, object.HracClassID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &moleculeHracClassificationR{}
			}

			for _, a := range args {
				if a == obj.HracClassID {
					continue Outer
				}
			}

			args = append(args, obj.HracClassID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`hrac_classification`),
		qm.WhereIn(`hrac_classification.hrac_class_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load HracClassification")
	}

	var resultSlice []*HracClassification
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice HracClassification")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for hrac_classification")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for hrac_classification")
	}

	if len(moleculeHracClassificationAfterSelectHooks) != 0 {
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
		object.R.HracClass = foreign
		if foreign.R == nil {
			foreign.R = &hracClassificationR{}
		}
		foreign.R.HracClassMoleculeHracClassifications = append(foreign.R.HracClassMoleculeHracClassifications, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.HracClassID == foreign.HracClassID {
				local.R.HracClass = foreign
				if foreign.R == nil {
					foreign.R = &hracClassificationR{}
				}
				foreign.R.HracClassMoleculeHracClassifications = append(foreign.R.HracClassMoleculeHracClassifications, local)
				break
			}
		}
	}

	return nil
}

// SetMolregnoMoleculeDictionary of the moleculeHracClassification to the related item.
// Sets o.R.MolregnoMoleculeDictionary to related.
// Adds o to related.R.MolregnoMoleculeHracClassifications.
func (o *MoleculeHracClassification) SetMolregnoMoleculeDictionary(ctx context.Context, exec boil.ContextExecutor, insert bool, related *MoleculeDictionary) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"molecule_hrac_classification\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"molregno"}),
		strmangle.WhereClause("\"", "\"", 0, moleculeHracClassificationPrimaryKeyColumns),
	)
	values := []interface{}{related.Molregno, o.MolHracID}

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
		o.R = &moleculeHracClassificationR{
			MolregnoMoleculeDictionary: related,
		}
	} else {
		o.R.MolregnoMoleculeDictionary = related
	}

	if related.R == nil {
		related.R = &moleculeDictionaryR{
			MolregnoMoleculeHracClassifications: MoleculeHracClassificationSlice{o},
		}
	} else {
		related.R.MolregnoMoleculeHracClassifications = append(related.R.MolregnoMoleculeHracClassifications, o)
	}

	return nil
}

// SetHracClass of the moleculeHracClassification to the related item.
// Sets o.R.HracClass to related.
// Adds o to related.R.HracClassMoleculeHracClassifications.
func (o *MoleculeHracClassification) SetHracClass(ctx context.Context, exec boil.ContextExecutor, insert bool, related *HracClassification) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"molecule_hrac_classification\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"hrac_class_id"}),
		strmangle.WhereClause("\"", "\"", 0, moleculeHracClassificationPrimaryKeyColumns),
	)
	values := []interface{}{related.HracClassID, o.MolHracID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.HracClassID = related.HracClassID
	if o.R == nil {
		o.R = &moleculeHracClassificationR{
			HracClass: related,
		}
	} else {
		o.R.HracClass = related
	}

	if related.R == nil {
		related.R = &hracClassificationR{
			HracClassMoleculeHracClassifications: MoleculeHracClassificationSlice{o},
		}
	} else {
		related.R.HracClassMoleculeHracClassifications = append(related.R.HracClassMoleculeHracClassifications, o)
	}

	return nil
}

// MoleculeHracClassifications retrieves all the records using an executor.
func MoleculeHracClassifications(mods ...qm.QueryMod) moleculeHracClassificationQuery {
	mods = append(mods, qm.From("\"molecule_hrac_classification\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"molecule_hrac_classification\".*"})
	}

	return moleculeHracClassificationQuery{q}
}

// FindMoleculeHracClassification retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMoleculeHracClassification(ctx context.Context, exec boil.ContextExecutor, molHracID int64, selectCols ...string) (*MoleculeHracClassification, error) {
	moleculeHracClassificationObj := &MoleculeHracClassification{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"molecule_hrac_classification\" where \"mol_hrac_id\"=?", sel,
	)

	q := queries.Raw(query, molHracID)

	err := q.Bind(ctx, exec, moleculeHracClassificationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from molecule_hrac_classification")
	}

	if err = moleculeHracClassificationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return moleculeHracClassificationObj, err
	}

	return moleculeHracClassificationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MoleculeHracClassification) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no molecule_hrac_classification provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(moleculeHracClassificationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	moleculeHracClassificationInsertCacheMut.RLock()
	cache, cached := moleculeHracClassificationInsertCache[key]
	moleculeHracClassificationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			moleculeHracClassificationAllColumns,
			moleculeHracClassificationColumnsWithDefault,
			moleculeHracClassificationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(moleculeHracClassificationType, moleculeHracClassificationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(moleculeHracClassificationType, moleculeHracClassificationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"molecule_hrac_classification\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"molecule_hrac_classification\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into molecule_hrac_classification")
	}

	if !cached {
		moleculeHracClassificationInsertCacheMut.Lock()
		moleculeHracClassificationInsertCache[key] = cache
		moleculeHracClassificationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MoleculeHracClassification.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MoleculeHracClassification) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	moleculeHracClassificationUpdateCacheMut.RLock()
	cache, cached := moleculeHracClassificationUpdateCache[key]
	moleculeHracClassificationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			moleculeHracClassificationAllColumns,
			moleculeHracClassificationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update molecule_hrac_classification, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"molecule_hrac_classification\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, moleculeHracClassificationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(moleculeHracClassificationType, moleculeHracClassificationMapping, append(wl, moleculeHracClassificationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update molecule_hrac_classification row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for molecule_hrac_classification")
	}

	if !cached {
		moleculeHracClassificationUpdateCacheMut.Lock()
		moleculeHracClassificationUpdateCache[key] = cache
		moleculeHracClassificationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q moleculeHracClassificationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for molecule_hrac_classification")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for molecule_hrac_classification")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MoleculeHracClassificationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), moleculeHracClassificationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"molecule_hrac_classification\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, moleculeHracClassificationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in moleculeHracClassification slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all moleculeHracClassification")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MoleculeHracClassification) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no molecule_hrac_classification provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(moleculeHracClassificationColumnsWithDefault, o)

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

	moleculeHracClassificationUpsertCacheMut.RLock()
	cache, cached := moleculeHracClassificationUpsertCache[key]
	moleculeHracClassificationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			moleculeHracClassificationAllColumns,
			moleculeHracClassificationColumnsWithDefault,
			moleculeHracClassificationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			moleculeHracClassificationAllColumns,
			moleculeHracClassificationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert molecule_hrac_classification, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(moleculeHracClassificationPrimaryKeyColumns))
			copy(conflict, moleculeHracClassificationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"molecule_hrac_classification\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(moleculeHracClassificationType, moleculeHracClassificationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(moleculeHracClassificationType, moleculeHracClassificationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert molecule_hrac_classification")
	}

	if !cached {
		moleculeHracClassificationUpsertCacheMut.Lock()
		moleculeHracClassificationUpsertCache[key] = cache
		moleculeHracClassificationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MoleculeHracClassification record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MoleculeHracClassification) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MoleculeHracClassification provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), moleculeHracClassificationPrimaryKeyMapping)
	sql := "DELETE FROM \"molecule_hrac_classification\" WHERE \"mol_hrac_id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from molecule_hrac_classification")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for molecule_hrac_classification")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q moleculeHracClassificationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no moleculeHracClassificationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from molecule_hrac_classification")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for molecule_hrac_classification")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MoleculeHracClassificationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(moleculeHracClassificationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), moleculeHracClassificationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"molecule_hrac_classification\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, moleculeHracClassificationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from moleculeHracClassification slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for molecule_hrac_classification")
	}

	if len(moleculeHracClassificationAfterDeleteHooks) != 0 {
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
func (o *MoleculeHracClassification) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMoleculeHracClassification(ctx, exec, o.MolHracID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MoleculeHracClassificationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MoleculeHracClassificationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), moleculeHracClassificationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"molecule_hrac_classification\".* FROM \"molecule_hrac_classification\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, moleculeHracClassificationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MoleculeHracClassificationSlice")
	}

	*o = slice

	return nil
}

// MoleculeHracClassificationExists checks if the MoleculeHracClassification row exists.
func MoleculeHracClassificationExists(ctx context.Context, exec boil.ContextExecutor, molHracID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"molecule_hrac_classification\" where \"mol_hrac_id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, molHracID)
	}
	row := exec.QueryRowContext(ctx, sql, molHracID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if molecule_hrac_classification exists")
	}

	return exists, nil
}