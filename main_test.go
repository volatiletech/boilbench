package main

import (
	"database/sql/driver"
	"os"
	"testing"

	"github.com/volatiletech/boilbench/mimic"
	"gorm.io/driver/postgres"
	"xorm.io/xorm/dialects"
)

func jetQuery() mimic.QueryResult {
	return mimic.QueryResult{
		Query: &mimic.Query{
			Cols: []string{"id", "pilot_id", "airport_id", "name", "color", "uuid", "identifier", "cargo", "manifest"},
			Vals: [][]driver.Value{
				{
					int64(1), int64(1), int64(1), "test", nil, "test", "test", []byte("test"), []byte("test"),
				},
				{
					int64(2), int64(2), int64(2), "test", nil, "test", "test", []byte("test"), []byte("test"),
				},
				{
					int64(3), int64(3), int64(3), "test", nil, "test", "test", []byte("test"), []byte("test"),
				},
				{
					int64(4), int64(4), int64(4), "test", nil, "test", "test", []byte("test"), []byte("test"),
				},
				{
					int64(5), int64(5), int64(5), "test", nil, "test", "test", []byte("test"), []byte("test"),
				},
			},
		},
	}
}

func jetQuerySubset() mimic.QueryResult {
	return mimic.QueryResult{
		Query: &mimic.Query{
			Cols: []string{"id", "name", "color", "uuid", "identifier", "cargo", "manifest"},
			Vals: [][]driver.Value{
				{
					int64(1), "test", nil, "test", "test", []byte("test"), []byte("test"),
				},
				{
					int64(2), "test", nil, "test", "test", []byte("test"), []byte("test"),
				},
				{
					int64(3), "test", nil, "test", "test", []byte("test"), []byte("test"),
				},
				{
					int64(4), "test", nil, "test", "test", []byte("test"), []byte("test"),
				},
				{
					int64(5), "test", nil, "test", "test", []byte("test"), []byte("test"),
				},
			},
		},
	}
}

func pilotQuery() mimic.QueryResult {
	return mimic.QueryResult{
		Query: &mimic.Query{
			Cols: []string{"id", "name"},
			Vals: [][]driver.Value{
				{
					int64(1), "test",
				},
				{
					int64(2), "test",
				},
				{
					int64(3), "test",
				},
				{
					int64(4), "test",
				},
				{
					int64(5), "test",
				},
			},
		},
	}
}

func languageQuery() mimic.QueryResult {
	return mimic.QueryResult{
		Query: &mimic.Query{
			Cols: []string{"id", "name"},
			Vals: [][]driver.Value{
				{
					int64(1), "test",
				},
				{
					int64(2), "test",
				},
				{
					int64(3), "test",
				},
				{
					int64(4), "test",
				},
				{
					int64(5), "test",
				},
				{
					int64(6), "test",
				},
				{
					int64(7), "test",
				},
				{
					int64(8), "test",
				},
				{
					int64(9), "test",
				},
				{
					int64(10), "test",
				},
			},
		},
	}
}

func jetQueryUpdate() mimic.QueryResult {
	return mimic.QueryResult{
		Query: &mimic.Query{
			Cols: []string{"id", "pilot_id", "airport_id", "name", "color", "uuid", "identifier", "cargo", "manifest"},
			Vals: [][]driver.Value{
				{
					int64(1), int64(1), int64(1), "test", nil, "test", "test", []byte("test"), []byte("test"),
				},
			},
		},
	}
}

func jetQueryInsert() mimic.QueryResult {
	return mimic.QueryResult{
		Query: &mimic.Query{
			Cols: []string{"id"},
			Vals: [][]driver.Value{
				{
					int64(1),
				},
			},
		},
	}
}

func jetExec() mimic.QueryResult {
	return mimic.QueryResult{
		Result: &mimic.Result{
			NumRows: 5,
		},
	}
}

func jetExecUpdate() mimic.QueryResult {
	return mimic.QueryResult{
		Result: &mimic.Result{
			NumRows: 1,
		},
	}
}

var gormMimicDialector = postgres.New(postgres.Config{
	DriverName: "mimic",
})

func TestMain(m *testing.M) {
	dialects.RegisterDriver("mimic", &mimic.XormDriver{})
	if dialects.QueryDriver("mimic") == nil {
		panic("failed to register xorm driver")
	}

	code := m.Run()
	os.Exit(code)
}
