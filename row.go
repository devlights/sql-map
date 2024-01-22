package sqlmap

import "errors"

var (
	// ErrNotFound is raised when the specified column name does not exist.
	ErrNotFound = errors.New("column name is not found")
)

// RowView represents a row-view in the database.
type RowView map[string]any

// Get retrieves the value of a given column name.
func (me RowView) Get(col string) (any, error) {
	v, ok := me[col]
	if !ok {
		return nil, ErrNotFound
	}

	return v, nil
}
