# sql-map

A small library that maps the results of Go's *sql.Rows to map[string]any.

## Install

```sh
go get github.com/devlights/sql-map@latest
```

## Usage

```go
db, _ := sql.Open("sqlite", "path/to/db")
defer db.Close()

rows, _ := db.Query("SELECT * FROM xxxxx")
defer rows.Close()

mapRows, err := sqlmap.MapRows(rows)
if err != nil {
    return err
}

fmt.Println(mapRows)
```