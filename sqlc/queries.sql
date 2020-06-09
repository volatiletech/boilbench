-- name: DeleteJet :exec
DELETE FROM "jets"
WHERE "id"=$1;

-- name: CreateJet :exec
INSERT INTO "jets" (
"id","pilot_id","airport_id","name","color","uuid","identifier","cargo","manifest"
) VALUES (
$1,$2,$3,$4,$5,$6,$7,$8,$9
);



-- name: UpdateJets :exec
UPDATE "jets"
SET "pilot_id"=$1,"airport_id"=$2,"name"=$3,"color"=$4,"uuid"=$5,"identifier"=$6,"cargo"=$7,"manifest"=$8
WHERE "id"=$9;


-- name: ListJets :many
SELECT * FROM "jets";

-- name: ListJetsSubset :many
SELECT id, name, color, uuid, identifier, cargo, manifest FROM "jets";

-- name: ListJetsComplex :many
SELECT id, name, color, uuid, identifier, cargo, manifest FROM "jets"
WHERE (id > $1) AND (name <> $2) GROUP BY id LIMIT 1 OFFSET 1;