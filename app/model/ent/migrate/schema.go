// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// AccountColumns holds the columns for the "account" table.
	AccountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "pid", Type: field.TypeString},
		{Name: "account", Type: field.TypeString},
		{Name: "account_typ", Type: field.TypeString},
		{Name: "account_kid", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "password_salt", Type: field.TypeString},
		{Name: "password_type", Type: field.TypeString},
		{Name: "verify_secret", Type: field.TypeString},
		{Name: "verify_type", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
		{Name: "status", Type: field.TypeInt},
		{Name: "description", Type: field.TypeString},
		{Name: "oa2_token", Type: field.TypeString},
		{Name: "oa2_expired", Type: field.TypeTime},
		{Name: "oa2_fake", Type: field.TypeString},
		{Name: "oa2_client", Type: field.TypeInt},
		{Name: "creator", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "version", Type: field.TypeInt, Default: 1},
		{Name: "string_1", Type: field.TypeString},
		{Name: "string_2", Type: field.TypeString},
		{Name: "string_3", Type: field.TypeString},
		{Name: "number_1", Type: field.TypeInt},
		{Name: "number_2", Type: field.TypeInt},
		{Name: "number_3", Type: field.TypeInt},
	}
	// AccountTable holds the schema information for the "account" table.
	AccountTable = &schema.Table{
		Name:        "account",
		Columns:     AccountColumns,
		PrimaryKey:  []*schema.Column{AccountColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountTable,
	}
)

func init() {
}
