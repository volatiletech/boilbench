-- name: QueryRows :many
SELECT * FROM records LIMIT 100;

-- name: ModelHasOne :many
SELECT *
FROM "books"
LEFT JOIN "authors"
ON "authors"."id" = "books"."author_id"
LIMIT 100;

-- name: ModelHasManyBooks :many
SELECT * FROM "books" LIMIT 100;

-- name: ModelHasManyRelatedTranslations :many
SELECT * FROM "translations"
WHERE book_id = ANY($1::bigint[]);

-- name: ModelHasMany2ManyBooks :many
SELECT * FROM "books" LIMIT 100;

-- name: ModelHasMany2ManyRelatedGenres :many
SELECT book_genres.book_id, book_genres.genre_id, book_genres.genre__rating, genre.name, genre.parent_id
FROM genres AS "genre" JOIN "book_genres"
ON ("book_genres"."book_id") = ANY($1::bigint[])
WHERE ("genre"."id" = "book_genres"."genre_id");

-- name: Exec :exec
INSERT INTO exec_test (name) VALUES ($1);