## SQLBoiler Benchmarks

### Requirements

These tests require go1.7

### Instructions

1. Clone this repository.
1. Create a `sqlboiler.toml` in the root directory.
1. Run `./initdb.sh` to create your db and generate your models folder.
1. Run the benches: `go test -bench . -benchmem`

**Note**: There are some ruby and python scripts for generating graphs from
many runs of these benchmarks. They can be used to help update the sqlboiler
README with new graphs.

Graphs can be found in the [SQLBoiler](https://github.com/volatiletech/sqlboiler) readme.

The homepage for the [SQLBoiler](https://github.com/volatiletech/sqlboiler) is located at: https://github.com/volatiletech/sqlboiler 
