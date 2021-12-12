package pops

import "github.com/volatiletech/null/v8"

// Pilot struct
type Pilot struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// Jet struct
type Jet struct {
	ID         int         `db:"id"`
	PilotID    int         `db:"pilot_id"`
	AirportID  int         `db:"airport_id"`
	Name       string      `db:"name"`
	Color      null.String `db:"color"`
	UUID       string      `db:"uuid"`
	Identifier string      `db:"identifier"`
	Cargo      []byte      `db:"cargo"`
	Manifest   []byte      `db:"manifest"`
}

// Airport struct
type Airport struct {
	ID   int      `db:"id"`
	Size null.Int `db:"size"`
}

// License struct
type License struct {
	ID      int `db:"id"`
	PilotID int `db:"pilot_id"`
}

// Hangar struct
type Hangar struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// Language struct
type Language struct {
	ID       int    `db:"id"`
	Language string `db:"language"`
}
