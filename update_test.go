package main

import (
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	"github.com/vattle/boilbench/gorms"
	"github.com/vattle/boilbench/gorps"
	"github.com/vattle/boilbench/kallaxes"
	"github.com/vattle/boilbench/mimic"
	"github.com/vattle/boilbench/models"
	"github.com/vattle/boilbench/xorms"
	gorp "gopkg.in/gorp.v1"
)

func BenchmarkGORMUpdate(b *testing.B) {
	store := gorms.Jet{
		ID: 1,
	}

	exec := jetExecUpdate()
	exec.NumInput = -1
	mimic.NewResult(exec)

	gormdb, err := gorm.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("gorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := gormdb.Model(&store).Updates(store).Error
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGORPUpdate(b *testing.B) {
	store := gorps.Jet{
		ID: 1,
	}

	exec := jetExecUpdate()
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
			_, err := gorpdb.Update(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkXORMUpdate(b *testing.B) {
	store := xorms.Jet{
		Id: 1,
	}

	exec := jetExecUpdate()
	exec.NumInput = -1
	mimic.NewResult(exec)

	xormdb, err := xorm.NewEngine("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("xorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := xormdb.Id(store.Id).Update(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkKallaxUpdate(b *testing.B) {
	query := jetQuery()
	query.NumInput = -1
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
	store, err := jetStore.FindOne(kallaxes.NewJetQuery())
	if err != nil {
		b.Fatal(err)
	}

	exec := jetExecUpdate()
	exec.NumInput = -1
	mimic.NewResult(exec)

	db, err = sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}
	jetStore = kallaxes.NewJetStore(db)

	b.Run("kallax", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := jetStore.Update(store, kallaxes.Schema.Jet.Columns()...)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkBoilUpdate(b *testing.B) {
	store := models.Jet{
		ID: 1,
	}

	exec := jetExecUpdate()
	exec.NumInput = -1
	mimic.NewResult(exec)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("boil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := store.Update(db)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
