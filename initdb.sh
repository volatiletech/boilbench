#!/usr/bin/env sh

set -o errexit
set -o verbose

export PGUSER="${PGUSER:-postgres}"

# cd into initdb.sh directory
cd "$(dirname "$0")"

# Drop sqlboiler models
rm -rf models

# Drop and create DB
dropdb boilbench --if-exists
createdb -O "${PGUSER}" boilbench

# Import schema
psql boilbench -f schema.sql --quiet

# Install SQLBoiler
go get github.com/volatiletech/sqlboiler

# Install SQLBoiler PSQL driver
go get -u -t github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

# Generate models
sqlboiler --wipe --no-context --output ./models psql -t "db"

# Generate sqlc models
go generate ./sqlc
