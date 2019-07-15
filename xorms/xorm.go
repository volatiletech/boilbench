package xorms

import (
	"github.com/volatiletech/null"

	// Shutup linter
	_ "github.com/lib/pq"
)

// Pilot struct
type Pilot struct {
	Id        int        `xorm:"pk"`
	Name      string     `xorm:"not null"`
	Languages []Language `xorm:"extends"`
}

// Jet struct
type Jet struct {
	Id int `xorm:"pk"`

	PilotId int `xorm:"not null"`

	AirportId int `xorm:"not null"`

	Name       string `xorm:"not null"`
	Color      null.String
	Uuid       string `xorm:"not null"`
	Identifier string `xorm:"not null"`
	Cargo      []byte `xorm:"not null"`
	Manifest   []byte `xorm:"not null"`
}

// Airport struct
type Airport struct {
	Id   int `xorm:"pk"`
	Size null.Int
}

// License struct
type License struct {
	Id int `xorm:"pk"`

	Pilot   Pilot
	PilotId int
}

// Hangar struct
type Hangar struct {
	Id   int    `xorm:"pk"`
	Name string `xorm:"not null"`
}

// Language struct
type Language struct {
	Id       int    `xorm:"pk"`
	Language string `xorm:"index not null"`
}

type PilotLanguage struct {
	PilotId    int `xorm:"pk"`
	LanguageId int `xorm:"pk"`
}
