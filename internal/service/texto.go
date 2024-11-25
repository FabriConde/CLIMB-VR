package service

import (
	"climb_vr/internal/models"
	"fmt"
	"strconv"
	"strings"
)

func AgregarSectorDesdeTexto(mapaSectores *models.MapaSectores, texto string) (*models.MapaSectores, error) {
	ok, lineasConSectorYVias, err := ComprobarQueSePuedeCrearSectorYVias(texto)
	if !ok {
		return nil, err
	}

	infoSector := SepararEnPartesUnaLineaPorLaComa(lineasConSectorYVias[0])
	nombreSector := ObtenerNombreSector(infoSector[0])
	ubicacion, err := GuardarUbicacionSector(infoSector)
	if err != nil {
		return nil, err
	}

	ok, sector := mapaSectores.ComprobarSiExisteSectorEnElMapaPorNombre(nombreSector)
	if !ok {
		sector = models.NewSector(nombreSector, ubicacion, make(map[string]*models.Via))
	}

	vias, err := CrearVias(lineasConSectorYVias[1:], sector.GetVias())
	if err != nil {
		return nil, err
	}

	for _, via := range vias {
		sector.AñadirVia(via)
	}
	mapaSectores.AddSector(sector)

	return mapaSectores, nil
}

func ComprobarQueSePuedeCrearSectorYVias(texto string) (bool, []string, error) {
	if ok, err := ComprobarQueTextoTieneSaltoDeLineasYComas(texto); !ok {
		return false, nil, err
	}

	lineasTexto := DividirTextoEnLineasEliminadoLasLineasVacias(texto)
	if ok, err := ComprobarQueSoloExisteInformacionDelSectorAlInicioYNoSeRepite(lineasTexto); !ok {
		return false, nil, err
	}

	if ok, err := ComprobarQueExisteLaPalabraVias(lineasTexto); !ok {
		return false, nil, err
	}

	lineasConVias := ObtenerLasLineasDondeEstanLasViasElimnadoLasPalabrasViasQueSeRepite(lineasTexto)
	lineasConSectorYVias := append(lineasTexto[:1], lineasConVias...)
	if ok, err := ComprobarQueLasLineasSonSuficientesParaCrearSectorYVias(lineasConSectorYVias); !ok {
		return false, nil, err
	}

	infoSector := SepararEnPartesUnaLineaPorLaComa(lineasConSectorYVias[0])
	if ok, err := ComprobarQueLaPrimeraLineaContieneInformacionNecesariaParaCrearSector(infoSector); !ok {
		return false, nil, err
	}

	if ok, err := ComprobarHayInformacionNecesariaParaCrearUnaVia(lineasConSectorYVias[1:]); !ok {
		return false, nil, err
	}

	return true, lineasConSectorYVias, nil
}

func ComprobarQueTextoTieneSaltoDeLineasYComas(texto string) (bool, error) {
	if !strings.Contains(texto, "\n") {
		return false, fmt.Errorf("el texto no contiene saltos de línea")
	}

	if !strings.Contains(texto, ",") {
		return false, fmt.Errorf("el texto no contiene comas")
	}

	return true, nil
}

func DividirTextoEnLineasEliminadoLasLineasVacias(texto string) []string {
	lineas := strings.Split(texto, "\n")
	lineas = EliminarLineasVaciasDelTexto(lineas)
	return lineas
}

func EliminarLineasVaciasDelTexto(lineas []string) []string {
	var lineasSinBlanco []string
	for _, linea := range lineas {
		linea = strings.TrimSpace(linea)
		if linea != "" {
			lineasSinBlanco = append(lineasSinBlanco, linea)
		}
	}
	return lineasSinBlanco
}

func ComprobarQueSoloExisteInformacionDelSectorAlInicioYNoSeRepite(lineas []string) (bool, error) {
	if !strings.HasPrefix(lineas[0], "Sector:") {
		return false, fmt.Errorf("la primera línea no contiene 'Sector:'")
	}

	for _, linea := range lineas[1:] {
		if strings.Contains(linea, "Sector:") {
			return false, fmt.Errorf("existen varias lineas que contienen 'Sector:'")
		}
	}

	return true, nil
}

func ComprobarQueExisteLaPalabraVias(lineas []string) (bool, error) {
	for _, linea := range lineas {
		if strings.TrimSpace(linea) == "Vias:" {
			return true, nil
		}
	}
	return false, fmt.Errorf("no se encontró 'Vias:' en el texto")
}

func ObtenerLasLineasDondeEstanLasViasElimnadoLasPalabrasViasQueSeRepite(lineas []string) []string {
	var lineasDeVias []string
	seEncuntraPalabraVias := false

	for _, linea := range lineas {
		if strings.TrimSpace(linea) == "Vias:" {
			seEncuntraPalabraVias = true
			continue
		}

		if seEncuntraPalabraVias {
			if strings.TrimSpace(linea) == "" {
				seEncuntraPalabraVias = false
			} else {
				lineasDeVias = append(lineasDeVias, strings.TrimSpace(linea))
			}
		}
	}

	return lineasDeVias
}

func ComprobarQueLasLineasSonSuficientesParaCrearSectorYVias(lineas []string) (bool, error) {
	if len(lineas) < 2 {
		return false, fmt.Errorf("el texto no contiene suficiente información para crear un sector y sus vías")
	}

	return true, nil
}

func SepararEnPartesUnaLineaPorLaComa(linea string) []string {
	partes := strings.Split(linea, ",")
	partes = EliminarEspaciosEnBlanco(partes)
	return partes
}

func EliminarEspaciosEnBlanco(partes []string) []string {
	for i, parte := range partes {
		partes[i] = strings.TrimSpace(parte)
	}
	return partes
}

func ComprobarQueLaPrimeraLineaContieneInformacionNecesariaParaCrearSector(parteSector []string) (bool, error) {
	if len(parteSector) < 2 {
		return false, fmt.Errorf("no contiene suficiente información para crear un sector")
	}

	return true, nil
}

func ComprobarHayInformacionNecesariaParaCrearUnaVia(lineasConVias []string) (bool, error) {
	for _, linea := range lineasConVias {
		partesDeLinea := SepararEnPartesUnaLineaPorLaComa(linea)
		if len(partesDeLinea) < 3 {
			return false, fmt.Errorf("no contiene suficiente información para crear una vía")
		}
	}
	return true, nil
}

func ObtenerNombreSector(linea string) string {
	partes := strings.Split(linea, ":")
	return strings.TrimSpace(partes[1])
}

func GuardarUbicacionSector(infoSector []string) (*models.Ubicacion, error) {
	descripcionUbicacion := strings.TrimSpace(infoSector[1])
	var ubicacionMaps string

	if len(infoSector) > 2 {
		ubicacionMaps = infoSector[2]
		if ubicacionMaps != "" && !strings.HasPrefix(ubicacionMaps, "https") {
			return nil, fmt.Errorf("la URL de ubicacionMaps debe comenzar con 'https'")
		}
	}
	ubicacion := models.NewUbicacion(descripcionUbicacion, ubicacionMaps)

	return ubicacion, nil
}

func ExisteViaEnSector(nombreVia string, vias map[string]*models.Via) error {
	if _, ok := vias[nombreVia]; ok {
		return fmt.Errorf("la vía ya existe: %s", nombreVia)
	}
	return nil
}

func AddDificultadVia(partesTextoVia []string) (models.DificultadEscalada, error) {
	if !models.EsUnaDificultadValida(partesTextoVia[1]) {
		return "", fmt.Errorf("la dificultad no cumple con los estandares: %s", partesTextoVia[1])
	}

	return models.DificultadEscalada(partesTextoVia[1]), nil
}

func AddLongitudVia(partesTextoVia []string) (float32, error) {
	longitud, err := strconv.ParseFloat(partesTextoVia[2], 32)
	if err != nil {
		return 0, fmt.Errorf("longitud inválida: %s", partesTextoVia[2])
	}

	return float32(longitud), nil
}

func AddInfoExtraVia(partesTextoVia []string) *string {
	if len(partesTextoVia) > 3 {
		return &partesTextoVia[3]
	}
	return nil
}

func CrearVias(infoVias []string, vias map[string]*models.Via) (map[string]*models.Via, error) {
	totalViasAddPorTexto := make(map[string]*models.Via)

	for _, linea := range infoVias {
		partesDeVia := SepararEnPartesUnaLineaPorLaComa(linea)
		nombreVia := partesDeVia[0]

		if err := ExisteViaEnSector(nombreVia, vias); err != nil {
			return nil, err
		}

		dificultad, err := AddDificultadVia(partesDeVia)
		if err != nil {
			return nil, err
		}

		longitud, err := AddLongitudVia(partesDeVia)
		if err != nil {
			return nil, err
		}

		posibleInfoExtra := AddInfoExtraVia(partesDeVia)
		var infoExtra string
		if posibleInfoExtra != nil {
			infoExtra = *posibleInfoExtra
		}

		via := models.NewVia(nombreVia, dificultad, float32(longitud), infoExtra)
		totalViasAddPorTexto[nombreVia] = via
	}

	return totalViasAddPorTexto, nil
}
