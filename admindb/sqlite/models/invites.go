// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
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

// Invite is an object representing the database table.
type Invite struct {
	ID              int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Token           string `boil:"token" json:"token" toml:"token" yaml:"token"`
	CreatedBy       int64  `boil:"created_by" json:"created_by" toml:"created_by" yaml:"created_by"`
	AliasSuggestion string `boil:"alias_suggestion" json:"alias_suggestion" toml:"alias_suggestion" yaml:"alias_suggestion"`
	Active          bool   `boil:"active" json:"active" toml:"active" yaml:"active"`

	R *inviteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L inviteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InviteColumns = struct {
	ID              string
	Token           string
	CreatedBy       string
	AliasSuggestion string
	Active          string
}{
	ID:              "id",
	Token:           "token",
	CreatedBy:       "created_by",
	AliasSuggestion: "alias_suggestion",
	Active:          "active",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var InviteWhere = struct {
	ID              whereHelperint64
	Token           whereHelperstring
	CreatedBy       whereHelperint64
	AliasSuggestion whereHelperstring
	Active          whereHelperbool
}{
	ID:              whereHelperint64{field: "\"invites\".\"id\""},
	Token:           whereHelperstring{field: "\"invites\".\"token\""},
	CreatedBy:       whereHelperint64{field: "\"invites\".\"created_by\""},
	AliasSuggestion: whereHelperstring{field: "\"invites\".\"alias_suggestion\""},
	Active:          whereHelperbool{field: "\"invites\".\"active\""},
}

// InviteRels is where relationship names are stored.
var InviteRels = struct {
	CreatedByAuthFallback string
}{
	CreatedByAuthFallback: "CreatedByAuthFallback",
}

// inviteR is where relationships are stored.
type inviteR struct {
	CreatedByAuthFallback *AuthFallback `boil:"CreatedByAuthFallback" json:"CreatedByAuthFallback" toml:"CreatedByAuthFallback" yaml:"CreatedByAuthFallback"`
}

// NewStruct creates a new relationship struct
func (*inviteR) NewStruct() *inviteR {
	return &inviteR{}
}

// inviteL is where Load methods for each relationship are stored.
type inviteL struct{}

var (
	inviteAllColumns            = []string{"id", "token", "created_by", "alias_suggestion", "active"}
	inviteColumnsWithoutDefault = []string{}
	inviteColumnsWithDefault    = []string{"id", "token", "created_by", "alias_suggestion", "active"}
	invitePrimaryKeyColumns     = []string{"id"}
)

type (
	// InviteSlice is an alias for a slice of pointers to Invite.
	// This should generally be used opposed to []Invite.
	InviteSlice []*Invite
	// InviteHook is the signature for custom Invite hook methods
	InviteHook func(context.Context, boil.ContextExecutor, *Invite) error

	inviteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	inviteType                 = reflect.TypeOf(&Invite{})
	inviteMapping              = queries.MakeStructMapping(inviteType)
	invitePrimaryKeyMapping, _ = queries.BindMapping(inviteType, inviteMapping, invitePrimaryKeyColumns)
	inviteInsertCacheMut       sync.RWMutex
	inviteInsertCache          = make(map[string]insertCache)
	inviteUpdateCacheMut       sync.RWMutex
	inviteUpdateCache          = make(map[string]updateCache)
	inviteUpsertCacheMut       sync.RWMutex
	inviteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var inviteBeforeInsertHooks []InviteHook
var inviteBeforeUpdateHooks []InviteHook
var inviteBeforeDeleteHooks []InviteHook
var inviteBeforeUpsertHooks []InviteHook

var inviteAfterInsertHooks []InviteHook
var inviteAfterSelectHooks []InviteHook
var inviteAfterUpdateHooks []InviteHook
var inviteAfterDeleteHooks []InviteHook
var inviteAfterUpsertHooks []InviteHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Invite) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Invite) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Invite) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Invite) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Invite) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Invite) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Invite) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Invite) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Invite) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range inviteAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInviteHook registers your hook function for all future operations.
func AddInviteHook(hookPoint boil.HookPoint, inviteHook InviteHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		inviteBeforeInsertHooks = append(inviteBeforeInsertHooks, inviteHook)
	case boil.BeforeUpdateHook:
		inviteBeforeUpdateHooks = append(inviteBeforeUpdateHooks, inviteHook)
	case boil.BeforeDeleteHook:
		inviteBeforeDeleteHooks = append(inviteBeforeDeleteHooks, inviteHook)
	case boil.BeforeUpsertHook:
		inviteBeforeUpsertHooks = append(inviteBeforeUpsertHooks, inviteHook)
	case boil.AfterInsertHook:
		inviteAfterInsertHooks = append(inviteAfterInsertHooks, inviteHook)
	case boil.AfterSelectHook:
		inviteAfterSelectHooks = append(inviteAfterSelectHooks, inviteHook)
	case boil.AfterUpdateHook:
		inviteAfterUpdateHooks = append(inviteAfterUpdateHooks, inviteHook)
	case boil.AfterDeleteHook:
		inviteAfterDeleteHooks = append(inviteAfterDeleteHooks, inviteHook)
	case boil.AfterUpsertHook:
		inviteAfterUpsertHooks = append(inviteAfterUpsertHooks, inviteHook)
	}
}

// One returns a single invite record from the query.
func (q inviteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Invite, error) {
	o := &Invite{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for invites")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Invite records from the query.
func (q inviteQuery) All(ctx context.Context, exec boil.ContextExecutor) (InviteSlice, error) {
	var o []*Invite

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Invite slice")
	}

	if len(inviteAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Invite records in the query.
func (q inviteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count invites rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q inviteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if invites exists")
	}

	return count > 0, nil
}

// CreatedByAuthFallback pointed to by the foreign key.
func (o *Invite) CreatedByAuthFallback(mods ...qm.QueryMod) authFallbackQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CreatedBy),
	}

	queryMods = append(queryMods, mods...)

	query := AuthFallbacks(queryMods...)
	queries.SetFrom(query.Query, "\"auth_fallback\"")

	return query
}

// LoadCreatedByAuthFallback allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (inviteL) LoadCreatedByAuthFallback(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInvite interface{}, mods queries.Applicator) error {
	var slice []*Invite
	var object *Invite

	if singular {
		object = maybeInvite.(*Invite)
	} else {
		slice = *maybeInvite.(*[]*Invite)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &inviteR{}
		}
		args = append(args, object.CreatedBy)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &inviteR{}
			}

			for _, a := range args {
				if a == obj.CreatedBy {
					continue Outer
				}
			}

			args = append(args, obj.CreatedBy)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`auth_fallback`),
		qm.WhereIn(`auth_fallback.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthFallback")
	}

	var resultSlice []*AuthFallback
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthFallback")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for auth_fallback")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for auth_fallback")
	}

	if len(inviteAfterSelectHooks) != 0 {
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
		object.R.CreatedByAuthFallback = foreign
		if foreign.R == nil {
			foreign.R = &authFallbackR{}
		}
		foreign.R.CreatedByInvites = append(foreign.R.CreatedByInvites, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CreatedBy == foreign.ID {
				local.R.CreatedByAuthFallback = foreign
				if foreign.R == nil {
					foreign.R = &authFallbackR{}
				}
				foreign.R.CreatedByInvites = append(foreign.R.CreatedByInvites, local)
				break
			}
		}
	}

	return nil
}

// SetCreatedByAuthFallback of the invite to the related item.
// Sets o.R.CreatedByAuthFallback to related.
// Adds o to related.R.CreatedByInvites.
func (o *Invite) SetCreatedByAuthFallback(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AuthFallback) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"invites\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"created_by"}),
		strmangle.WhereClause("\"", "\"", 0, invitePrimaryKeyColumns),
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

	o.CreatedBy = related.ID
	if o.R == nil {
		o.R = &inviteR{
			CreatedByAuthFallback: related,
		}
	} else {
		o.R.CreatedByAuthFallback = related
	}

	if related.R == nil {
		related.R = &authFallbackR{
			CreatedByInvites: InviteSlice{o},
		}
	} else {
		related.R.CreatedByInvites = append(related.R.CreatedByInvites, o)
	}

	return nil
}

// Invites retrieves all the records using an executor.
func Invites(mods ...qm.QueryMod) inviteQuery {
	mods = append(mods, qm.From("\"invites\""))
	return inviteQuery{NewQuery(mods...)}
}

// FindInvite retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInvite(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Invite, error) {
	inviteObj := &Invite{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"invites\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, inviteObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from invites")
	}

	return inviteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Invite) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no invites provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(inviteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	inviteInsertCacheMut.RLock()
	cache, cached := inviteInsertCache[key]
	inviteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			inviteAllColumns,
			inviteColumnsWithDefault,
			inviteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(inviteType, inviteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(inviteType, inviteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"invites\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"invites\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"invites\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, invitePrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into invites")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == inviteMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for invites")
	}

CacheNoHooks:
	if !cached {
		inviteInsertCacheMut.Lock()
		inviteInsertCache[key] = cache
		inviteInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Invite.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Invite) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	inviteUpdateCacheMut.RLock()
	cache, cached := inviteUpdateCache[key]
	inviteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			inviteAllColumns,
			invitePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update invites, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"invites\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, invitePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(inviteType, inviteMapping, append(wl, invitePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update invites row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for invites")
	}

	if !cached {
		inviteUpdateCacheMut.Lock()
		inviteUpdateCache[key] = cache
		inviteUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q inviteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for invites")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for invites")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InviteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"invites\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, invitePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in invite slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all invite")
	}
	return rowsAff, nil
}

// Delete deletes a single Invite record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Invite) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Invite provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), invitePrimaryKeyMapping)
	sql := "DELETE FROM \"invites\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from invites")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for invites")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q inviteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no inviteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from invites")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for invites")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InviteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(inviteBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"invites\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, invitePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from invite slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for invites")
	}

	if len(inviteAfterDeleteHooks) != 0 {
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
func (o *Invite) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInvite(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InviteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InviteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"invites\".* FROM \"invites\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, invitePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InviteSlice")
	}

	*o = slice

	return nil
}

// InviteExists checks if the Invite row exists.
func InviteExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"invites\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if invites exists")
	}

	return exists, nil
}
