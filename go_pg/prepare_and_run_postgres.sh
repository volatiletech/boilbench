#!/usr/bin/env bash

set -o errexit
set -o verbose

SCRIPTPATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

# cd into initdb_postgres.sh directory
cd "${SCRIPTPATH}"

# DB port
port=5432

# Stop any docker linked to that port
for id in $(docker ps -q)
do
    if [[ $(docker port "${id}") == *"${port}"* ]]; then
        echo "stopping container ${id}"
        docker stop "${id}"
    fi
done

# Start postgres docker
dockerID=`docker run -d -p ${port}:5432 \
-e POSTGRES_PASSWORD=postgres \
-e POSTGRES_USER=postgres \
-e POSTGRES_DB=boilbench \
--mount type=bind,source=${SCRIPTPATH}/postgres_schema.sql,target=/docker-entrypoint-initdb.d/postgres_schema.sql \
postgres`

# Wait for docker to start
sleep 2

# Generate models
go generate ./...

go run github.com/volatiletech/sqlboiler --wipe --no-context --output ./boilgenerated psql -t "db"

DB_URL="postgresql://postgres:postgres@localhost:${port}/boilbench?sslmode=disable" go test -bench . -benchmem

# stop docker
docker stop ${dockerID}
docker rm -f ${dockerID}