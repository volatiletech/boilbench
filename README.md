## SQLBoiler Benchmarks

### Requirements

These tests require go1.11 (go1.11 has support for go modules that are used here, go1.14 is recommended as current benchmarks were run using go1.14)

### Instructions

1. Clone this repository.
1. Update a `sqlboiler.toml` in the root directory if needed.
1. Modify port `initdb.sh` and run `./initdb.sh` to create your db and generate your models folder.
1. Run the benches against mocked (mimic) driver: `go test -bench . -benchmem`

Previous results can be found in `./results/mocked_driver.txt`

### go-pg benchmarks
As [go-pg](https://github.com/go-pg/pg) is not implementing `database/sql`,
to compare `database/sql` based implementations benchmarks should be run against real db

in `/go_pg` directory you will find files related to benchmarks comparison of [sqlc](https://github.com/kyleconroy/sqlc), [sqlboiler](https://github.com/volatiletech/sqlboiler) against go-pg both for CockroachDB and PostgresDB (all using docker), benchmarks are using same approach as described in [go-pg](https://github.com/go-pg/pg/blob/v10.0.0-beta.1/bench_test.go) repository

1. Run benchmark against CockroachDB (it will also set up DB, generate all needed files): `./go_pg/prepare_and_run_cockroach.sh`
    Previous results can be found in `./results/go_pg_cockroach.txt`
1. Run benchmark against PostgresDB (it will also set up DB, generate all needed files): `./go_pg/prepare_and_run_postgres.sh`
    Previous results can be found in `./results/go_pg_postgres.txt`

Inside each correspondent script variable `port` is defined. Override it if standard DB ports are already occupied on your machine.

**Note**: There are some ruby and python scripts for generating graphs from
many runs of these benchmarks. They can be used to help update the sqlboiler
README with new graphs.

Graphs can be found in the [SQLBoiler](https://github.com/volatiletech/sqlboiler) readme.

The homepage for the [SQLBoiler](https://github.com/volatiletech/sqlboiler) is located at: https://github.com/volatiletech/sqlboiler 
