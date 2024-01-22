package sqlmap

import "database/sql"

// MapRows maps the data in the given *sql.Rows to []RowView.
//
// RowView is based on map[string]any.
//
// This code is based on the ideas presented at the following URL:
//   - https://kylewbanks.com/blog/query-result-to-map-in-golang
func MapRows(rows *sql.Rows) ([]RowView, error) {
	var (
		cols []string
		err  error
	)

	cols, err = rows.Columns()
	if err != nil {
		return nil, err
	}

	var (
		result = make([]RowView, 0)
	)

	for rows.Next() {
		var (
			c  = make([]any, len(cols)) // column values
			cp = make([]any, len(cols)) // pointer of column value
		)

		for i := 0; i < len(c); i++ {
			cp[i] = &c[i]
		}

		err = rows.Scan(cp...)
		if err != nil {
			return nil, err
		}

		var (
			m = make(RowView)
		)

		for i, name := range cols {
			v := cp[i].(*any)
			m[name] = *v
		}

		result = append(result, m)
	}

	return result, nil
}
