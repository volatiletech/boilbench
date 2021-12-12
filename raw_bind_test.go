package main

import (
	"context"
	"database/sql"
	"github.com/gobuffalo/pop/v6"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/boilbench/gorms"
	"github.com/volatiletech/boilbench/gorps"
	"github.com/volatiletech/boilbench/mimic"
	"github.com/volatiletech/boilbench/models"
	"github.com/volatiletech/boilbench/xorms"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"gopkg.in/gorp.v1"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

func BenchmarkGORMRawBind(b *testing.B) {
	query := jetQuery()
	mimic.NewQuery(query)

	gormdb, err := gorm.Open(gormMimicDialector, &gorm.Config{})
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
	b.Skip("broken")

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
			err = queries.Raw("select * from jets").Bind(context.Background(), db, &slice)
			if err != nil {
				b.Fatal(err)
			}
			slice = nil
		}
	})
}

func BenchmarkPopRawBind(b *testing.B) {
	dsn := "postgres://BenchmarkPopRawBind"
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
			var slice []models.Jet
			err = popdb.RawQuery("select * from jets").All(&slice)
			if err != nil {
				b.Fatal(err)
			}
			slice = nil
		}
	})
}
