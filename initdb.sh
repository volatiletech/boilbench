#!/usr/bin/env sh

# cd into initdb.sh directory
cd "$(dirname "$0")"

# Drop sqlboiler models
rm -rf models

# Drop and create DB
dropdb -U postgres boilbench --if-exists || { exit 1; }
createdb -U postgres -O postgres boilbench || { exit 1; }

# Import schema
psql -U postgres boilbench -f schema.sql --quiet || { exit 1; }

# Install SQLBoiler
go install github.com/vattle/sqlboiler

# Generate models
sqlboiler -o ./models postgres -t "db"
