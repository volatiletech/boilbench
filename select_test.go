package main

import (
	"database/sql"
	"testing"

	"gopkg.in/gorp.v1"

	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/nullbio/boilbench/gorms"
	"github.com/nullbio/boilbench/gorps"
	"github.com/nullbio/boilbench/mimic"
	"github.com/nullbio/boilbench/models"
	"github.com/nullbio/boilbench/xorms"
	"github.com/vattle/sqlboiler/boil/qm"
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

func BenchmarkSQLXSelectAll(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	db, err := sqlx.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("sqlx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var slice []models.Jet
			err = db.Select(&slice, "select * from jets")
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

func BenchmarkSQLXSelectSubset(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	db, err := sqlx.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("sqlx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var slice []models.Jet
			err = db.Select(&slice, "select id, name, color, uuid, identifier, cargo, manifest")
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

func BenchmarkSQLXSelectComplex(b *testing.B) {
	query := jetQuery()
	query.NumInput = 2
	mimic.NewQuery(query)

	db, err := sqlx.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("sqlx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var slice []models.Jet
			err = db.Select(&slice, "select id, name, color, uuid, identifier, cargo, manifest where id > ? and name <> ? limit 1 group by id offset 1", 1, "thing")
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkBoilSelectAllNormalSlice(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("boil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var jet []models.Jet
			err = models.Jets(db).Bind(&jet)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
