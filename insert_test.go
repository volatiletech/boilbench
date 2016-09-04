package main

import (
	"database/sql"
	"testing"

	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	"github.com/nullbio/boilbench/gorms"
	"github.com/nullbio/boilbench/gorps"
	"github.com/nullbio/boilbench/mimic"
	"github.com/nullbio/boilbench/models"
	"github.com/nullbio/boilbench/xorms"
	"gopkg.in/gorp.v1"
)

func BenchmarkGORMInsert(b *testing.B) {
	store := gorms.Jet{
		ID: 1,
	}

	exec := jetExec()
	exec.NumInput = -1
	mimic.NewResult(exec)

	gormdb, err := gorm.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("gorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := gormdb.Create(&store).Error
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGORPInsert(b *testing.B) {
	store := gorps.Jet{
		ID: 1,
	}

	query := jetQueryInsert()
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

	gorpdb.AddTable(gorps.Jet{}).SetKeys(true, "ID")

	b.Run("gorp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := gorpdb.Insert(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkXORMInsert(b *testing.B) {
	store := xorms.Jet{
		Id: 1,
	}

	exec := jetExec()
	exec.NumInput = -1
	mimic.NewResult(exec)

	xormdb, err := xorm.NewEngine("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("xorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := xormdb.Insert(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkBoilInsert(b *testing.B) {
	store := models.Jet{
		ID: 1,
	}

	exec := jetExec()
	exec.NumInput = -1
	mimic.NewResult(exec)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("boil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := store.Insert(db)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
