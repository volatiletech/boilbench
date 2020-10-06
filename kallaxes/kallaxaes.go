package kallaxes

import "github.com/networkteam/go-kallax"

//go:generate go run github.com/networkteam/go-kallax/generator/cli/kallax

type Airport struct {
	kallax.Model `table:"airports"`

	ID   int64 `pk:"autoincr"`
	Size *int
}

type Hangar struct {
	kallax.Model `table:"hangars"`

	ID   int64 `pk:"autoincr"`
	Name *string
}

type Jet struct {
	kallax.Model `table:"jets"`

	ID         int64 `pk:"autoincr"`
	PilotID    int
	AirportID  int
	Name       string
	Color      *string
	UUID       string
	Identifier string
	Cargo      []byte
	Manifest   []byte
}

type Language struct {
	kallax.Model `table:"languages"`

	ID       int64 `pk:"autoincr"`
	Language string
}

type License struct {
	kallax.Model `table:"licenses"`

	ID      int64  `pk:"autoincr"`
	PilotID *int64 `fk:"pilot_id"`
}

type Pilot struct {
	kallax.Model `table:"pilots"`

	ID   int64 `pk:"autoincr"`
	Name string
}
