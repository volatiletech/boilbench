#!/usr/bin/env bash

set -o errexit
set -o verbose

# cd into initdb.sh directory
cd "$(dirname "$0")"

# Drop sqlboiler models
rm -rf models

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

schemaContent=`cat schema.sql`

# Import schema
docker exec -ti ${dockerID} ./cockroach sql --insecure --database boilbench --execute="${schemaContent}"

# Install SQLBoiler
go get github.com/volatiletech/sqlboiler

# Install SQLBoiler Cockroach DB driver
go get -u -t github.com/glerchundi/sqlboiler-crdb

# Generate models
sqlboiler --wipe --no-context --output ./models crdb -t "db"

# stop docker
docker stop ${dockerID}
docker rm -f ${dockerID}

# Generate sqlc models
go generate ./sqlc
