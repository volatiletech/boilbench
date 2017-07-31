package main

import (
	"database/sql"
	"testing"

	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/boilbench/gorms"
	"github.com/volatiletech/boilbench/gorps"
	"github.com/volatiletech/boilbench/kallaxes"
	"github.com/volatiletech/boilbench/mimic"
	"github.com/volatiletech/boilbench/models"
	"github.com/volatiletech/boilbench/xorms"
	"github.com/volatiletech/sqlboiler/queries"
	"gopkg.in/gorp.v1"
)

func BenchmarkGORMRawBind(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	gormdb, err := gorm.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("gorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []gorms.Jet
			err := gormdb.Raw("select * from jets").Scan(&store).Error
			if err != nil {
				b.Fatal(err)
			}
			store = nil
		}
	})
}

func BenchmarkGORPRawBind(b *testing.B) {
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

func BenchmarkXORMRawBind(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	xormdb, err := xorm.NewEngine("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("xorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []xorms.Jet
			err = xormdb.SQL("select * from jets").Find(&store)
			if err != nil {
				b.Fatal(err)
			}
			store = nil
		}
	})
}

func BenchmarkKallaxRawBind(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	jetStore := kallaxes.NewJetStore(db)

	b.Run("kallax", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var store []kallaxes.Jet
			rs, err := jetStore.RawQuery("select * from jets")
			if err != nil {
				b.Fatal(err)
			}

			// This kind of discounts the benchmark, since this is not an ORM
			// type wrap-function. It simply tunnels through to the annoying
			// API that is database/sql.
			for rs.Next() {
				var jet kallaxes.Jet
				err = rs.RawScan(
					&jet.ID,
					&jet.PilotID,
					&jet.AirportID,
					&jet.Name,
					&jet.Color,
					&jet.UUID,
					&jet.Identifier,
					&jet.Cargo,
					&jet.Manifest,
				)
				if err != nil {
					b.Fatal(err)
				}
				store = append(store, jet)
			}

			store = nil
		}
	})
}

func BenchmarkSQLXRawBind(b *testing.B) {
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
			slice = nil
		}
	})
}

func BenchmarkBoilRawBind(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("boil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var slice []models.Jet
			err = queries.Raw(db, "select * from jets").Bind(&slice)
			if err != nil {
				b.Fatal(err)
			}
			slice = nil
		}
	})
}
