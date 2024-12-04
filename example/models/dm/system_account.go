// Code generated by SQLBoiler 4.17.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dm

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

// SystemAccount is an object representing the database table.
type SystemAccount struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"createdAt" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updatedAt" toml:"updated_at" yaml:"updated_at"`
	DeletedAt null.Time `boil:"deleted_at" json:"deletedAt,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *systemAccountR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L systemAccountL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SystemAccountColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

var SystemAccountTableColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "system_account.id",
	CreatedAt: "system_account.created_at",
	UpdatedAt: "system_account.updated_at",
	DeletedAt: "system_account.deleted_at",
}

// Generated where

var SystemAccountWhere = struct {
	ID        whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	DeletedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "\"system_account\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"system_account\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"system_account\".\"updated_at\""},
	DeletedAt: whereHelpernull_Time{field: "\"system_account\".\"deleted_at\""},
}

// SystemAccountRels is where relationship names are stored.
var SystemAccountRels = struct {
	Causers string
}{
	Causers: "Causers",
}

// systemAccountR is where relationships are stored.
type systemAccountR struct {
	Causers CauserSlice `boil:"Causers" json:"Causers" toml:"Causers" yaml:"Causers"`
}

// NewStruct creates a new relationship struct
func (*systemAccountR) NewStruct() *systemAccountR {
	return &systemAccountR{}
}

func (r *systemAccountR) GetCausers() CauserSlice {
	if r == nil {
		return nil
	}
	return r.Causers
}

// systemAccountL is where Load methods for each relationship are stored.
type systemAccountL struct{}

var (
	systemAccountAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at"}
	systemAccountColumnsWithoutDefault = []string{}
	systemAccountColumnsWithDefault    = []string{"id", "created_at", "updated_at", "deleted_at"}
	systemAccountPrimaryKeyColumns     = []string{"id"}
	systemAccountGeneratedColumns      = []string{}
)

type (
	// SystemAccountSlice is an alias for a slice of pointers to SystemAccount.
	// This should almost always be used instead of []SystemAccount.
	SystemAccountSlice []*SystemAccount
	// SystemAccountHook is the signature for custom SystemAccount hook methods
	SystemAccountHook func(context.Context, boil.ContextExecutor, *SystemAccount) error

	systemAccountQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	systemAccountType                 = reflect.TypeOf(&SystemAccount{})
	systemAccountMapping              = queries.MakeStructMapping(systemAccountType)
	systemAccountPrimaryKeyMapping, _ = queries.BindMapping(systemAccountType, systemAccountMapping, systemAccountPrimaryKeyColumns)
	systemAccountInsertCacheMut       sync.RWMutex
	systemAccountInsertCache          = make(map[string]insertCache)
	systemAccountUpdateCacheMut       sync.RWMutex
	systemAccountUpdateCache          = make(map[string]updateCache)
	systemAccountUpsertCacheMut       sync.RWMutex
	systemAccountUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var systemAccountAfterSelectMu sync.Mutex
var systemAccountAfterSelectHooks []SystemAccountHook

var systemAccountBeforeInsertMu sync.Mutex
var systemAccountBeforeInsertHooks []SystemAccountHook
var systemAccountAfterInsertMu sync.Mutex
var systemAccountAfterInsertHooks []SystemAccountHook

var systemAccountBeforeUpdateMu sync.Mutex
var systemAccountBeforeUpdateHooks []SystemAccountHook
var systemAccountAfterUpdateMu sync.Mutex
var systemAccountAfterUpdateHooks []SystemAccountHook

var systemAccountBeforeDeleteMu sync.Mutex
var systemAccountBeforeDeleteHooks []SystemAccountHook
var systemAccountAfterDeleteMu sync.Mutex
var systemAccountAfterDeleteHooks []SystemAccountHook

var systemAccountBeforeUpsertMu sync.Mutex
var systemAccountBeforeUpsertHooks []SystemAccountHook
var systemAccountAfterUpsertMu sync.Mutex
var systemAccountAfterUpsertHooks []SystemAccountHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SystemAccount) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range systemAccountAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SystemAccount) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range systemAccountBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SystemAccount) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range systemAccountAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SystemAccount) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range systemAccountBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SystemAccount) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range systemAccountAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SystemAccount) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range systemAccountBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SystemAccount) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range systemAccountAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SystemAccount) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range systemAccountBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SystemAccount) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range systemAccountAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSystemAccountHook registers your hook function for all future operations.
func AddSystemAccountHook(hookPoint boil.HookPoint, systemAccountHook SystemAccountHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		systemAccountAfterSelectMu.Lock()
		systemAccountAfterSelectHooks = append(systemAccountAfterSelectHooks, systemAccountHook)
		systemAccountAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		systemAccountBeforeInsertMu.Lock()
		systemAccountBeforeInsertHooks = append(systemAccountBeforeInsertHooks, systemAccountHook)
		systemAccountBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		systemAccountAfterInsertMu.Lock()
		systemAccountAfterInsertHooks = append(systemAccountAfterInsertHooks, systemAccountHook)
		systemAccountAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		systemAccountBeforeUpdateMu.Lock()
		systemAccountBeforeUpdateHooks = append(systemAccountBeforeUpdateHooks, systemAccountHook)
		systemAccountBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		systemAccountAfterUpdateMu.Lock()
		systemAccountAfterUpdateHooks = append(systemAccountAfterUpdateHooks, systemAccountHook)
		systemAccountAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		systemAccountBeforeDeleteMu.Lock()
		systemAccountBeforeDeleteHooks = append(systemAccountBeforeDeleteHooks, systemAccountHook)
		systemAccountBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		systemAccountAfterDeleteMu.Lock()
		systemAccountAfterDeleteHooks = append(systemAccountAfterDeleteHooks, systemAccountHook)
		systemAccountAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		systemAccountBeforeUpsertMu.Lock()
		systemAccountBeforeUpsertHooks = append(systemAccountBeforeUpsertHooks, systemAccountHook)
		systemAccountBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		systemAccountAfterUpsertMu.Lock()
		systemAccountAfterUpsertHooks = append(systemAccountAfterUpsertHooks, systemAccountHook)
		systemAccountAfterUpsertMu.Unlock()
	}
}

// One returns a single systemAccount record from the query.
func (q systemAccountQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SystemAccount, error) {
	o := &SystemAccount{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dm: failed to execute a one query for system_account")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SystemAccount records from the query.
func (q systemAccountQuery) All(ctx context.Context, exec boil.ContextExecutor) (SystemAccountSlice, error) {
	var o []*SystemAccount

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dm: failed to assign all query results to SystemAccount slice")
	}

	if len(systemAccountAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SystemAccount records in the query.
func (q systemAccountQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dm: failed to count system_account rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q systemAccountQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dm: failed to check if system_account exists")
	}

	return count > 0, nil
}

// Causers retrieves all the causer's Causers with an executor.
func (o *SystemAccount) Causers(mods ...qm.QueryMod) causerQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"causer\".\"system_account_id\"=?", o.ID),
	)

	return Causers(queryMods...)
}

// LoadCausers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (systemAccountL) LoadCausers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSystemAccount interface{}, mods queries.Applicator) error {
	var slice []*SystemAccount
	var object *SystemAccount

	if singular {
		var ok bool
		object, ok = maybeSystemAccount.(*SystemAccount)
		if !ok {
			object = new(SystemAccount)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSystemAccount)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSystemAccount))
			}
		}
	} else {
		s, ok := maybeSystemAccount.(*[]*SystemAccount)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSystemAccount)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSystemAccount))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &systemAccountR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &systemAccountR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`causer`),
		qm.WhereIn(`causer.system_account_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load causer")
	}

	var resultSlice []*Causer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice causer")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on causer")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for causer")
	}

	if len(causerAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Causers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &causerR{}
			}
			foreign.R.SystemAccount = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.SystemAccountID) {
				local.R.Causers = append(local.R.Causers, foreign)
				if foreign.R == nil {
					foreign.R = &causerR{}
				}
				foreign.R.SystemAccount = local
				break
			}
		}
	}

	return nil
}

// AddCausers adds the given related objects to the existing relationships
// of the system_account, optionally inserting them as new records.
// Appends related to o.R.Causers.
// Sets related.R.SystemAccount appropriately.
func (o *SystemAccount) AddCausers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Causer) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.SystemAccountID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"causer\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"system_account_id"}),
				strmangle.WhereClause("\"", "\"", 2, causerPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.SystemAccountID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &systemAccountR{
			Causers: related,
		}
	} else {
		o.R.Causers = append(o.R.Causers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &causerR{
				SystemAccount: o,
			}
		} else {
			rel.R.SystemAccount = o
		}
	}
	return nil
}

// SetCausers removes all previously related items of the
// system_account replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.SystemAccount's Causers accordingly.
// Replaces o.R.Causers with related.
// Sets related.R.SystemAccount's Causers accordingly.
func (o *SystemAccount) SetCausers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Causer) error {
	query := "update \"causer\" set \"system_account_id\" = null where \"system_account_id\" = $1"
	values := []interface{}{o.ID}
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
		for _, rel := range o.R.Causers {
			queries.SetScanner(&rel.SystemAccountID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.SystemAccount = nil
		}
		o.R.Causers = nil
	}

	return o.AddCausers(ctx, exec, insert, related...)
}

// RemoveCausers relationships from objects passed in.
// Removes related items from R.Causers (uses pointer comparison, removal does not keep order)
// Sets related.R.SystemAccount.
func (o *SystemAccount) RemoveCausers(ctx context.Context, exec boil.ContextExecutor, related ...*Causer) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.SystemAccountID, nil)
		if rel.R != nil {
			rel.R.SystemAccount = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("system_account_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Causers {
			if rel != ri {
				continue
			}

			ln := len(o.R.Causers)
			if ln > 1 && i < ln-1 {
				o.R.Causers[i] = o.R.Causers[ln-1]
			}
			o.R.Causers = o.R.Causers[:ln-1]
			break
		}
	}

	return nil
}

// SystemAccounts retrieves all the records using an executor.
func SystemAccounts(mods ...qm.QueryMod) systemAccountQuery {
	mods = append(mods, qm.From("\"system_account\""), qmhelper.WhereIsNull("\"system_account\".\"deleted_at\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"system_account\".*"})
	}

	return systemAccountQuery{q}
}

// FindSystemAccount retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSystemAccount(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*SystemAccount, error) {
	systemAccountObj := &SystemAccount{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"system_account\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, systemAccountObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dm: unable to select from system_account")
	}

	if err = systemAccountObj.doAfterSelectHooks(ctx, exec); err != nil {
		return systemAccountObj, err
	}

	return systemAccountObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SystemAccount) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dm: no system_account provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(systemAccountColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	systemAccountInsertCacheMut.RLock()
	cache, cached := systemAccountInsertCache[key]
	systemAccountInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			systemAccountAllColumns,
			systemAccountColumnsWithDefault,
			systemAccountColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(systemAccountType, systemAccountMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(systemAccountType, systemAccountMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"system_account\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"system_account\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "dm: unable to insert into system_account")
	}

	if !cached {
		systemAccountInsertCacheMut.Lock()
		systemAccountInsertCache[key] = cache
		systemAccountInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SystemAccount.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SystemAccount) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	systemAccountUpdateCacheMut.RLock()
	cache, cached := systemAccountUpdateCache[key]
	systemAccountUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			systemAccountAllColumns,
			systemAccountPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dm: unable to update system_account, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"system_account\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, systemAccountPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(systemAccountType, systemAccountMapping, append(wl, systemAccountPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dm: unable to update system_account row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dm: failed to get rows affected by update for system_account")
	}

	if !cached {
		systemAccountUpdateCacheMut.Lock()
		systemAccountUpdateCache[key] = cache
		systemAccountUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q systemAccountQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dm: unable to update all for system_account")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dm: unable to retrieve rows affected for system_account")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SystemAccountSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dm: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), systemAccountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"system_account\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, systemAccountPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dm: unable to update all in systemAccount slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dm: unable to retrieve rows affected all in update all systemAccount")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SystemAccount) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("dm: no system_account provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(systemAccountColumnsWithDefault, o)

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

	systemAccountUpsertCacheMut.RLock()
	cache, cached := systemAccountUpsertCache[key]
	systemAccountUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			systemAccountAllColumns,
			systemAccountColumnsWithDefault,
			systemAccountColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			systemAccountAllColumns,
			systemAccountPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dm: unable to upsert system_account, could not build update column list")
		}

		ret := strmangle.SetComplement(systemAccountAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(systemAccountPrimaryKeyColumns) == 0 {
				return errors.New("dm: unable to upsert system_account, could not build conflict column list")
			}

			conflict = make([]string, len(systemAccountPrimaryKeyColumns))
			copy(conflict, systemAccountPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"system_account\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(systemAccountType, systemAccountMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(systemAccountType, systemAccountMapping, ret)
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
		return errors.Wrap(err, "dm: unable to upsert system_account")
	}

	if !cached {
		systemAccountUpsertCacheMut.Lock()
		systemAccountUpsertCache[key] = cache
		systemAccountUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SystemAccount record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SystemAccount) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("dm: no SystemAccount provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), systemAccountPrimaryKeyMapping)
		sql = "DELETE FROM \"system_account\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"system_account\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(systemAccountType, systemAccountMapping, append(wl, systemAccountPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), valueMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dm: unable to delete from system_account")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dm: failed to get rows affected by delete for system_account")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q systemAccountQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dm: no systemAccountQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dm: unable to delete all from system_account")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dm: failed to get rows affected by deleteall for system_account")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SystemAccountSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(systemAccountBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), systemAccountPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"system_account\" WHERE " +
			strmangle.WhereInClause(string(dialect.LQ), string(dialect.RQ), 1, systemAccountPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), systemAccountPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"system_account\" SET %s WHERE "+
			strmangle.WhereInClause(string(dialect.LQ), string(dialect.RQ), 2, systemAccountPrimaryKeyColumns, len(o)),
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		args = append([]interface{}{currTime}, args...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dm: unable to delete all from systemAccount slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dm: failed to get rows affected by deleteall for system_account")
	}

	if len(systemAccountAfterDeleteHooks) != 0 {
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
func (o *SystemAccount) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSystemAccount(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SystemAccountSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SystemAccountSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), systemAccountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"system_account\".* FROM \"system_account\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, systemAccountPrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dm: unable to reload all in SystemAccountSlice")
	}

	*o = slice

	return nil
}

// SystemAccountExists checks if the SystemAccount row exists.
func SystemAccountExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"system_account\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dm: unable to check if system_account exists")
	}

	return exists, nil
}

// Exists checks if the SystemAccount row exists.
func (o *SystemAccount) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SystemAccountExists(ctx, exec, o.ID)
}
