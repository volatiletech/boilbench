package gopg

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10/orm"
)

type Genre struct {
	// tableName is an optional field that specifies custom table name and alias.
	// By default go-pg generates table name and alias from struct name.
	tableName struct{} `pg:"genres,alias:genre"` // default values are the same

	Id     int // Id is automatically detected as primary key
	Name   string
	Rating int `pg:"-"` // - is used to ignore field

	Books []Book `pg:"many2many:book_genres"` // many to many relation

	ParentId  int
	Subgenres []Genre `pg:"fk:parent_id"`
}

func (g Genre) String() string {
	return fmt.Sprintf("Genre<Id=%d Name=%q>", g.Id, g.Name)
}

type Image struct {
	Id   int
	Path string
}

type Author struct {
	ID    int     // both "Id" and "ID" are detected as primary key
	Name  string  `pg:",unique"`
	Books []*Book // has many relation

	AvatarId int
	Avatar   Image
}

func (a Author) String() string {
	return fmt.Sprintf("Author<ID=%d Name=%q>", a.ID, a.Name)
}

type BookGenre struct {
	tableName struct{} `pg:"alias:bg"` // custom table alias

	BookId  int `pg:",pk"` // pk tag is used to mark field as primary key
	Book    *Book
	GenreId int `pg:",pk"`
	Genre   *Genre

	Genre_Rating int // belongs to and is copied to Genre model
}

type Book struct {
	Id        int
	Title     string
	AuthorID  int
	Author    Author // has one relation
	EditorID  int
	Editor    *Author   // has one relation
	CreatedAt time.Time `pg:"default:now()"`
	UpdatedAt time.Time

	Genres       []Genre       `pg:"many2many:book_genres"` // many to many relation
	Translations []Translation // has many relation
	Comments     []Comment     `pg:"polymorphic:trackable_"` // has many polymorphic relation
}

var _ orm.BeforeInsertHook = (*Book)(nil)

func (b Book) String() string {
	return fmt.Sprintf("Book<Id=%d Title=%q>", b.Id, b.Title)
}

func (b *Book) BeforeInsert(c context.Context) (context.Context, error) {
	if b.CreatedAt.IsZero() {
		b.CreatedAt = time.Now()
	}
	return c, nil
}

// BookWithCommentCount is like Book model, but has additional CommentCount
// field that is used to select data into it. The use of `pg:",inherit"` tag
// is essential here so it inherits internal model properties such as table name.
type BookWithCommentCount struct {
	Book `pg:",inherit"`

	CommentCount int
}

type Translation struct {
	tableName struct{} `pg:",alias:tr"` // custom table alias

	Id     int
	BookId int    `pg:"unique:book_id_lang"`
	Book   *Book  // has one relation
	Lang   string `pg:"unique:book_id_lang"`

	Comments []Comment `pg:",polymorphic:trackable_"` // has many polymorphic relation
}

type Comment struct {
	Id            int    `pg:",pk"`
	TrackableId   int    // Book.Id or Translation.Id
	TrackableType string // "Book" or "Translation"
	Text          string
}
