#!/usr/bin/env bash

set -o errexit
set -o verbose

# cd into initdb.sh directory
cd "$(dirname "$0")"

# DB port
port=26257

# Stop any docker linked to that port
for id in $(docker ps -q)
do
    if [[ $(docker port "${id}") == *"${port}"* ]]; then
        echo "stopping container ${id}"
        docker stop "${id}"
    fi
done

# Start cockroach docker
dockerID=`docker run -d -p ${port}:26257 cockroachdb/cockroach:v19.1.0 start --insecure`
# Wait for docker to start
sleep 2
# Create DB
docker exec -ti ${dockerID} ./cockroach sql --insecure --execute="CREATE DATABASE IF NOT EXISTS boilbench;"

schemaContent=`cat cockroach_schema.sql`

# Import schema
docker exec -ti ${dockerID} ./cockroach sql --insecure --database boilbench --execute="${schemaContent}"

# Wait for docker to start
sleep 2

# Generate models
go generate ./...

go run github.com/volatiletech/sqlboiler --wipe --no-context --output ./boilgenerated crdb -t "db"

DB_URL="postgresql://root@localhost:${port}/boilbench?sslmode=disable" go test -bench . -benchmem

# stop docker
docker stop ${dockerID}
docker rm -f ${dockerID}

