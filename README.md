## SQLBoiler Benchmarks

### Requirements

This repo requires Go 1.14+ to run.

To generate models, a running docker daemon is required.

### Instructions

To run the benchmarks, run: `go test -bench . -benchmem`

To generate the models, run: `./scripts/gen-models`

**Note**: There are some ruby and python scripts for generating graphs from
many runs of these benchmarks. They can be used to help update the sqlboiler
README with new graphs.

Graphs can be found in the [SQLBoiler](https://github.com/volatiletech/sqlboiler) readme.

The homepage for the [SQLBoiler](https://github.com/volatiletech/sqlboiler) is located at: https://github.com/volatiletech/sqlboiler
