CREATE TABLE authors (
	id BIGINT NOT NULL DEFAULT unique_rowid(),
	name TEXT NULL,
	avatar_id BIGINT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id )
);

CREATE UNIQUE INDEX authors_name_key ON authors (name);

CREATE TABLE book_genres (
	book_id BIGINT NOT NULL DEFAULT unique_rowid(),
	genre_id BIGINT NOT NULL DEFAULT unique_rowid(),
	genre__rating BIGINT NULL,
	CONSTRAINT "primary" PRIMARY KEY (book_id , genre_id )
);

CREATE TABLE books (
	id BIGINT NOT NULL DEFAULT unique_rowid(),
	title TEXT NULL,
	author_id BIGINT NULL,
	editor_id BIGINT NULL,
	created_at TIMESTAMPTZ NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NULL,
	CONSTRAINT "primary" PRIMARY KEY (id )
);

CREATE TABLE comments (
    id BIGINT NOT NULL DEFAULT unique_rowid(),
	trackable_id BIGINT NULL,
	trackable_type TEXT NULL,
	text TEXT NULL,
    CONSTRAINT "primary" PRIMARY KEY (id )
);

CREATE TABLE genres (
	id BIGINT NOT NULL DEFAULT unique_rowid(),
	name TEXT NULL,
	parent_id BIGINT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id )
);

CREATE TABLE images (
	id BIGINT NOT NULL DEFAULT unique_rowid(),
	path TEXT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id )
);

CREATE TABLE records (
    id BIGINT NOT NULL DEFAULT unique_rowid(),
	num1 BIGINT NOT NULL DEFAULT unique_rowid(),
	num2 BIGINT NOT NULL DEFAULT unique_rowid(),
	num3 BIGINT NOT NULL DEFAULT unique_rowid(),
	str1 TEXT NULL,
	str2 TEXT NULL,
	str3 TEXT NULL,
    CONSTRAINT "primary" PRIMARY KEY (id )
);

CREATE TABLE translations (
	id BIGINT NOT NULL DEFAULT unique_rowid(),
	book_id BIGINT NULL,
	lang TEXT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id)
);

CREATE UNIQUE INDEX translations_book_id_lang_key ON translations(book_id, lang);

CREATE TABLE exec_test
(
    id   bigserial,
    name VARCHAR(500),
    CONSTRAINT "primary" PRIMARY KEY (id)
);

ALTER TABLE ONLY book_genres ADD CONSTRAINT book_genres_book_id_book_id_foreign FOREIGN KEY (book_id) REFERENCES books(id);
ALTER TABLE ONLY book_genres ADD CONSTRAINT book_genres_genre_id_genre_id_foreign FOREIGN KEY (genre_id) REFERENCES genres(id);
ALTER TABLE ONLY books ADD CONSTRAINT books_author_id_author_id_foreign FOREIGN KEY (author_id) REFERENCES authors(id);
ALTER TABLE ONLY books ADD CONSTRAINT books_editor_id_author_id_foreign FOREIGN KEY (editor_id) REFERENCES authors(id);
ALTER TABLE ONLY genres ADD CONSTRAINT genres_parent_id_genre_id_foreign FOREIGN KEY (parent_id) REFERENCES genres(id);
ALTER TABLE ONLY translations ADD CONSTRAINT translations_book_id_book_id_foreign FOREIGN KEY (book_id) REFERENCES books(id);
