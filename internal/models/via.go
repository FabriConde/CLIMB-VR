package models

import (
	"fmt"
)

type DificultadEscalada string

const (
	Quinto     DificultadEscalada = "V"
	QuintoPlus DificultadEscalada = "V+"
	SeisA      DificultadEscalada = "6a"
	SeisAPlus  DificultadEscalada = "6a+"
	SeisB      DificultadEscalada = "6b"
	SeisBPlus  DificultadEscalada = "6b+"
	SeisC      DificultadEscalada = "6c"
	SeisCPlus  DificultadEscalada = "6c+"
	SieteA     DificultadEscalada = "7a"
	SieteAPlus DificultadEscalada = "7a+"
	SieteB     DificultadEscalada = "7b"
	SieteBPlus DificultadEscalada = "7b+"
	SieteC     DificultadEscalada = "7c"
	SieteCPlus DificultadEscalada = "7c+"
	OchoA      DificultadEscalada = "8a"
	OchoAPlus  DificultadEscalada = "8a+"
	OchoB      DificultadEscalada = "8b"
	OchoBPlus  DificultadEscalada = "8b+"
	OchoC      DificultadEscalada = "8c"
	OchoCPlus  DificultadEscalada = "8c+"
	NueveA     DificultadEscalada = "9a"
	NueveAPlus DificultadEscalada = "9a+"
	NueveB     DificultadEscalada = "9b"
	NueveBPlus DificultadEscalada = "9b+"
	NueveC     DificultadEscalada = "9c"
	NueveCPlus DificultadEscalada = "9c+"
)

type Via struct {
	nombre     string
	dificultad DificultadEscalada
	longitud   float32
	infoExtra  *string
}

func NewVia(nombre string, dificultad DificultadEscalada, longitud float32, infoExtra string) *Via {
	var infoExtraPtr *string
	if infoExtra != "" {
		infoExtraPtr = &infoExtra
	}
	return &Via{
		nombre:     nombre,
		dificultad: dificultad,
		longitud:   longitud,
		infoExtra:  infoExtraPtr,
	}
}

func EsUnaDificultadValida(dificultad string) bool {
	switch DificultadEscalada(dificultad) {
	case Quinto, QuintoPlus, SeisA, SeisAPlus, SeisB, SeisBPlus, SeisC, SeisCPlus,
		SieteA, SieteAPlus, SieteB, SieteBPlus, SieteC, SieteCPlus,
		OchoA, OchoAPlus, OchoB, OchoBPlus, OchoC, OchoCPlus,
		NueveA, NueveAPlus, NueveB, NueveBPlus, NueveC, NueveCPlus:
		return true
	default:
		return false
	}
}

func (v *Via) GetInfoExtra() string {
	if v.infoExtra != nil {
		return *v.infoExtra
	}
	return ""
}

func (v *Via) ToString() string {
	return fmt.Sprintf("Vía: %s, Dificultad: %s, Longitud: %.2f, Info Extra: %s",
		v.nombre, v.dificultad, v.longitud, v.GetInfoExtra())
}
