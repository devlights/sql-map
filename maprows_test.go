package sqlmap

import (
	"database/sql"
	"testing"

	_ "github.com/glebarez/go-sqlite"
)

func TestMapRows(t *testing.T) {
	//
	// Arrange
	//
	db, err := sql.Open("sqlite", "testdb/chinook.db")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM artists ORDER BY ArtistId DESC LIMIT 5")
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()

	//
	// Act
	//
	m, err := MapRows(rows)

	//
	// Assert
	//
	if err != nil {
		t.Error(err)
	}

	if m == nil {
		t.Errorf("[want] not nil\t[got] nil")
	}

	for _, v := range m {
		t.Logf("[row] %v", v)
	}
}
