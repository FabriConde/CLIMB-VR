package models

import (
	"fmt"
	"strings"
)

type Sector struct {
	nombre    string
	ubicacion *Ubicacion
	vias      map[string]*Via
	numVias   int
}

func NewSector(nombre string, ubicacion *Ubicacion, vias map[string]*Via) *Sector {
	return &Sector{
		nombre:    nombre,
		ubicacion: ubicacion,
		vias:      vias,
		numVias:   len(vias),
	}
}

func (s *Sector) GetVias() map[string]*Via {
	return s.vias
}

func (s *Sector) AñadirVia(via *Via) {
	s.vias[via.nombre] = via
	s.numVias = len(s.vias)
}

func (s *Sector) ToString() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Sector: %s\n", s.nombre))
	sb.WriteString(fmt.Sprintf("Ubicación: %s\n", s.ubicacion.ToString()))
	sb.WriteString(fmt.Sprintf("Número de Vías: %d\n", s.numVias))
	sb.WriteString("Vías:\n")
	for _, via := range s.GetVias() {
		sb.WriteString(fmt.Sprintf("  %s\n", via.ToString()))
	}
	return sb.String()
}
