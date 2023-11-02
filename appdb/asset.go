// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package appdb

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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Asset is an object representing the database table.
type Asset struct {
	ID              int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ConfigurationID int64      `boil:"configuration_id" json:"configuration_id" toml:"configuration_id" yaml:"configuration_id"`
	ProjectID       string     `boil:"project_id" json:"project_id" toml:"project_id" yaml:"project_id"`
	GlobalAssetID   string     `boil:"global_asset_id" json:"global_asset_id" toml:"global_asset_id" yaml:"global_asset_id"`
	AssetID         null.Int32 `boil:"asset_id" json:"asset_id,omitempty" toml:"asset_id" yaml:"asset_id,omitempty"`

	R *assetR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L assetL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AssetColumns = struct {
	ID              string
	ConfigurationID string
	ProjectID       string
	GlobalAssetID   string
	AssetID         string
}{
	ID:              "id",
	ConfigurationID: "configuration_id",
	ProjectID:       "project_id",
	GlobalAssetID:   "global_asset_id",
	AssetID:         "asset_id",
}

var AssetTableColumns = struct {
	ID              string
	ConfigurationID string
	ProjectID       string
	GlobalAssetID   string
	AssetID         string
}{
	ID:              "asset.id",
	ConfigurationID: "asset.configuration_id",
	ProjectID:       "asset.project_id",
	GlobalAssetID:   "asset.global_asset_id",
	AssetID:         "asset.asset_id",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int32 struct{ field string }

func (w whereHelpernull_Int32) EQ(x null.Int32) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int32) NEQ(x null.Int32) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int32) LT(x null.Int32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int32) LTE(x null.Int32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int32) GT(x null.Int32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int32) GTE(x null.Int32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int32) IN(slice []int32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int32) NIN(slice []int32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int32) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int32) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var AssetWhere = struct {
	ID              whereHelperint64
	ConfigurationID whereHelperint64
	ProjectID       whereHelperstring
	GlobalAssetID   whereHelperstring
	AssetID         whereHelpernull_Int32
}{
	ID:              whereHelperint64{field: "\"mystrom\".\"asset\".\"id\""},
	ConfigurationID: whereHelperint64{field: "\"mystrom\".\"asset\".\"configuration_id\""},
	ProjectID:       whereHelperstring{field: "\"mystrom\".\"asset\".\"project_id\""},
	GlobalAssetID:   whereHelperstring{field: "\"mystrom\".\"asset\".\"global_asset_id\""},
	AssetID:         whereHelpernull_Int32{field: "\"mystrom\".\"asset\".\"asset_id\""},
}

// AssetRels is where relationship names are stored.
var AssetRels = struct {
	Configuration string
}{
	Configuration: "Configuration",
}

// assetR is where relationships are stored.
type assetR struct {
	Configuration *Configuration `boil:"Configuration" json:"Configuration" toml:"Configuration" yaml:"Configuration"`
}

// NewStruct creates a new relationship struct
func (*assetR) NewStruct() *assetR {
	return &assetR{}
}

func (r *assetR) GetConfiguration() *Configuration {
	if r == nil {
		return nil
	}
	return r.Configuration
}

// assetL is where Load methods for each relationship are stored.
type assetL struct{}

var (
	assetAllColumns            = []string{"id", "configuration_id", "project_id", "global_asset_id", "asset_id"}
	assetColumnsWithoutDefault = []string{"project_id", "global_asset_id"}
	assetColumnsWithDefault    = []string{"id", "configuration_id", "asset_id"}
	assetPrimaryKeyColumns     = []string{"id"}
	assetGeneratedColumns      = []string{}
)

type (
	// AssetSlice is an alias for a slice of pointers to Asset.
	// This should almost always be used instead of []Asset.
	AssetSlice []*Asset
	// AssetHook is the signature for custom Asset hook methods
	AssetHook func(context.Context, boil.ContextExecutor, *Asset) error

	assetQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	assetType                 = reflect.TypeOf(&Asset{})
	assetMapping              = queries.MakeStructMapping(assetType)
	assetPrimaryKeyMapping, _ = queries.BindMapping(assetType, assetMapping, assetPrimaryKeyColumns)
	assetInsertCacheMut       sync.RWMutex
	assetInsertCache          = make(map[string]insertCache)
	assetUpdateCacheMut       sync.RWMutex
	assetUpdateCache          = make(map[string]updateCache)
	assetUpsertCacheMut       sync.RWMutex
	assetUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var assetAfterSelectHooks []AssetHook

var assetBeforeInsertHooks []AssetHook
var assetAfterInsertHooks []AssetHook

var assetBeforeUpdateHooks []AssetHook
var assetAfterUpdateHooks []AssetHook

var assetBeforeDeleteHooks []AssetHook
var assetAfterDeleteHooks []AssetHook

var assetBeforeUpsertHooks []AssetHook
var assetAfterUpsertHooks []AssetHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Asset) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range assetAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Asset) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range assetBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Asset) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range assetAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Asset) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range assetBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Asset) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range assetAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Asset) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range assetBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Asset) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range assetAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Asset) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range assetBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Asset) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range assetAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAssetHook registers your hook function for all future operations.
func AddAssetHook(hookPoint boil.HookPoint, assetHook AssetHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		assetAfterSelectHooks = append(assetAfterSelectHooks, assetHook)
	case boil.BeforeInsertHook:
		assetBeforeInsertHooks = append(assetBeforeInsertHooks, assetHook)
	case boil.AfterInsertHook:
		assetAfterInsertHooks = append(assetAfterInsertHooks, assetHook)
	case boil.BeforeUpdateHook:
		assetBeforeUpdateHooks = append(assetBeforeUpdateHooks, assetHook)
	case boil.AfterUpdateHook:
		assetAfterUpdateHooks = append(assetAfterUpdateHooks, assetHook)
	case boil.BeforeDeleteHook:
		assetBeforeDeleteHooks = append(assetBeforeDeleteHooks, assetHook)
	case boil.AfterDeleteHook:
		assetAfterDeleteHooks = append(assetAfterDeleteHooks, assetHook)
	case boil.BeforeUpsertHook:
		assetBeforeUpsertHooks = append(assetBeforeUpsertHooks, assetHook)
	case boil.AfterUpsertHook:
		assetAfterUpsertHooks = append(assetAfterUpsertHooks, assetHook)
	}
}

// OneG returns a single asset record from the query using the global executor.
func (q assetQuery) OneG(ctx context.Context) (*Asset, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single asset record from the query.
func (q assetQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Asset, error) {
	o := &Asset{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "appdb: failed to execute a one query for asset")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Asset records from the query using the global executor.
func (q assetQuery) AllG(ctx context.Context) (AssetSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Asset records from the query.
func (q assetQuery) All(ctx context.Context, exec boil.ContextExecutor) (AssetSlice, error) {
	var o []*Asset

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "appdb: failed to assign all query results to Asset slice")
	}

	if len(assetAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Asset records in the query using the global executor
func (q assetQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Asset records in the query.
func (q assetQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to count asset rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q assetQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q assetQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "appdb: failed to check if asset exists")
	}

	return count > 0, nil
}

// Configuration pointed to by the foreign key.
func (o *Asset) Configuration(mods ...qm.QueryMod) configurationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ConfigurationID),
	}

	queryMods = append(queryMods, mods...)

	return Configurations(queryMods...)
}

// LoadConfiguration allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (assetL) LoadConfiguration(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAsset interface{}, mods queries.Applicator) error {
	var slice []*Asset
	var object *Asset

	if singular {
		var ok bool
		object, ok = maybeAsset.(*Asset)
		if !ok {
			object = new(Asset)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAsset)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAsset))
			}
		}
	} else {
		s, ok := maybeAsset.(*[]*Asset)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAsset)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAsset))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &assetR{}
		}
		args = append(args, object.ConfigurationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &assetR{}
			}

			for _, a := range args {
				if a == obj.ConfigurationID {
					continue Outer
				}
			}

			args = append(args, obj.ConfigurationID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`mystrom.configuration`),
		qm.WhereIn(`mystrom.configuration.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Configuration")
	}

	var resultSlice []*Configuration
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Configuration")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for configuration")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for configuration")
	}

	if len(configurationAfterSelectHooks) != 0 {
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
		object.R.Configuration = foreign
		if foreign.R == nil {
			foreign.R = &configurationR{}
		}
		foreign.R.Assets = append(foreign.R.Assets, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ConfigurationID == foreign.ID {
				local.R.Configuration = foreign
				if foreign.R == nil {
					foreign.R = &configurationR{}
				}
				foreign.R.Assets = append(foreign.R.Assets, local)
				break
			}
		}
	}

	return nil
}

// SetConfigurationG of the asset to the related item.
// Sets o.R.Configuration to related.
// Adds o to related.R.Assets.
// Uses the global database handle.
func (o *Asset) SetConfigurationG(ctx context.Context, insert bool, related *Configuration) error {
	return o.SetConfiguration(ctx, boil.GetContextDB(), insert, related)
}

// SetConfiguration of the asset to the related item.
// Sets o.R.Configuration to related.
// Adds o to related.R.Assets.
func (o *Asset) SetConfiguration(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Configuration) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"mystrom\".\"asset\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"configuration_id"}),
		strmangle.WhereClause("\"", "\"", 2, assetPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ConfigurationID = related.ID
	if o.R == nil {
		o.R = &assetR{
			Configuration: related,
		}
	} else {
		o.R.Configuration = related
	}

	if related.R == nil {
		related.R = &configurationR{
			Assets: AssetSlice{o},
		}
	} else {
		related.R.Assets = append(related.R.Assets, o)
	}

	return nil
}

// Assets retrieves all the records using an executor.
func Assets(mods ...qm.QueryMod) assetQuery {
	mods = append(mods, qm.From("\"mystrom\".\"asset\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"mystrom\".\"asset\".*"})
	}

	return assetQuery{q}
}

// FindAssetG retrieves a single record by ID.
func FindAssetG(ctx context.Context, iD int64, selectCols ...string) (*Asset, error) {
	return FindAsset(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindAsset retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAsset(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Asset, error) {
	assetObj := &Asset{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"mystrom\".\"asset\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, assetObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "appdb: unable to select from asset")
	}

	if err = assetObj.doAfterSelectHooks(ctx, exec); err != nil {
		return assetObj, err
	}

	return assetObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Asset) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Asset) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("appdb: no asset provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(assetColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	assetInsertCacheMut.RLock()
	cache, cached := assetInsertCache[key]
	assetInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			assetAllColumns,
			assetColumnsWithDefault,
			assetColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(assetType, assetMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(assetType, assetMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"mystrom\".\"asset\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"mystrom\".\"asset\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "appdb: unable to insert into asset")
	}

	if !cached {
		assetInsertCacheMut.Lock()
		assetInsertCache[key] = cache
		assetInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Asset record using the global executor.
// See Update for more documentation.
func (o *Asset) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Asset.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Asset) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	assetUpdateCacheMut.RLock()
	cache, cached := assetUpdateCache[key]
	assetUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			assetAllColumns,
			assetPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("appdb: unable to update asset, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"mystrom\".\"asset\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, assetPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(assetType, assetMapping, append(wl, assetPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "appdb: unable to update asset row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by update for asset")
	}

	if !cached {
		assetUpdateCacheMut.Lock()
		assetUpdateCache[key] = cache
		assetUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q assetQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q assetQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to update all for asset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to retrieve rows affected for asset")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AssetSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AssetSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("appdb: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), assetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"mystrom\".\"asset\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, assetPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to update all in asset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to retrieve rows affected all in update all asset")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Asset) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Asset) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("appdb: no asset provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(assetColumnsWithDefault, o)

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

	assetUpsertCacheMut.RLock()
	cache, cached := assetUpsertCache[key]
	assetUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			assetAllColumns,
			assetColumnsWithDefault,
			assetColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			assetAllColumns,
			assetPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("appdb: unable to upsert asset, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(assetPrimaryKeyColumns))
			copy(conflict, assetPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"mystrom\".\"asset\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(assetType, assetMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(assetType, assetMapping, ret)
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
		return errors.Wrap(err, "appdb: unable to upsert asset")
	}

	if !cached {
		assetUpsertCacheMut.Lock()
		assetUpsertCache[key] = cache
		assetUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Asset record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Asset) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Asset record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Asset) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("appdb: no Asset provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), assetPrimaryKeyMapping)
	sql := "DELETE FROM \"mystrom\".\"asset\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete from asset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by delete for asset")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q assetQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q assetQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("appdb: no assetQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete all from asset")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by deleteall for asset")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AssetSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AssetSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(assetBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), assetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"mystrom\".\"asset\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, assetPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete all from asset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by deleteall for asset")
	}

	if len(assetAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Asset) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("appdb: no Asset provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Asset) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAsset(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AssetSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("appdb: empty AssetSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AssetSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AssetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), assetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"mystrom\".\"asset\".* FROM \"mystrom\".\"asset\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, assetPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "appdb: unable to reload all in AssetSlice")
	}

	*o = slice

	return nil
}

// AssetExistsG checks if the Asset row exists.
func AssetExistsG(ctx context.Context, iD int64) (bool, error) {
	return AssetExists(ctx, boil.GetContextDB(), iD)
}

// AssetExists checks if the Asset row exists.
func AssetExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"mystrom\".\"asset\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "appdb: unable to check if asset exists")
	}

	return exists, nil
}

// Exists checks if the Asset row exists.
func (o *Asset) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AssetExists(ctx, exec, o.ID)
}
