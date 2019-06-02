package persona

import "github.com/fochoac/go-model-ws/prueba"

type PersonaRespuesta struct {
	CodigoRetorno   prueba.CodigoRetorno
	Personas        []Persona
	PersonasNativas []PersonaNativa
}
