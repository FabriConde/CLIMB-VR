package models

import (
	"fmt"
	"strings"
)

type MapaSectores struct {
	sectores map[string]*Sector
}

func NewMapaSectores() *MapaSectores {
	return &MapaSectores{
		sectores: make(map[string]*Sector),
	}
}

func (ms *MapaSectores) AddSector(sector *Sector) {
	if ms.sectores == nil {
		ms.sectores = make(map[string]*Sector)
	}
	ms.sectores[sector.nombre] = sector
}

func (ms *MapaSectores) ComprobarSiExisteSectorEnElMapaPorNombre(nombreSector string) (bool, *Sector) {
	sector, ok := ms.sectores[nombreSector]
	if !ok {
		return false, nil
	}
	return true, sector
}

func (ms *MapaSectores) ToString() string {
	var sb strings.Builder
	sb.WriteString("Mapa de Sectores:\n")
	for _, sector := range ms.sectores {
		sb.WriteString(fmt.Sprintf("%s\n", sector.ToString()))
	}
	return sb.String()
}
