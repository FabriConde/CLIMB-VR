package models

import (
	"fmt"
)

type Ubicacion struct {
	descripcion   string
	ubicacionMaps *string
}

func NewUbicacion(descripcion, ubicacionMaps string) *Ubicacion {
	var ubicacionMapsPtr *string
	if ubicacionMaps != "" {
		ubicacionMapsPtr = &ubicacionMaps
	}

	return &Ubicacion{
		descripcion:   descripcion,
		ubicacionMaps: ubicacionMapsPtr,
	}
}

func (u *Ubicacion) ToString() string {
	if u.ubicacionMaps != nil {
		return fmt.Sprintf("Descripción: %s, Ubicación Maps: %s", u.descripcion, *u.ubicacionMaps)
	}
	return fmt.Sprintf("Descripción: %s", u.descripcion)
}
