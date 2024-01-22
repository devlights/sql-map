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
    mapRows []RowView
)

mapRows, err := sqlmap.MapRows(rows)
if err != nil {
    return err
}

fmt.Println(mapRows)

for _, v := range m {
    name, err := v.Get("Name")
    if err != nil {
        t.Error(err)
    }

    fmt.Printf("[row] %v\n", name)
}
```
