# sqlmap

A small library that maps the results of ```*sql.Rows``` to ```[]map[string]any```.

## Install

```sh
go get github.com/devlights/sqlmap@latest
```

## Usage

```go
var (
    db  *sql.DB
    err error
)

db, _ := sql.Open("sqlite", "path/to/db")
defer db.Close()

var (
    rows *sql.Rows
)

rows, _ := db.Query("SELECT * FROM xxxxx")
defer rows.Close()

var (
    mapRows []map[string]any
)

mapRows, err := sqlmap.MapRows(rows)
if err != nil {
    return err
}

fmt.Println(mapRows)
```
