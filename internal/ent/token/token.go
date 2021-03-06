// Code generated by entc, DO NOT EDIT.

package token

const (
	// Label holds the string label denoting the token type in the database.
	Label = "token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the token in the database.
	Table = "tokens"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "tokens"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "token_user"
)

// Columns holds all SQL columns for token fields.
var Columns = []string{
	FieldID,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Token type.
var ForeignKeys = []string{
	"token_user",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ValueValidator is a validator for the "value" field. It is called by the builders before save.
	ValueValidator func(string) error
)
