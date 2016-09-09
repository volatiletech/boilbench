## SQLBoiler Benchmarks

### Requirements

These tests require go1.7

### Instructions

1. Install `sqlboiler`: `go get -u -t github.com/vattle/sqlboiler`
2. Clone this repository.
3. Create a `sqlboiler.toml` in the root directory.
4. Run `./initdb.sh` to create your db and generate your models folder.
5. Run the benches: `go test -bench . -benchmem`

`sqlboiler.toml` Example:

```
[postgres]
  host="localhost"
  port=5432
  user="postgres"
  pass="yourpassword"
  dbname="boilbench"
```

Graphs can be found in the [SQLBoiler](https://github.com/vattle/sqlboiler) readme.

The homepage for the [SQLBoiler](https://github.com/vattle/sqlboiler)  [Golang ORM](https://github.com/vattle/sqlboiler) generator is located at: https://github.com/vattle/sqlboiler 
