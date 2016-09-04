package mimic

import (
	"database/sql"
	"database/sql/driver"
	"testing"
)

func TestIt(t *testing.T) {
	NewQuery(QueryResult{
		Query: &Query{
			Cols: []string{"id"},
			Vals: [][]driver.Value{
				[]driver.Value{
					int64(56),
				},
				[]driver.Value{
					int64(76),
				},
				[]driver.Value{
					int64(333),
				},
			},
		},
	})

	db, err := sql.Open("mimic", "")
	if err != nil {
		t.Fatal(err)
	}

	rows, err := db.Query("LOL")
	if err != nil {
		t.Fatal(err)
	}

	for rows.Next() {
		var thing int

		if err = rows.Scan(&thing); err != nil {
			t.Fatal(err)
		}
	}

	if err = rows.Err(); err != nil {
		t.Fatal(err)
	}
}
