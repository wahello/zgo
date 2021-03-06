// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/suisrc/zgo/app/model/ent/account"
	"github.com/suisrc/zgo/app/model/ent/predicate"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks      []Hook
	mutation   *AccountMutation
	predicates []predicate.Account
}

// Where adds a new predicate for the builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetPid sets the pid field.
func (au *AccountUpdate) SetPid(s string) *AccountUpdate {
	au.mutation.SetPid(s)
	return au
}

// SetAccount sets the account field.
func (au *AccountUpdate) SetAccount(s string) *AccountUpdate {
	au.mutation.SetAccount(s)
	return au
}

// SetAccountTyp sets the account_typ field.
func (au *AccountUpdate) SetAccountTyp(s string) *AccountUpdate {
	au.mutation.SetAccountTyp(s)
	return au
}

// SetAccountKid sets the account_kid field.
func (au *AccountUpdate) SetAccountKid(s string) *AccountUpdate {
	au.mutation.SetAccountKid(s)
	return au
}

// SetPassword sets the password field.
func (au *AccountUpdate) SetPassword(s string) *AccountUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetPasswordSalt sets the password_salt field.
func (au *AccountUpdate) SetPasswordSalt(s string) *AccountUpdate {
	au.mutation.SetPasswordSalt(s)
	return au
}

// SetPasswordType sets the password_type field.
func (au *AccountUpdate) SetPasswordType(s string) *AccountUpdate {
	au.mutation.SetPasswordType(s)
	return au
}

// SetVerifySecret sets the verify_secret field.
func (au *AccountUpdate) SetVerifySecret(s string) *AccountUpdate {
	au.mutation.SetVerifySecret(s)
	return au
}

// SetVerifyType sets the verify_type field.
func (au *AccountUpdate) SetVerifyType(s string) *AccountUpdate {
	au.mutation.SetVerifyType(s)
	return au
}

// SetUserID sets the user_id field.
func (au *AccountUpdate) SetUserID(i int) *AccountUpdate {
	au.mutation.ResetUserID()
	au.mutation.SetUserID(i)
	return au
}

// AddUserID adds i to user_id.
func (au *AccountUpdate) AddUserID(i int) *AccountUpdate {
	au.mutation.AddUserID(i)
	return au
}

// SetRoleID sets the role_id field.
func (au *AccountUpdate) SetRoleID(i int) *AccountUpdate {
	au.mutation.ResetRoleID()
	au.mutation.SetRoleID(i)
	return au
}

// AddRoleID adds i to role_id.
func (au *AccountUpdate) AddRoleID(i int) *AccountUpdate {
	au.mutation.AddRoleID(i)
	return au
}

// SetStatus sets the status field.
func (au *AccountUpdate) SetStatus(i int) *AccountUpdate {
	au.mutation.ResetStatus()
	au.mutation.SetStatus(i)
	return au
}

// AddStatus adds i to status.
func (au *AccountUpdate) AddStatus(i int) *AccountUpdate {
	au.mutation.AddStatus(i)
	return au
}

// SetDescription sets the description field.
func (au *AccountUpdate) SetDescription(s string) *AccountUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetOa2Token sets the oa2_token field.
func (au *AccountUpdate) SetOa2Token(s string) *AccountUpdate {
	au.mutation.SetOa2Token(s)
	return au
}

// SetOa2Expired sets the oa2_expired field.
func (au *AccountUpdate) SetOa2Expired(t time.Time) *AccountUpdate {
	au.mutation.SetOa2Expired(t)
	return au
}

// SetOa2Fake sets the oa2_fake field.
func (au *AccountUpdate) SetOa2Fake(s string) *AccountUpdate {
	au.mutation.SetOa2Fake(s)
	return au
}

// SetOa2Client sets the oa2_client field.
func (au *AccountUpdate) SetOa2Client(i int) *AccountUpdate {
	au.mutation.ResetOa2Client()
	au.mutation.SetOa2Client(i)
	return au
}

// AddOa2Client adds i to oa2_client.
func (au *AccountUpdate) AddOa2Client(i int) *AccountUpdate {
	au.mutation.AddOa2Client(i)
	return au
}

// SetCreator sets the creator field.
func (au *AccountUpdate) SetCreator(s string) *AccountUpdate {
	au.mutation.SetCreator(s)
	return au
}

// SetCreatedAt sets the created_at field.
func (au *AccountUpdate) SetCreatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (au *AccountUpdate) SetNillableCreatedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the updated_at field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (au *AccountUpdate) SetNillableUpdatedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// SetVersion sets the version field.
func (au *AccountUpdate) SetVersion(i int) *AccountUpdate {
	au.mutation.ResetVersion()
	au.mutation.SetVersion(i)
	return au
}

// SetNillableVersion sets the version field if the given value is not nil.
func (au *AccountUpdate) SetNillableVersion(i *int) *AccountUpdate {
	if i != nil {
		au.SetVersion(*i)
	}
	return au
}

// AddVersion adds i to version.
func (au *AccountUpdate) AddVersion(i int) *AccountUpdate {
	au.mutation.AddVersion(i)
	return au
}

// SetString1 sets the string_1 field.
func (au *AccountUpdate) SetString1(s string) *AccountUpdate {
	au.mutation.SetString1(s)
	return au
}

// SetString2 sets the string_2 field.
func (au *AccountUpdate) SetString2(s string) *AccountUpdate {
	au.mutation.SetString2(s)
	return au
}

// SetString3 sets the string_3 field.
func (au *AccountUpdate) SetString3(s string) *AccountUpdate {
	au.mutation.SetString3(s)
	return au
}

// SetNumber1 sets the number_1 field.
func (au *AccountUpdate) SetNumber1(i int) *AccountUpdate {
	au.mutation.ResetNumber1()
	au.mutation.SetNumber1(i)
	return au
}

// AddNumber1 adds i to number_1.
func (au *AccountUpdate) AddNumber1(i int) *AccountUpdate {
	au.mutation.AddNumber1(i)
	return au
}

// SetNumber2 sets the number_2 field.
func (au *AccountUpdate) SetNumber2(i int) *AccountUpdate {
	au.mutation.ResetNumber2()
	au.mutation.SetNumber2(i)
	return au
}

// AddNumber2 adds i to number_2.
func (au *AccountUpdate) AddNumber2(i int) *AccountUpdate {
	au.mutation.AddNumber2(i)
	return au
}

// SetNumber3 sets the number_3 field.
func (au *AccountUpdate) SetNumber3(i int) *AccountUpdate {
	au.mutation.ResetNumber3()
	au.mutation.SetNumber3(i)
	return au
}

// AddNumber3 adds i to number_3.
func (au *AccountUpdate) AddNumber3(i int) *AccountUpdate {
	au.mutation.AddNumber3(i)
	return au
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Pid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPid,
		})
	}
	if value, ok := au.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAccount,
		})
	}
	if value, ok := au.mutation.AccountTyp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAccountTyp,
		})
	}
	if value, ok := au.mutation.AccountKid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAccountKid,
		})
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPassword,
		})
	}
	if value, ok := au.mutation.PasswordSalt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPasswordSalt,
		})
	}
	if value, ok := au.mutation.PasswordType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPasswordType,
		})
	}
	if value, ok := au.mutation.VerifySecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldVerifySecret,
		})
	}
	if value, ok := au.mutation.VerifyType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldVerifyType,
		})
	}
	if value, ok := au.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldUserID,
		})
	}
	if value, ok := au.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldUserID,
		})
	}
	if value, ok := au.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldRoleID,
		})
	}
	if value, ok := au.mutation.AddedRoleID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldRoleID,
		})
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldStatus,
		})
	}
	if value, ok := au.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldStatus,
		})
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldDescription,
		})
	}
	if value, ok := au.mutation.Oa2Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldOa2Token,
		})
	}
	if value, ok := au.mutation.Oa2Expired(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldOa2Expired,
		})
	}
	if value, ok := au.mutation.Oa2Fake(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldOa2Fake,
		})
	}
	if value, ok := au.mutation.Oa2Client(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldOa2Client,
		})
	}
	if value, ok := au.mutation.AddedOa2Client(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldOa2Client,
		})
	}
	if value, ok := au.mutation.Creator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldCreator,
		})
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldVersion,
		})
	}
	if value, ok := au.mutation.AddedVersion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldVersion,
		})
	}
	if value, ok := au.mutation.String1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldString1,
		})
	}
	if value, ok := au.mutation.String2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldString2,
		})
	}
	if value, ok := au.mutation.String3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldString3,
		})
	}
	if value, ok := au.mutation.Number1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber1,
		})
	}
	if value, ok := au.mutation.AddedNumber1(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber1,
		})
	}
	if value, ok := au.mutation.Number2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber2,
		})
	}
	if value, ok := au.mutation.AddedNumber2(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber2,
		})
	}
	if value, ok := au.mutation.Number3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber3,
		})
	}
	if value, ok := au.mutation.AddedNumber3(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber3,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// SetPid sets the pid field.
func (auo *AccountUpdateOne) SetPid(s string) *AccountUpdateOne {
	auo.mutation.SetPid(s)
	return auo
}

// SetAccount sets the account field.
func (auo *AccountUpdateOne) SetAccount(s string) *AccountUpdateOne {
	auo.mutation.SetAccount(s)
	return auo
}

// SetAccountTyp sets the account_typ field.
func (auo *AccountUpdateOne) SetAccountTyp(s string) *AccountUpdateOne {
	auo.mutation.SetAccountTyp(s)
	return auo
}

// SetAccountKid sets the account_kid field.
func (auo *AccountUpdateOne) SetAccountKid(s string) *AccountUpdateOne {
	auo.mutation.SetAccountKid(s)
	return auo
}

// SetPassword sets the password field.
func (auo *AccountUpdateOne) SetPassword(s string) *AccountUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetPasswordSalt sets the password_salt field.
func (auo *AccountUpdateOne) SetPasswordSalt(s string) *AccountUpdateOne {
	auo.mutation.SetPasswordSalt(s)
	return auo
}

// SetPasswordType sets the password_type field.
func (auo *AccountUpdateOne) SetPasswordType(s string) *AccountUpdateOne {
	auo.mutation.SetPasswordType(s)
	return auo
}

// SetVerifySecret sets the verify_secret field.
func (auo *AccountUpdateOne) SetVerifySecret(s string) *AccountUpdateOne {
	auo.mutation.SetVerifySecret(s)
	return auo
}

// SetVerifyType sets the verify_type field.
func (auo *AccountUpdateOne) SetVerifyType(s string) *AccountUpdateOne {
	auo.mutation.SetVerifyType(s)
	return auo
}

// SetUserID sets the user_id field.
func (auo *AccountUpdateOne) SetUserID(i int) *AccountUpdateOne {
	auo.mutation.ResetUserID()
	auo.mutation.SetUserID(i)
	return auo
}

// AddUserID adds i to user_id.
func (auo *AccountUpdateOne) AddUserID(i int) *AccountUpdateOne {
	auo.mutation.AddUserID(i)
	return auo
}

// SetRoleID sets the role_id field.
func (auo *AccountUpdateOne) SetRoleID(i int) *AccountUpdateOne {
	auo.mutation.ResetRoleID()
	auo.mutation.SetRoleID(i)
	return auo
}

// AddRoleID adds i to role_id.
func (auo *AccountUpdateOne) AddRoleID(i int) *AccountUpdateOne {
	auo.mutation.AddRoleID(i)
	return auo
}

// SetStatus sets the status field.
func (auo *AccountUpdateOne) SetStatus(i int) *AccountUpdateOne {
	auo.mutation.ResetStatus()
	auo.mutation.SetStatus(i)
	return auo
}

// AddStatus adds i to status.
func (auo *AccountUpdateOne) AddStatus(i int) *AccountUpdateOne {
	auo.mutation.AddStatus(i)
	return auo
}

// SetDescription sets the description field.
func (auo *AccountUpdateOne) SetDescription(s string) *AccountUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetOa2Token sets the oa2_token field.
func (auo *AccountUpdateOne) SetOa2Token(s string) *AccountUpdateOne {
	auo.mutation.SetOa2Token(s)
	return auo
}

// SetOa2Expired sets the oa2_expired field.
func (auo *AccountUpdateOne) SetOa2Expired(t time.Time) *AccountUpdateOne {
	auo.mutation.SetOa2Expired(t)
	return auo
}

// SetOa2Fake sets the oa2_fake field.
func (auo *AccountUpdateOne) SetOa2Fake(s string) *AccountUpdateOne {
	auo.mutation.SetOa2Fake(s)
	return auo
}

// SetOa2Client sets the oa2_client field.
func (auo *AccountUpdateOne) SetOa2Client(i int) *AccountUpdateOne {
	auo.mutation.ResetOa2Client()
	auo.mutation.SetOa2Client(i)
	return auo
}

// AddOa2Client adds i to oa2_client.
func (auo *AccountUpdateOne) AddOa2Client(i int) *AccountUpdateOne {
	auo.mutation.AddOa2Client(i)
	return auo
}

// SetCreator sets the creator field.
func (auo *AccountUpdateOne) SetCreator(s string) *AccountUpdateOne {
	auo.mutation.SetCreator(s)
	return auo
}

// SetCreatedAt sets the created_at field.
func (auo *AccountUpdateOne) SetCreatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableCreatedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the updated_at field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableUpdatedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// SetVersion sets the version field.
func (auo *AccountUpdateOne) SetVersion(i int) *AccountUpdateOne {
	auo.mutation.ResetVersion()
	auo.mutation.SetVersion(i)
	return auo
}

// SetNillableVersion sets the version field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableVersion(i *int) *AccountUpdateOne {
	if i != nil {
		auo.SetVersion(*i)
	}
	return auo
}

// AddVersion adds i to version.
func (auo *AccountUpdateOne) AddVersion(i int) *AccountUpdateOne {
	auo.mutation.AddVersion(i)
	return auo
}

// SetString1 sets the string_1 field.
func (auo *AccountUpdateOne) SetString1(s string) *AccountUpdateOne {
	auo.mutation.SetString1(s)
	return auo
}

// SetString2 sets the string_2 field.
func (auo *AccountUpdateOne) SetString2(s string) *AccountUpdateOne {
	auo.mutation.SetString2(s)
	return auo
}

// SetString3 sets the string_3 field.
func (auo *AccountUpdateOne) SetString3(s string) *AccountUpdateOne {
	auo.mutation.SetString3(s)
	return auo
}

// SetNumber1 sets the number_1 field.
func (auo *AccountUpdateOne) SetNumber1(i int) *AccountUpdateOne {
	auo.mutation.ResetNumber1()
	auo.mutation.SetNumber1(i)
	return auo
}

// AddNumber1 adds i to number_1.
func (auo *AccountUpdateOne) AddNumber1(i int) *AccountUpdateOne {
	auo.mutation.AddNumber1(i)
	return auo
}

// SetNumber2 sets the number_2 field.
func (auo *AccountUpdateOne) SetNumber2(i int) *AccountUpdateOne {
	auo.mutation.ResetNumber2()
	auo.mutation.SetNumber2(i)
	return auo
}

// AddNumber2 adds i to number_2.
func (auo *AccountUpdateOne) AddNumber2(i int) *AccountUpdateOne {
	auo.mutation.AddNumber2(i)
	return auo
}

// SetNumber3 sets the number_3 field.
func (auo *AccountUpdateOne) SetNumber3(i int) *AccountUpdateOne {
	auo.mutation.ResetNumber3()
	auo.mutation.SetNumber3(i)
	return auo
}

// AddNumber3 adds i to number_3.
func (auo *AccountUpdateOne) AddNumber3(i int) *AccountUpdateOne {
	auo.mutation.AddNumber3(i)
	return auo
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// Save executes the query and returns the updated entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (a *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Account.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Pid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPid,
		})
	}
	if value, ok := auo.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAccount,
		})
	}
	if value, ok := auo.mutation.AccountTyp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAccountTyp,
		})
	}
	if value, ok := auo.mutation.AccountKid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAccountKid,
		})
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPassword,
		})
	}
	if value, ok := auo.mutation.PasswordSalt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPasswordSalt,
		})
	}
	if value, ok := auo.mutation.PasswordType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPasswordType,
		})
	}
	if value, ok := auo.mutation.VerifySecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldVerifySecret,
		})
	}
	if value, ok := auo.mutation.VerifyType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldVerifyType,
		})
	}
	if value, ok := auo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldUserID,
		})
	}
	if value, ok := auo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldUserID,
		})
	}
	if value, ok := auo.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldRoleID,
		})
	}
	if value, ok := auo.mutation.AddedRoleID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldRoleID,
		})
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldStatus,
		})
	}
	if value, ok := auo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldStatus,
		})
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldDescription,
		})
	}
	if value, ok := auo.mutation.Oa2Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldOa2Token,
		})
	}
	if value, ok := auo.mutation.Oa2Expired(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldOa2Expired,
		})
	}
	if value, ok := auo.mutation.Oa2Fake(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldOa2Fake,
		})
	}
	if value, ok := auo.mutation.Oa2Client(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldOa2Client,
		})
	}
	if value, ok := auo.mutation.AddedOa2Client(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldOa2Client,
		})
	}
	if value, ok := auo.mutation.Creator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldCreator,
		})
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldVersion,
		})
	}
	if value, ok := auo.mutation.AddedVersion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldVersion,
		})
	}
	if value, ok := auo.mutation.String1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldString1,
		})
	}
	if value, ok := auo.mutation.String2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldString2,
		})
	}
	if value, ok := auo.mutation.String3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldString3,
		})
	}
	if value, ok := auo.mutation.Number1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber1,
		})
	}
	if value, ok := auo.mutation.AddedNumber1(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber1,
		})
	}
	if value, ok := auo.mutation.Number2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber2,
		})
	}
	if value, ok := auo.mutation.AddedNumber2(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber2,
		})
	}
	if value, ok := auo.mutation.Number3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber3,
		})
	}
	if value, ok := auo.mutation.AddedNumber3(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber3,
		})
	}
	a = &Account{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}
