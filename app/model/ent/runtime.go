// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/suisrc/zgo/app/model/ent/account"
	"github.com/suisrc/zgo/app/model/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[18].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[19].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// accountDescVersion is the schema descriptor for version field.
	accountDescVersion := accountFields[20].Descriptor()
	// account.DefaultVersion holds the default value on creation for the version field.
	account.DefaultVersion = accountDescVersion.Default.(int)
}
