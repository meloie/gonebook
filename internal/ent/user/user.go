// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// EdgeContacts holds the string denoting the contacts edge name in mutations.
	EdgeContacts = "contacts"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"

	// Table holds the table name of the user in the database.
	Table = "users"
	// ContactsTable is the table the holds the contacts relation/edge.
	ContactsTable = "contacts"
	// ContactsInverseTable is the table name for the Contact entity.
	// It exists in this package in order to avoid circular dependency with the "contact" package.
	ContactsInverseTable = "contacts"
	// ContactsColumn is the table column denoting the contacts relation/edge.
	ContactsColumn = "contact_owner"
	// TokenTable is the table the holds the token relation/edge.
	TokenTable = "tokens"
	// TokenInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokenInverseTable = "tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "token_user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
