// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/suisrc/zgo/app/model/ent/account"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
}

// SetPid sets the pid field.
func (ac *AccountCreate) SetPid(s string) *AccountCreate {
	ac.mutation.SetPid(s)
	return ac
}

// SetAccount sets the account field.
func (ac *AccountCreate) SetAccount(s string) *AccountCreate {
	ac.mutation.SetAccount(s)
	return ac
}

// SetAccountTyp sets the account_typ field.
func (ac *AccountCreate) SetAccountTyp(s string) *AccountCreate {
	ac.mutation.SetAccountTyp(s)
	return ac
}

// SetAccountKid sets the account_kid field.
func (ac *AccountCreate) SetAccountKid(s string) *AccountCreate {
	ac.mutation.SetAccountKid(s)
	return ac
}

// SetPassword sets the password field.
func (ac *AccountCreate) SetPassword(s string) *AccountCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetPasswordSalt sets the password_salt field.
func (ac *AccountCreate) SetPasswordSalt(s string) *AccountCreate {
	ac.mutation.SetPasswordSalt(s)
	return ac
}

// SetPasswordType sets the password_type field.
func (ac *AccountCreate) SetPasswordType(s string) *AccountCreate {
	ac.mutation.SetPasswordType(s)
	return ac
}

// SetVerifySecret sets the verify_secret field.
func (ac *AccountCreate) SetVerifySecret(s string) *AccountCreate {
	ac.mutation.SetVerifySecret(s)
	return ac
}

// SetVerifyType sets the verify_type field.
func (ac *AccountCreate) SetVerifyType(s string) *AccountCreate {
	ac.mutation.SetVerifyType(s)
	return ac
}

// SetUserID sets the user_id field.
func (ac *AccountCreate) SetUserID(i int) *AccountCreate {
	ac.mutation.SetUserID(i)
	return ac
}

// SetRoleID sets the role_id field.
func (ac *AccountCreate) SetRoleID(i int) *AccountCreate {
	ac.mutation.SetRoleID(i)
	return ac
}

// SetStatus sets the status field.
func (ac *AccountCreate) SetStatus(i int) *AccountCreate {
	ac.mutation.SetStatus(i)
	return ac
}

// SetDescription sets the description field.
func (ac *AccountCreate) SetDescription(s string) *AccountCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetOa2Token sets the oa2_token field.
func (ac *AccountCreate) SetOa2Token(s string) *AccountCreate {
	ac.mutation.SetOa2Token(s)
	return ac
}

// SetOa2Expired sets the oa2_expired field.
func (ac *AccountCreate) SetOa2Expired(t time.Time) *AccountCreate {
	ac.mutation.SetOa2Expired(t)
	return ac
}

// SetOa2Fake sets the oa2_fake field.
func (ac *AccountCreate) SetOa2Fake(s string) *AccountCreate {
	ac.mutation.SetOa2Fake(s)
	return ac
}

// SetOa2Client sets the oa2_client field.
func (ac *AccountCreate) SetOa2Client(i int) *AccountCreate {
	ac.mutation.SetOa2Client(i)
	return ac
}

// SetCreator sets the creator field.
func (ac *AccountCreate) SetCreator(s string) *AccountCreate {
	ac.mutation.SetCreator(s)
	return ac
}

// SetCreatedAt sets the created_at field.
func (ac *AccountCreate) SetCreatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the updated_at field.
func (ac *AccountCreate) SetUpdatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetVersion sets the version field.
func (ac *AccountCreate) SetVersion(i int) *AccountCreate {
	ac.mutation.SetVersion(i)
	return ac
}

// SetNillableVersion sets the version field if the given value is not nil.
func (ac *AccountCreate) SetNillableVersion(i *int) *AccountCreate {
	if i != nil {
		ac.SetVersion(*i)
	}
	return ac
}

// SetString1 sets the string_1 field.
func (ac *AccountCreate) SetString1(s string) *AccountCreate {
	ac.mutation.SetString1(s)
	return ac
}

// SetString2 sets the string_2 field.
func (ac *AccountCreate) SetString2(s string) *AccountCreate {
	ac.mutation.SetString2(s)
	return ac
}

// SetString3 sets the string_3 field.
func (ac *AccountCreate) SetString3(s string) *AccountCreate {
	ac.mutation.SetString3(s)
	return ac
}

// SetNumber1 sets the number_1 field.
func (ac *AccountCreate) SetNumber1(i int) *AccountCreate {
	ac.mutation.SetNumber1(i)
	return ac
}

// SetNumber2 sets the number_2 field.
func (ac *AccountCreate) SetNumber2(i int) *AccountCreate {
	ac.mutation.SetNumber2(i)
	return ac
}

// SetNumber3 sets the number_3 field.
func (ac *AccountCreate) SetNumber3(i int) *AccountCreate {
	ac.mutation.SetNumber3(i)
	return ac
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	if _, ok := ac.mutation.Pid(); !ok {
		return nil, &ValidationError{Name: "pid", err: errors.New("ent: missing required field \"pid\"")}
	}
	if _, ok := ac.mutation.Account(); !ok {
		return nil, &ValidationError{Name: "account", err: errors.New("ent: missing required field \"account\"")}
	}
	if _, ok := ac.mutation.AccountTyp(); !ok {
		return nil, &ValidationError{Name: "account_typ", err: errors.New("ent: missing required field \"account_typ\"")}
	}
	if _, ok := ac.mutation.AccountKid(); !ok {
		return nil, &ValidationError{Name: "account_kid", err: errors.New("ent: missing required field \"account_kid\"")}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return nil, &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if _, ok := ac.mutation.PasswordSalt(); !ok {
		return nil, &ValidationError{Name: "password_salt", err: errors.New("ent: missing required field \"password_salt\"")}
	}
	if _, ok := ac.mutation.PasswordType(); !ok {
		return nil, &ValidationError{Name: "password_type", err: errors.New("ent: missing required field \"password_type\"")}
	}
	if _, ok := ac.mutation.VerifySecret(); !ok {
		return nil, &ValidationError{Name: "verify_secret", err: errors.New("ent: missing required field \"verify_secret\"")}
	}
	if _, ok := ac.mutation.VerifyType(); !ok {
		return nil, &ValidationError{Name: "verify_type", err: errors.New("ent: missing required field \"verify_type\"")}
	}
	if _, ok := ac.mutation.UserID(); !ok {
		return nil, &ValidationError{Name: "user_id", err: errors.New("ent: missing required field \"user_id\"")}
	}
	if _, ok := ac.mutation.RoleID(); !ok {
		return nil, &ValidationError{Name: "role_id", err: errors.New("ent: missing required field \"role_id\"")}
	}
	if _, ok := ac.mutation.Status(); !ok {
		return nil, &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return nil, &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := ac.mutation.Oa2Token(); !ok {
		return nil, &ValidationError{Name: "oa2_token", err: errors.New("ent: missing required field \"oa2_token\"")}
	}
	if _, ok := ac.mutation.Oa2Expired(); !ok {
		return nil, &ValidationError{Name: "oa2_expired", err: errors.New("ent: missing required field \"oa2_expired\"")}
	}
	if _, ok := ac.mutation.Oa2Fake(); !ok {
		return nil, &ValidationError{Name: "oa2_fake", err: errors.New("ent: missing required field \"oa2_fake\"")}
	}
	if _, ok := ac.mutation.Oa2Client(); !ok {
		return nil, &ValidationError{Name: "oa2_client", err: errors.New("ent: missing required field \"oa2_client\"")}
	}
	if _, ok := ac.mutation.Creator(); !ok {
		return nil, &ValidationError{Name: "creator", err: errors.New("ent: missing required field \"creator\"")}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := account.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := account.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.Version(); !ok {
		v := account.DefaultVersion
		ac.mutation.SetVersion(v)
	}
	if _, ok := ac.mutation.String1(); !ok {
		return nil, &ValidationError{Name: "string_1", err: errors.New("ent: missing required field \"string_1\"")}
	}
	if _, ok := ac.mutation.String2(); !ok {
		return nil, &ValidationError{Name: "string_2", err: errors.New("ent: missing required field \"string_2\"")}
	}
	if _, ok := ac.mutation.String3(); !ok {
		return nil, &ValidationError{Name: "string_3", err: errors.New("ent: missing required field \"string_3\"")}
	}
	if _, ok := ac.mutation.Number1(); !ok {
		return nil, &ValidationError{Name: "number_1", err: errors.New("ent: missing required field \"number_1\"")}
	}
	if _, ok := ac.mutation.Number2(); !ok {
		return nil, &ValidationError{Name: "number_2", err: errors.New("ent: missing required field \"number_2\"")}
	}
	if _, ok := ac.mutation.Number3(); !ok {
		return nil, &ValidationError{Name: "number_3", err: errors.New("ent: missing required field \"number_3\"")}
	}
	var (
		err  error
		node *Account
	)
	if len(ac.hooks) == 0 {
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	a, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	a.ID = int(id)
	return a, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		a     = &Account{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: account.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Pid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPid,
		})
		a.Pid = value
	}
	if value, ok := ac.mutation.Account(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAccount,
		})
		a.Account = value
	}
	if value, ok := ac.mutation.AccountTyp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAccountTyp,
		})
		a.AccountTyp = value
	}
	if value, ok := ac.mutation.AccountKid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAccountKid,
		})
		a.AccountKid = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPassword,
		})
		a.Password = value
	}
	if value, ok := ac.mutation.PasswordSalt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPasswordSalt,
		})
		a.PasswordSalt = value
	}
	if value, ok := ac.mutation.PasswordType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPasswordType,
		})
		a.PasswordType = value
	}
	if value, ok := ac.mutation.VerifySecret(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldVerifySecret,
		})
		a.VerifySecret = value
	}
	if value, ok := ac.mutation.VerifyType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldVerifyType,
		})
		a.VerifyType = value
	}
	if value, ok := ac.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldUserID,
		})
		a.UserID = value
	}
	if value, ok := ac.mutation.RoleID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldRoleID,
		})
		a.RoleID = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldStatus,
		})
		a.Status = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldDescription,
		})
		a.Description = value
	}
	if value, ok := ac.mutation.Oa2Token(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldOa2Token,
		})
		a.Oa2Token = value
	}
	if value, ok := ac.mutation.Oa2Expired(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldOa2Expired,
		})
		a.Oa2Expired = value
	}
	if value, ok := ac.mutation.Oa2Fake(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldOa2Fake,
		})
		a.Oa2Fake = value
	}
	if value, ok := ac.mutation.Oa2Client(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldOa2Client,
		})
		a.Oa2Client = value
	}
	if value, ok := ac.mutation.Creator(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldCreator,
		})
		a.Creator = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
		a.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
		a.UpdatedAt = value
	}
	if value, ok := ac.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldVersion,
		})
		a.Version = value
	}
	if value, ok := ac.mutation.String1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldString1,
		})
		a.String1 = value
	}
	if value, ok := ac.mutation.String2(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldString2,
		})
		a.String2 = value
	}
	if value, ok := ac.mutation.String3(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldString3,
		})
		a.String3 = value
	}
	if value, ok := ac.mutation.Number1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber1,
		})
		a.Number1 = value
	}
	if value, ok := ac.mutation.Number2(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber2,
		})
		a.Number2 = value
	}
	if value, ok := ac.mutation.Number3(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: account.FieldNumber3,
		})
		a.Number3 = value
	}
	return a, _spec
}
