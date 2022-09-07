package main

import (
	"context"
	"database/sql"
	"testing"

	"github.com/gobuffalo/pop/v6"
	"github.com/olachat/gola/coredb"

	golas "github.com/volatiletech/boilbench/gola"
	"github.com/volatiletech/boilbench/gorms"
	"github.com/volatiletech/boilbench/gorps"
	"github.com/volatiletech/boilbench/mimic"
	"github.com/volatiletech/boilbench/models"
	"github.com/volatiletech/boilbench/xorms"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gopkg.in/gorp.v1"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

func BenchmarkGORMSelectAll(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	gormdb, err := gorm.Open(gormMimicDialector, &gorm.Config{})
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

func BenchmarkGOLASelectAll(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	coredb.Setup(func(_ string, _ coredb.DBMode) *sql.DB {
		return db
	})

	b.Run("gola", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			store := golas.Select().All()
			if len(store) != 5 {
				b.Fatal("gola load failed: ")
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
		ctx := context.Background()
		for i := 0; i < b.N; i++ {
			_, err = models.Jets().All(ctx, db)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPopSelectAll(b *testing.B) {
	dsn := "postgres://BenchmarkPopSelectAll"

	query := jetQuery()
	mimic.NewQueryDSN(dsn, query)

	popdb, err := pop.NewConnection(&pop.ConnectionDetails{Driver: "mimic", Dialect: "postgres", URL: dsn})
	if err != nil {
		panic(err)
	}

	err = popdb.Open()
	if err != nil {
		panic(err)
	}

	b.Run("pop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []gorps.Jet
			err = popdb.All(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGORMSelectSubset(b *testing.B) {
	var store []gorms.Jet
	query := jetQuerySubset()
	mimic.NewQuery(query)

	gormdb, err := gorm.Open(gormMimicDialector, &gorm.Config{})
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

func BenchmarkGOLASelectSubset(b *testing.B) {
	query := jetQuerySubset()
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	coredb.Setup(func(_ string, _ coredb.DBMode) *sql.DB {
		return db
	})

	b.Run("gola", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			store := golas.SelectFields[struct {
				golas.Id
				golas.Name
				golas.Color
				golas.Uuid
				golas.Identifier
				golas.Cargo
				golas.Manifest
			}]().All()
			if len(store) != 5 {
				b.Fatal("gola load failed: ")
			}
			store = nil
		}
	})
}

func BenchmarkGORPSelectSubset(b *testing.B) {
	query := jetQuerySubset()
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
	query := jetQuerySubset()
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
	query := jetQuerySubset()
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("boil", func(b *testing.B) {
		ctx := context.Background()
		for i := 0; i < b.N; i++ {
			_, err = models.Jets(qm.Select("id, name, color, uuid, identifier, cargo, manifest")).All(ctx, db)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPopSelectSubset(b *testing.B) {
	dsn := "postgres://BenchmarkPopSelectSubset"

	query := jetQuerySubset()
	mimic.NewQueryDSN(dsn, query)

	popdb, err := pop.NewConnection(&pop.ConnectionDetails{Driver: "mimic", Dialect: "postgres", URL: dsn})
	if err != nil {
		panic(err)
	}

	err = popdb.Open()
	if err != nil {
		panic(err)
	}

	b.Run("pop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []gorps.Jet
			err = popdb.Select("id, name, color, uuid, identifier, cargo, manifest").All(&store)
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

	gormdb, err := gorm.Open(gormMimicDialector, &gorm.Config{})
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

func BenchmarkGOLASelectComplex(b *testing.B) {
	query := jetQuerySubset()
	query.NumInput = -1
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	coredb.Setup(func(_ string, _ coredb.DBMode) *sql.DB {
		return db
	})

	type miniJet struct {
		golas.Id
		golas.Name
		golas.Color
		golas.Uuid
		golas.Identifier
		golas.Cargo
		golas.Manifest
	}

	b.Run("gola", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			store, err := golas.FindFields[miniJet]("where id > 1 and name <> ? group by id limit ?, ?", 1, "thing", 1, 1)
			if store == nil {
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
		ctx := context.Background()
		for i := 0; i < b.N; i++ {
			_, err = models.Jets(
				qm.Select("id, name, color, uuid, identifier, cargo, manifest"),
				qm.Where("id > ?", 1),
				qm.And("name <> ?", "thing"),
				qm.Limit(1),
				qm.GroupBy("id"),
				qm.Offset(1),
			).All(ctx, db)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPopSelectComplex(b *testing.B) {
	dsn := "postgres://BenchmarkPopSelectComplex"

	query := jetQuery()
	query.NumInput = -1
	mimic.NewQueryDSN(dsn, query)

	popdb, err := pop.NewConnection(&pop.ConnectionDetails{Driver: "mimic", Dialect: "postgres", URL: dsn})
	if err != nil {
		panic(err)
	}

	err = popdb.Open()
	if err != nil {
		panic(err)
	}

	b.Run("pop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []gorps.Jet
			err = popdb.Select(
				"id, name, color, uuid, identifier, cargo, manifest").
				Where("id > ? AND name <> ?", 1, "thing").
				Limit(1).
				GroupBy("id").
				// Offset does not exist, you need to use a paginator.
				// The Paginator however does another query, which breaks
				// this integration here, so I have removed it.
				//
				//
				// Paginate(1, 1).
				All(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
