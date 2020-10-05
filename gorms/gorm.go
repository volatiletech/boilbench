package gorms

import (
	"github.com/volatiletech/null/v8"
)

// Pilot struct
type Pilot struct {
	ID        int
	Name      string     `gorm:"not null"`
	Languages []Language `gorm:"many2many:pilot_languages;"`
}

// Jet struct
type Jet struct {
	ID int

	Pilot   Pilot `gorm:"ForeignKey:PilotID"`
	PilotID int   `gorm:"not null"`

	Airport   Airport `gorm:"ForeignKey:Airport"`
	AirportID int     `gorm:"not null"`

	Name       string `gorm:"not null"`
	Color      null.String
	UUID       string `gorm:"not null"`
	Identifier string `gorm:"not null"`
	Cargo      []byte `gorm:"not null"`
	Manifest   []byte `gorm:"not null"`
}

// Airport struct
type Airport struct {
	ID   int
	Size null.Int
}

// License struct
type License struct {
	ID int

	Pilot   Pilot `gorm:"ForeignKey:PilotID"`
	PilotID int
}

// Hangar struct
type Hangar struct {
	ID   int
	Name string `gorm:"not null"`
}

// Language struct
type Language struct {
	ID       int
	Language string `gorm:"index:idx_pilot_languages;not null"`
}
