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
go install github.com/volatiletech/sqlboiler

# Generate models
sqlboiler --wipe --output ./models postgres -t "db"
