// Code generated by entc, DO NOT EDIT.

package account

import (
	"time"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldAccountTyp holds the string denoting the account_typ field in the database.
	FieldAccountTyp = "account_typ"
	// FieldAccountKid holds the string denoting the account_kid field in the database.
	FieldAccountKid = "account_kid"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldPasswordSalt holds the string denoting the password_salt field in the database.
	FieldPasswordSalt = "password_salt"
	// FieldPasswordType holds the string denoting the password_type field in the database.
	FieldPasswordType = "password_type"
	// FieldVerifySecret holds the string denoting the verify_secret field in the database.
	FieldVerifySecret = "verify_secret"
	// FieldVerifyType holds the string denoting the verify_type field in the database.
	FieldVerifyType = "verify_type"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldOa2Token holds the string denoting the oa2_token field in the database.
	FieldOa2Token = "oa2_token"
	// FieldOa2Expired holds the string denoting the oa2_expired field in the database.
	FieldOa2Expired = "oa2_expired"
	// FieldOa2Fake holds the string denoting the oa2_fake field in the database.
	FieldOa2Fake = "oa2_fake"
	// FieldOa2Client holds the string denoting the oa2_client field in the database.
	FieldOa2Client = "oa2_client"
	// FieldCreator holds the string denoting the creator field in the database.
	FieldCreator = "creator"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldString1 holds the string denoting the string_1 field in the database.
	FieldString1 = "string_1"
	// FieldString2 holds the string denoting the string_2 field in the database.
	FieldString2 = "string_2"
	// FieldString3 holds the string denoting the string_3 field in the database.
	FieldString3 = "string_3"
	// FieldNumber1 holds the string denoting the number_1 field in the database.
	FieldNumber1 = "number_1"
	// FieldNumber2 holds the string denoting the number_2 field in the database.
	FieldNumber2 = "number_2"
	// FieldNumber3 holds the string denoting the number_3 field in the database.
	FieldNumber3 = "number_3"

	// Table holds the table name of the account in the database.
	Table = "account"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldPid,
	FieldAccount,
	FieldAccountTyp,
	FieldAccountKid,
	FieldPassword,
	FieldPasswordSalt,
	FieldPasswordType,
	FieldVerifySecret,
	FieldVerifyType,
	FieldUserID,
	FieldRoleID,
	FieldStatus,
	FieldDescription,
	FieldOa2Token,
	FieldOa2Expired,
	FieldOa2Fake,
	FieldOa2Client,
	FieldCreator,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVersion,
	FieldString1,
	FieldString2,
	FieldString3,
	FieldNumber1,
	FieldNumber2,
	FieldNumber3,
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// DefaultVersion holds the default value on creation for the version field.
	DefaultVersion int
)
