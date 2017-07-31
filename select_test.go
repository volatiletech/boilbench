package main

import (
	"database/sql"
	"database/sql/driver"
	"testing"

	"gopkg.in/gorp.v1"
	"gopkg.in/src-d/go-kallax.v1"

	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	"github.com/volatiletech/boilbench/gorms"
	"github.com/volatiletech/boilbench/gorps"
	"github.com/volatiletech/boilbench/kallaxes"
	"github.com/volatiletech/boilbench/mimic"
	"github.com/volatiletech/boilbench/models"
	"github.com/volatiletech/boilbench/xorms"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func BenchmarkGORMSelectAll(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	gormdb, err := gorm.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("gorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []gorms.Jet
			err := gormdb.Find(&store).Error
			if err != nil {
				b.Fatal(err)
			}
			store = nil
		}
	})
}

func BenchmarkGORPSelectAll(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	gorpdb := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	if err != nil {
		panic(err)
	}

	b.Run("gorp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []gorps.Jet
			_, err = gorpdb.Select(&store, "select * from jets")
			if err != nil {
				b.Fatal(err)
			}
			store = nil
		}
	})
}

func BenchmarkXORMSelectAll(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	xormdb, err := xorm.NewEngine("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("xorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []xorms.Jet
			err = xormdb.Find(&store)
			if err != nil {
				b.Fatal(err)
			}
			store = nil
		}
	})
}

func BenchmarkKallaxSelectAll(b *testing.B) {
	query := jetQuery()
	query.Vals = [][]driver.Value{
		[]driver.Value{
			int64(1), int64(1), int64(1), "test", nil, "test", "test", []byte("{5}"), []byte("{3}"),
		},
	}
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	jetStore := kallaxes.NewJetStore(db)

	b.Run("kallax", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rs, err := jetStore.Find(kallaxes.NewJetQuery())
			if err != nil {
				b.Fatal(err)
			}
			_, err = rs.All()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkBoilSelectAll(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("boil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err = models.Jets(db).All()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGORMSelectSubset(b *testing.B) {
	var store []gorms.Jet
	query := jetQuery()
	mimic.NewQuery(query)

	gormdb, err := gorm.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("gorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err = gormdb.Select("id, name, color, uuid, identifier, cargo, manifest").Find(&store).Error
			if err != nil {
				b.Fatal(err)
			}
			store = nil
		}
	})
}

func BenchmarkGORPSelectSubset(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	gorpdb := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	if err != nil {
		panic(err)
	}

	b.Run("gorp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []gorps.Jet
			_, err = gorpdb.Select(&store, `select id, name, color, uuid, identifier, cargo, manifest from "jets"`)
			if err != nil {
				b.Fatal(err)
			}
			store = nil
		}
	})
}

func BenchmarkXORMSelectSubset(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	xormdb, err := xorm.NewEngine("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("xorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []xorms.Jet
			err = xormdb.Select("id, name, color, uuid, identifier, cargo, manifest").Find(&store)
			if err != nil {
				b.Fatal(err)
			}
			store = nil
		}
	})
}

func BenchmarkKallaxSelectSubset(b *testing.B) {
	query := jetQuery()
	query.Cols = []string{"id", "name", "color", "uuid", "identifier", "cargo", "manifest"}
	query.Vals = [][]driver.Value{
		[]driver.Value{
			int64(1), int64(1), int64(1), "test", "str", nil, "{3}",
		},
	}
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	jetStore := kallaxes.NewJetStore(db)

	b.Run("kallax", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rs, err := jetStore.Find(
				kallaxes.NewJetQuery().Select(
					kallaxes.Schema.Jet.ID,
					kallaxes.Schema.Jet.Name,
					kallaxes.Schema.Jet.Color,
					kallaxes.Schema.Jet.UUID,
					kallaxes.Schema.Jet.Identifier,
					kallaxes.Schema.Jet.Cargo,
					kallaxes.Schema.Jet.Manifest,
				))
			if err != nil {
				b.Fatal(err)
			}
			_, err = rs.All()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkBoilSelectSubset(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("boil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err = models.Jets(db, qm.Select("id, name, color, uuid, identifier, cargo, manifest")).All()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGORMSelectComplex(b *testing.B) {
	query := jetQuery()
	query.NumInput = -1
	mimic.NewQuery(query)

	gormdb, err := gorm.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("gorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []gorms.Jet
			err = gormdb.Where("id > ?", 1).
				Where("name <> ?", "thing").
				Limit(1).
				Group("id").
				Offset(1).
				Select("id, name, color, uuid, identifier, cargo, manifest").
				Find(&store).Error
			if err != nil {
				b.Fatal(err)
			}
			store = nil
		}
	})
}

func BenchmarkGORPSelectComplex(b *testing.B) {
	query := jetQuery()
	query.NumInput = -1
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	gorpdb := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	if err != nil {
		panic(err)
	}

	b.Run("gorp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []gorps.Jet
			_, err = gorpdb.Select(&store, `
			select id, name, color, uuid, identifier, cargo, manifest from "jets"
			where id > $1 and name <> $2 group by "id" offset $3 limit $4
		`, 1, "thing", 1, 1)
			if err != nil {
				b.Fatal(err)
			}
			store = nil
		}
	})
}

func BenchmarkXORMSelectComplex(b *testing.B) {
	query := jetQuery()
	query.NumInput = -1
	mimic.NewQuery(query)

	xormdb, err := xorm.NewEngine("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("xorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []xorms.Jet
			err = xormdb.
				Select("id, name, color, uuid, identifier, cargo, manifest").
				Where("id > ?", 1).
				Where("name <> ?", "thing").
				Limit(1, 1).
				GroupBy("id").
				Find(&store)
			if err != nil {
				b.Fatal(err)
			}
			store = nil
		}
	})
}

func BenchmarkKallaxSelectComplex(b *testing.B) {
	query := jetQuery()
	query.NumInput = 2
	query.Cols = []string{"id", "name", "color", "uuid", "identifier", "cargo", "manifest"}
	query.Vals = [][]driver.Value{
		[]driver.Value{
			int64(1), int64(1), int64(1), "test", "str", nil, "{3}",
		},
	}
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	jetStore := kallaxes.NewJetStore(db)

	b.Run("kallax", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rs, err := jetStore.Find(
				// Could not find GroupBy for this query
				kallaxes.NewJetQuery().Select(
					kallaxes.Schema.Jet.ID,
					kallaxes.Schema.Jet.Name,
					kallaxes.Schema.Jet.Color,
					kallaxes.Schema.Jet.UUID,
					kallaxes.Schema.Jet.Identifier,
					kallaxes.Schema.Jet.Cargo,
					kallaxes.Schema.Jet.Manifest,
				).
					Where(kallax.Gt(kallaxes.Schema.Jet.ID, 1)).
					Where(kallax.Not(kallax.Eq(kallaxes.Schema.Jet.Name, "thing"))).
					Limit(1).
					Offset(1),
			)
			if err != nil {
				b.Fatal(err)
			}
			_, err = rs.All()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkBoilSelectComplex(b *testing.B) {
	query := jetQuery()
	query.NumInput = -1
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("boil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err = models.Jets(
				db,
				qm.Select("id, name, color, uuid, identifier, cargo, manifest"),
				qm.Where("id > ?", 1),
				qm.And("name <> ?", "thing"),
				qm.Limit(1),
				qm.GroupBy("id"),
				qm.Offset(1),
			).All()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
