// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/suisrc/zgo/app/model/ent/account"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Pid holds the value of the "pid" field.
	Pid string `json:"pid,omitempty"`
	// Account holds the value of the "account" field.
	Account string `json:"account,omitempty"`
	// AccountTyp holds the value of the "account_typ" field.
	AccountTyp string `json:"account_typ,omitempty"`
	// AccountKid holds the value of the "account_kid" field.
	AccountKid string `json:"account_kid,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// PasswordSalt holds the value of the "password_salt" field.
	PasswordSalt string `json:"password_salt,omitempty"`
	// PasswordType holds the value of the "password_type" field.
	PasswordType string `json:"password_type,omitempty"`
	// VerifySecret holds the value of the "verify_secret" field.
	VerifySecret string `json:"verify_secret,omitempty"`
	// VerifyType holds the value of the "verify_type" field.
	VerifyType string `json:"verify_type,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID int `json:"role_id,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Oa2Token holds the value of the "oa2_token" field.
	Oa2Token string `json:"oa2_token,omitempty"`
	// Oa2Expired holds the value of the "oa2_expired" field.
	Oa2Expired time.Time `json:"oa2_expired,omitempty"`
	// Oa2Fake holds the value of the "oa2_fake" field.
	Oa2Fake string `json:"oa2_fake,omitempty"`
	// Oa2Client holds the value of the "oa2_client" field.
	Oa2Client int `json:"oa2_client,omitempty"`
	// Creator holds the value of the "creator" field.
	Creator string `json:"creator,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Version holds the value of the "version" field.
	Version int `json:"version,omitempty"`
	// String1 holds the value of the "string_1" field.
	String1 string `json:"string_1,omitempty"`
	// String2 holds the value of the "string_2" field.
	String2 string `json:"string_2,omitempty"`
	// String3 holds the value of the "string_3" field.
	String3 string `json:"string_3,omitempty"`
	// Number1 holds the value of the "number_1" field.
	Number1 int `json:"number_1,omitempty"`
	// Number2 holds the value of the "number_2" field.
	Number2 int `json:"number_2,omitempty"`
	// Number3 holds the value of the "number_3" field.
	Number3 int `json:"number_3,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // pid
		&sql.NullString{}, // account
		&sql.NullString{}, // account_typ
		&sql.NullString{}, // account_kid
		&sql.NullString{}, // password
		&sql.NullString{}, // password_salt
		&sql.NullString{}, // password_type
		&sql.NullString{}, // verify_secret
		&sql.NullString{}, // verify_type
		&sql.NullInt64{},  // user_id
		&sql.NullInt64{},  // role_id
		&sql.NullInt64{},  // status
		&sql.NullString{}, // description
		&sql.NullString{}, // oa2_token
		&sql.NullTime{},   // oa2_expired
		&sql.NullString{}, // oa2_fake
		&sql.NullInt64{},  // oa2_client
		&sql.NullString{}, // creator
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullInt64{},  // version
		&sql.NullString{}, // string_1
		&sql.NullString{}, // string_2
		&sql.NullString{}, // string_3
		&sql.NullInt64{},  // number_1
		&sql.NullInt64{},  // number_2
		&sql.NullInt64{},  // number_3
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(values ...interface{}) error {
	if m, n := len(values), len(account.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field pid", values[0])
	} else if value.Valid {
		a.Pid = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field account", values[1])
	} else if value.Valid {
		a.Account = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field account_typ", values[2])
	} else if value.Valid {
		a.AccountTyp = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field account_kid", values[3])
	} else if value.Valid {
		a.AccountKid = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[4])
	} else if value.Valid {
		a.Password = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password_salt", values[5])
	} else if value.Valid {
		a.PasswordSalt = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password_type", values[6])
	} else if value.Valid {
		a.PasswordType = value.String
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field verify_secret", values[7])
	} else if value.Valid {
		a.VerifySecret = value.String
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field verify_type", values[8])
	} else if value.Valid {
		a.VerifyType = value.String
	}
	if value, ok := values[9].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field user_id", values[9])
	} else if value.Valid {
		a.UserID = int(value.Int64)
	}
	if value, ok := values[10].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field role_id", values[10])
	} else if value.Valid {
		a.RoleID = int(value.Int64)
	}
	if value, ok := values[11].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[11])
	} else if value.Valid {
		a.Status = int(value.Int64)
	}
	if value, ok := values[12].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[12])
	} else if value.Valid {
		a.Description = value.String
	}
	if value, ok := values[13].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field oa2_token", values[13])
	} else if value.Valid {
		a.Oa2Token = value.String
	}
	if value, ok := values[14].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field oa2_expired", values[14])
	} else if value.Valid {
		a.Oa2Expired = value.Time
	}
	if value, ok := values[15].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field oa2_fake", values[15])
	} else if value.Valid {
		a.Oa2Fake = value.String
	}
	if value, ok := values[16].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field oa2_client", values[16])
	} else if value.Valid {
		a.Oa2Client = int(value.Int64)
	}
	if value, ok := values[17].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field creator", values[17])
	} else if value.Valid {
		a.Creator = value.String
	}
	if value, ok := values[18].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[18])
	} else if value.Valid {
		a.CreatedAt = value.Time
	}
	if value, ok := values[19].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[19])
	} else if value.Valid {
		a.UpdatedAt = value.Time
	}
	if value, ok := values[20].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field version", values[20])
	} else if value.Valid {
		a.Version = int(value.Int64)
	}
	if value, ok := values[21].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field string_1", values[21])
	} else if value.Valid {
		a.String1 = value.String
	}
	if value, ok := values[22].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field string_2", values[22])
	} else if value.Valid {
		a.String2 = value.String
	}
	if value, ok := values[23].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field string_3", values[23])
	} else if value.Valid {
		a.String3 = value.String
	}
	if value, ok := values[24].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field number_1", values[24])
	} else if value.Valid {
		a.Number1 = int(value.Int64)
	}
	if value, ok := values[25].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field number_2", values[25])
	} else if value.Valid {
		a.Number2 = int(value.Int64)
	}
	if value, ok := values[26].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field number_3", values[26])
	} else if value.Valid {
		a.Number3 = int(value.Int64)
	}
	return nil
}

// Update returns a builder for updating this Account.
// Note that, you need to call Account.Unwrap() before calling this method, if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", pid=")
	builder.WriteString(a.Pid)
	builder.WriteString(", account=")
	builder.WriteString(a.Account)
	builder.WriteString(", account_typ=")
	builder.WriteString(a.AccountTyp)
	builder.WriteString(", account_kid=")
	builder.WriteString(a.AccountKid)
	builder.WriteString(", password=")
	builder.WriteString(a.Password)
	builder.WriteString(", password_salt=")
	builder.WriteString(a.PasswordSalt)
	builder.WriteString(", password_type=")
	builder.WriteString(a.PasswordType)
	builder.WriteString(", verify_secret=")
	builder.WriteString(a.VerifySecret)
	builder.WriteString(", verify_type=")
	builder.WriteString(a.VerifyType)
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", a.UserID))
	builder.WriteString(", role_id=")
	builder.WriteString(fmt.Sprintf("%v", a.RoleID))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", a.Status))
	builder.WriteString(", description=")
	builder.WriteString(a.Description)
	builder.WriteString(", oa2_token=")
	builder.WriteString(a.Oa2Token)
	builder.WriteString(", oa2_expired=")
	builder.WriteString(a.Oa2Expired.Format(time.ANSIC))
	builder.WriteString(", oa2_fake=")
	builder.WriteString(a.Oa2Fake)
	builder.WriteString(", oa2_client=")
	builder.WriteString(fmt.Sprintf("%v", a.Oa2Client))
	builder.WriteString(", creator=")
	builder.WriteString(a.Creator)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", version=")
	builder.WriteString(fmt.Sprintf("%v", a.Version))
	builder.WriteString(", string_1=")
	builder.WriteString(a.String1)
	builder.WriteString(", string_2=")
	builder.WriteString(a.String2)
	builder.WriteString(", string_3=")
	builder.WriteString(a.String3)
	builder.WriteString(", number_1=")
	builder.WriteString(fmt.Sprintf("%v", a.Number1))
	builder.WriteString(", number_2=")
	builder.WriteString(fmt.Sprintf("%v", a.Number2))
	builder.WriteString(", number_3=")
	builder.WriteString(fmt.Sprintf("%v", a.Number3))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
