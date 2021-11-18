// Code generated by entc, DO NOT EDIT.

package allmethodsservice

const (
	// Label holds the string label denoting the allmethodsservice type in the database.
	Label = "all_methods_service"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the allmethodsservice in the database.
	Table = "all_methods_services"
)

// Columns holds all SQL columns for allmethodsservice fields.
var Columns = []string{
	FieldID,
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