package main

import (
	"database/sql"
	"testing"

	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/vattle/boilbench/gorms"
	"github.com/vattle/boilbench/gorps"
	"github.com/vattle/boilbench/mimic"
	"github.com/vattle/boilbench/models"
	"github.com/vattle/boilbench/xorms"
	"github.com/vattle/sqlboiler/boil"
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
			boil.SQL(db, "select * from jets").Bind(&slice)
			if err != nil {
				b.Fatal(err)
			}
			slice = nil
		}
	})
}
