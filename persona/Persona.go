package persona

import (
	"time"
)

// Persona es el modelo de la tabla Persona
type Persona struct {
	PersonaPk       uint32    `json:"persona_pk"`
	Cedula          string    `json:"cedula"`
	Nombres         string    `json:"nombres"`
	Apellidos       string    `json:"apellidos"`
	FechaNacimiento time.Time `json:"fechanacimiento"`
	Genero          string    `json:"genero"`
	FechaCreacion   time.Time `json:"fechacreacion"`
	Estado          bool      `json:"estado"`
}

// PersonaNativa es el modelo de la tabla Persona
type PersonaNativa struct {
	PersonaPk       string `json:"persona_pk"`
	Cedula          string `json:"cedula"`
	Nombres         string `json:"nombres"`
	Apellidos       string `json:"apellidos"`
	FechaNacimiento string `json:"fechanacimiento"`
	Genero          string `json:"genero"`
	FechaCreacion   string `json:"fechacreacion"`
	Estado          string `json:"estado"`
}
