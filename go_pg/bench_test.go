package gopg

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	boilmodels "github.com/volatiletech/boilbench/go_pg/boilgenerated"
	sqlc "github.com/volatiletech/boilbench/go_pg/sqlcgenerated"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

var dbURL = "postgresql://postgres:postgres@localhost:5432/boilbench?sslmode=disable"

func init() {
	url := os.Getenv("DB_URL")
	if url != "" {
		dbURL = url
	}
}

func benchmarkDB() *pg.DB {
	opts, err := pg.ParseURL(dbURL)
	if err != nil {
		panic(err)
	}
	return pg.Connect(opts)
}

type goPgDbLogger struct{}

func (d goPgDbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d goPgDbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, err := q.FormattedQuery()
	if err != nil {
		return err
	}
	fmt.Println(string(query))
	return nil
}

func BenchmarkQueryRowsBoil(b *testing.B) {
	seedDB()

	pqdb, err := pqdb()
	if err != nil {
		b.Fatal(err)
	}

	defer pqdb.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rs, err := boilmodels.Records(
				qm.Limit(100),
			).All(pqdb)
			if err != nil {
				b.Fatal(err)
			}
			if len(rs) != 100 {
				b.Fatalf("got %d, wanted 100", len(rs))
			}
		}
	})
}

func BenchmarkQueryRowsGopgReflect(b *testing.B) {
	seedDB()

	db := benchmarkDB()
	defer db.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var rs []Record
			_, err := db.Query(&rs, `SELECT * FROM records LIMIT 100`)
			if err != nil {
				b.Fatal(err)
			}
			if len(rs) != 100 {
				b.Fatalf("got %d, wanted 100", len(rs))
			}
		}
	})
}

func BenchmarkQueryRowsGopgORM(b *testing.B) {
	seedDB()

	db := benchmarkDB()
	defer db.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var rs []Record
			err := db.Model(&rs).Limit(100).Select()
			if err != nil {
				b.Fatal(err)
			}
			if len(rs) != 100 {
				b.Fatalf("got %d, wanted 100", len(rs))
			}
		}
	})
}

func BenchmarkQueryRowsSqlc(b *testing.B) {
	seedDB()

	pqdb, err := pqdb()
	if err != nil {
		b.Fatal(err)
	}

	sqlcdb := sqlc.New(pqdb)
	defer sqlcdb.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rs, err := sqlcdb.QueryRows(context.Background())
			if err != nil {
				b.Fatal(err)
			}
			if len(rs) != 100 {
				b.Fatalf("got %d, wanted 100", len(rs))
			}
		}
	})
}

func BenchmarkQueryRowsLibPq(b *testing.B) {
	seedDB()

	pqdb, err := pqdb()
	if err != nil {
		b.Fatal(err)
	}
	defer pqdb.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rows, err := pqdb.Query(`SELECT * FROM records LIMIT 100`)
			if err != nil {
				b.Fatal(err)
			}

			var rs []Record
			for rows.Next() {
				rs = append(rs, Record{})
				rec := &rs[len(rs)-1]

				err := rows.Scan(&rec.ID, &rec.Num1, &rec.Num2, &rec.Num3, &rec.Str1, &rec.Str2, &rec.Str3)
				if err != nil {
					b.Fatal(err)
				}
			}
			rows.Close()

			if len(rs) != 100 {
				b.Fatalf("got %d, wanted 100", len(rs))
			}
		}
	})
}

func BenchmarkQueryRowsGORM(b *testing.B) {
	seedDB()

	db, err := gormdb()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var rs []Record
			err := db.Limit(100).Find(&rs).Error
			if err != nil {
				b.Fatal(err)
			}

			if len(rs) != 100 {
				b.Fatalf("got %d, wanted 100", len(rs))
			}
		}
	})
}

func BenchmarkModelHasOneBoil(b *testing.B) {
	seedDB()

	pqdb, err := pqdb()
	if err != nil {
		b.Fatal(err)
	}

	defer pqdb.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rs, err := boilmodels.Books(
				qm.Load(boilmodels.BookRels.Author),
				qm.Limit(100),
			).All(pqdb)
			if err != nil {
				b.Fatal(err)
			}
			if len(rs) != 100 {
				b.Fatalf("got %d, wanted 100", len(rs))
			}
		}
	})
}

func BenchmarkModelHasOneGopg(b *testing.B) {
	seedDB()

	db := benchmarkDB()
	defer db.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var books []Book
			err := db.Model(&books).Column("book.*").Relation("Author").Limit(100).Select()
			if err != nil {
				b.Fatal(err)
			}

			if len(books) != 100 {
				b.Fatalf("got %d, wanted 100", len(books))
			}
		}
	})
}

func BenchmarkModelHasOneGORM(b *testing.B) {
	seedDB()

	db, err := gormdb()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var books []Book
			err := db.Preload("Author").Limit(100).Find(&books).Error
			if err != nil {
				b.Fatal(err)
			}

			if len(books) != 100 {
				b.Fatalf("got %d, wanted 100", len(books))
			}
		}
	})
}

func BenchmarkModelHasOneSqlc(b *testing.B) {
	seedDB()

	pqdb, err := pqdb()
	if err != nil {
		b.Fatal(err)
	}

	sqlcdb := sqlc.New(pqdb)
	defer sqlcdb.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			books, err := sqlcdb.ModelHasOne(context.Background())
			if err != nil {
				b.Fatal(err)
			}

			if len(books) != 100 {
				b.Fatalf("got %d, wanted 100", len(books))
			}
		}
	})
}

func BenchmarkModelHasManyBoil(b *testing.B) {
	seedDB()

	pqdb, err := pqdb()
	if err != nil {
		b.Fatal(err)
	}

	defer pqdb.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rs, err := boilmodels.Books(
				qm.Load(boilmodels.BookRels.Translations),
				qm.Limit(100),
			).All(pqdb)
			if err != nil {
				b.Fatal(err)
			}
			if len(rs) != 100 {
				b.Fatalf("got %d, wanted 100", len(rs))
			}
		}
	})
}

func BenchmarkModelHasManyGopg(b *testing.B) {
	seedDB()

	db := benchmarkDB()
	defer db.Close()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var books []Book
			err := db.Model(&books).Column("book.*").Relation("Translations").Limit(100).Select()
			if err != nil {
				b.Fatal(err)
			}

			if len(books) != 100 {
				b.Fatalf("got %d, wanted 100", len(books))
			}
			for _, book := range books {
				if len(book.Translations) != 10 {
					b.Fatalf("got %d, wanted 10", len(book.Translations))
				}
			}
		}
	})
}

func BenchmarkModelHasManyGORM(b *testing.B) {
	seedDB()

	db, err := gormdb()
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var books []Book
			err := db.Preload("Translations").Limit(100).Find(&books).Error
			if err != nil {
				b.Fatal(err)
			}

			if len(books) != 100 {
				b.Fatalf("got %d, wanted 100", len(books))
			}
			for _, book := range books {
				if len(book.Translations) != 10 {
					b.Fatalf("got %d, wanted 10", len(book.Translations))
				}
			}
		}
	})
}

func BenchmarkModelHasManySqlc(b *testing.B) {
	seedDB()

	pqdb, err := pqdb()
	if err != nil {
		b.Fatal(err)
	}

	sqlcdb := sqlc.New(pqdb)
	defer sqlcdb.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			books, err := sqlcdb.ModelHasManyBooks(context.Background())
			if err != nil {
				b.Fatal(err)
			}
			bookIDs := make([]int64, 0, len(books))
			for _, book := range books {
				bookIDs = append(bookIDs, book.ID)
			}
			translations, err := sqlcdb.ModelHasManyRelatedTranslations(context.Background(), bookIDs)
			if err != nil {
				b.Fatal(err)
			}
			if len(books) != 100 {
				b.Fatalf("got %d, wanted 100", len(books))
			}
			if len(translations) != 1000 {
				b.Fatalf("got %d, wanted 100", len(translations))
			}
		}
	})
}

func BenchmarkModelHasMany2ManyBoil(b *testing.B) {
	seedDB()

	pqdb, err := pqdb()
	if err != nil {
		b.Fatal(err)
	}

	defer pqdb.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rs, err := boilmodels.Books(
				qm.Load(boilmodels.BookRels.BookGenres),
				qm.Limit(100),
			).All(pqdb)
			if err != nil {
				b.Fatal(err)
			}
			if len(rs) != 100 {
				b.Fatalf("got %d, wanted 100", len(rs))
			}
		}
	})
}

func BenchmarkModelHasMany2ManyGopg(b *testing.B) {
	seedDB()

	db := benchmarkDB()
	defer db.Close()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var books []Book
			err := db.Model(&books).
				Column("book.*").Relation("Genres").
				Limit(100).
				Select()

			if err != nil {
				b.Fatal(err)
			}

			if len(books) != 100 {
				b.Fatalf("got %d, wanted 100", len(books))
			}
			for _, book := range books {
				if len(book.Genres) != 10 {
					b.Fatalf("got %d, wanted 10", len(book.Genres))
				}
			}
		}
	})
}

func BenchmarkModelHasMany2ManySqlc(b *testing.B) {
	seedDB()

	pqdb, err := pqdb()
	if err != nil {
		b.Fatal(err)
	}

	sqlcdb := sqlc.New(pqdb)
	defer sqlcdb.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			books, err := sqlcdb.ModelHasMany2ManyBooks(context.Background())
			if err != nil {
				b.Fatal(err)
			}
			bookIDs := make([]int64, 0, len(books))
			for _, book := range books {
				bookIDs = append(bookIDs, book.ID)
			}
			genres, err := sqlcdb.ModelHasMany2ManyRelatedGenres(context.Background(), bookIDs)
			if err != nil {
				b.Fatal(err)
			}
			if len(books) != 100 {
				b.Fatalf("got %d, wanted 100", len(books))
			}
			if len(genres) != 1000 {
				b.Fatalf("got %d, wanted 1000", len(genres))
			}
		}
	})
}

type numLoader struct {
	Num int
}

func BenchmarkExecBoil(b *testing.B) {
	db := benchmarkDB()
	defer db.Close()

	qs := []string{
		`DROP TABLE IF EXISTS exec_test`,
		`CREATE TABLE exec_test
(
    id   bigserial,
    name VARCHAR(500),
    CONSTRAINT "primary" PRIMARY KEY (id)
);`,
	}
	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			b.Fatal(err)
		}
	}

	pqdb, err := pqdb()
	if err != nil {
		b.Fatal(err)
	}
	defer pqdb.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := queries.Raw(`INSERT INTO exec_test (name) VALUES ($1)`, "hello world").Exec(pqdb)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkExecGopg(b *testing.B) {
	db := benchmarkDB()
	defer db.Close()

	qs := []string{
		`DROP TABLE IF EXISTS exec_test`,
		`CREATE TABLE exec_test
(
    id   bigserial,
    name VARCHAR(500),
    CONSTRAINT "primary" PRIMARY KEY (id)
);`,
	}
	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := db.Exec(`INSERT INTO exec_test (name) VALUES (?)`, "hello world")
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkExecSqlc(b *testing.B) {
	db := benchmarkDB()
	defer db.Close()

	qs := []string{
		`DROP TABLE IF EXISTS exec_test`,
		`CREATE TABLE exec_test
(
    id   bigserial,
    name VARCHAR(500),
    CONSTRAINT "primary" PRIMARY KEY (id)
);`,
	}
	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			b.Fatal(err)
		}
	}

	pqdb, err := pqdb()
	if err != nil {
		b.Fatal(err)
	}

	sqlcdb := sqlc.New(pqdb)
	defer sqlcdb.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := sqlcdb.Exec(context.Background(), sql.NullString{"hello world", true})
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func pqdb() (*sql.DB, error) {
	return sql.Open("postgres", dbURL)
}

func gormdb() (*gorm.DB, error) {
	return gorm.Open("postgres", dbURL)
}

type Record struct {
	ID               int64 `pg:"type:bigserial,pk"`
	Num1             int64 `pg:"type:serial"`
	Num2             int64 `pg:"type:serial"`
	Num3             int64 `pg:"type:serial"`
	Str1, Str2, Str3 string
}

func (r *Record) GetNum1() int64 {
	return r.Num1
}

func (r *Record) GetNum2() int64 {
	return r.Num2
}

func (r *Record) GetNum3() int64 {
	return r.Num3
}

func (r *Record) GetStr1() string {
	return r.Str1
}

func (r *Record) GetStr2() string {
	return r.Str2
}

func (r *Record) GetStr3() string {
	return r.Str3
}

var seedDBOnce sync.Once

func seedDB() {
	seedDBOnce.Do(func() {
		if err := _seedDB(); err != nil {
			panic(err)
		}
	})
}

func _seedDB() error {
	db := benchmarkDB()
	defer db.Close()

	for i := 0; i < 1000; i++ {
		_, err := db.Exec(`
			INSERT INTO records (str1, str2, str3) VALUES (?, ?, ?)
		`, randSeq(100), randSeq(200), randSeq(300))
		if err != nil {
			return err
		}
	}

	for i := 1; i < 100; i++ {
		genre := Genre{
			Id:   i,
			Name: fmt.Sprintf("genre %d", i),
		}
		err := db.Insert(&genre)
		if err != nil {
			return err
		}

		author := Author{
			ID:   i,
			Name: fmt.Sprintf("author %d", i),
		}
		err = db.Insert(&author)
		if err != nil {
			return err
		}
	}

	for i := 1; i <= 1000; i++ {
		err := db.Insert(&Book{
			Id:        i,
			Title:     fmt.Sprintf("book %d", i),
			AuthorID:  rand.Intn(99) + 1,
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}

		for j := 1; j <= 10; j++ {
			err = db.Insert(&BookGenre{
				BookId:  i,
				GenreId: j,
			})
			if err != nil {
				return err
			}

			err = db.Insert(&Translation{
				BookId: i,
				Lang:   fmt.Sprintf("%d", j),
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
