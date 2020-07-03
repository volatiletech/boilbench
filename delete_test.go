package main

import (
	"database/sql"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/volatiletech/boilbench/gorms"
	"github.com/volatiletech/boilbench/gorps"
	"github.com/volatiletech/boilbench/kallaxes"
	"github.com/volatiletech/boilbench/mimic"
	"github.com/volatiletech/boilbench/models"
	"github.com/volatiletech/boilbench/xorms"
	"gopkg.in/gorp.v1"
	"xorm.io/xorm"
)

func BenchmarkGORMDelete(b *testing.B) {
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
			err := gormdb.Delete(&store).Error
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGORPDelete(b *testing.B) {
	store := gorps.Jet{
		ID: 1,
	}

	exec := jetExec()
	exec.NumInput = -1
	mimic.NewResult(exec)

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
			_, err := gorpdb.Delete(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkXORMDelete(b *testing.B) {
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
			_, err := xormdb.Delete(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkKallaxDelete(b *testing.B) {
	store := kallaxes.Jet{
		ID: 1,
	}

	exec := jetExec()
	exec.NumInput = -1
	mimic.NewResult(exec)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	jetStore := kallaxes.NewJetStore(db)

	b.Run("kallax", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := jetStore.Delete(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkBoilDelete(b *testing.B) {
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
			_, err := store.Delete(db)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
