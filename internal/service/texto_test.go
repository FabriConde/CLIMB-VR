package service

import (
	"climb_vr/internal/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgregarSectorDesdeTextoCuandoSectorNoExiste(t *testing.T) {
	mapaSectores := models.NewMapaSectores()
	texto := "Sector: Sector Avispa, Desde que dejas el coche en el parking seguir el sendero a la derecha, https://ubiVias\n\nVías:\nAvispa, 6b, 15\nLepto, 7c, 34, placa"
	resultado, err := AgregarSectorDesdeTexto(mapaSectores, texto)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	fmt.Println(resultado.ToString())
}

func TestAgregarSectorDesdeTextoCuandoSectorYaExiste(t *testing.T) {
	mapaSectores := models.NewMapaSectores()
	ubicacion := models.NewUbicacion("Desde que dejas el coche en el parking seguir el sendero a la derecha", "https://ubiVias")
	vias := make(map[string]*models.Via)
	via := models.NewVia("Avispa pro", "9b", 15, "")
	vias["Avispa pro"] = via
	sector := models.NewSector("Sector Avispa", ubicacion, vias)
	mapaSectores.AddSector(sector)
	texto := "Sector: Sector Avispa, Desde que dejas el coche en el parking seguir el sendero a la derecha, https://ubiVias\n\nVias:\nAvispa, 6b, 15\nLepto, 7c, 34, placa,\n\n\nvías:"
	resultado, err := AgregarSectorDesdeTexto(mapaSectores, texto)

	assert.NoError(t, err)
	assert.NotNil(t, resultado)
	fmt.Println(resultado.ToString())
}

func TestNoSePuedeCrearSectorYViasPorQueTextoNoTieneSaltoDeLinea(t *testing.T) {
	mapaSectores := models.NewMapaSectores()
	texto := "Sector: Sector Avispa, Desde que dejas el coche en el parking seguir el sendero a la derecha, https://ubiVias Vias: Avispa, 6b, 15, Lepto, 7c, 34, placa"
	resultado, err := AgregarSectorDesdeTexto(mapaSectores, texto)

	assert.Error(t, err)
	assert.Equal(t, "el texto no contiene saltos de línea", err.Error())
	assert.Nil(t, resultado)
}

func TestNoSePuedeCrearSectorPorqueUbicacionNoEstaBienFormada(t *testing.T) {
	mapaSectores := models.NewMapaSectores()
	texto := "Sector: Sector Avispa, Desde que dejas el coche en el parking seguir el sendero a la derecha, ubiVias\n\nVias:\nAvispa, 6b, 15\nLepto, 7c, 34, placa"
	resultado, err := AgregarSectorDesdeTexto(mapaSectores, texto)

	assert.Error(t, err)
	assert.Equal(t, "la URL de ubicacionMaps debe comenzar con 'https'", err.Error())
	assert.Nil(t, resultado)
}

func TestNoSePuedeCrearViasPorqueYaExisteEsaVia(t *testing.T) {
	mapaSectores := models.NewMapaSectores()
	ubicacion := models.NewUbicacion("Desde que dejas el coche en el parking seguir el sendero a la derecha", "https://ubiVias")
	vias := make(map[string]*models.Via)
	via := models.NewVia("Avispa", "9b", 15, "")
	vias["Avispa"] = via
	sector := models.NewSector("Sector Avispa", ubicacion, vias)
	mapaSectores.AddSector(sector)
	texto := "Sector: Sector Avispa, Desde que dejas el coche en el parking seguir el sendero a la derecha, https://ubiVias\n\nVias:\nAvispa, 6b, 15\nLepto, 7c, 34, placa"
	resultado, err := AgregarSectorDesdeTexto(mapaSectores, texto)

	assert.Error(t, err)
	assert.Equal(t, "la vía ya existe en el sector", err.Error())
	assert.Nil(t, resultado)
}

func TestElTextoNoContieneComas(t *testing.T) {
	texto := "Sector: Desde que dejas el coche en el parking seguir el sendero a la derechahttps://ubiVias\n\nVias:\nAvispa 6b 15\nLepto 7c 34 placa"
	ok, resultado, err := ComprobarQueSePuedeCrearSectorYVias(texto)

	assert.False(t, ok)
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "el texto no contiene comas", err.Error())
}

func TestPrimeraLineaNoContieneLaPalabraSector(t *testing.T) {
	texto := "Desde que dejas el coche en el parking seguir el sendero a la derecha, https://ubiVias\n\nVias:\nAvispa, 6b, 15\nLepto, 7c, 34, placa"
	ok, resultado, err := ComprobarQueSePuedeCrearSectorYVias(texto)

	assert.False(t, ok)
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "la primera línea no contiene 'sector:'", err.Error())
}

func TestElTextoContieneMasDeUnaPalabraSector(t *testing.T) {
	texto := "Sector: Sector Avispa, Desde que dejas el coche en el parking seguir el sendero a la derecha, https://ubiVias\n\nVias:\nAvispa, 6b, 15\n Sector: Sector Avispa, Desde que dejas el coche en el parking seguir el sendero a la derecha, https://ubiVias \n Lepto, 7c, 34, placa"
	ok, resultado, err := ComprobarQueSePuedeCrearSectorYVias(texto)

	assert.False(t, ok)
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "existen varias lineas que contienen 'sector:'", err.Error())
}

func TestElTextoNoContieneLaPalabraVias(t *testing.T) {
	texto := "Sector: Sector Avispa, Desde que dejas el coche en el parking seguir el sendero a la derecha, https://ubiVias\n\n\nAvispa, 6b, 15\nLepto, 7c, 34, placa"
	ok, resultado, err := ComprobarQueSePuedeCrearSectorYVias(texto)

	assert.False(t, ok)
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "no se encontró 'Vias:' en el texto", err.Error())
}

func TestElTextoNoContieneLaInformacinSuficienteParaCrearUnSectorYVias(t *testing.T) {
	texto := "Sector: Sector Avispa, Desde que dejas el coche en el parking seguir el sendero a la derecha, https://ubiVias\n\nVias:"
	ok, resultado, err := ComprobarQueSePuedeCrearSectorYVias(texto)

	assert.False(t, ok)
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "el texto no contiene suficiente información para crear un sector y sus vías", err.Error())
}

func TestElTextoNoContieneLaInformacinSuficienteParaCrearUnSector(t *testing.T) {
	texto := "Sector: Sector Avispa,\n\nVias:\n Avispa, 6b, 15\nLepto, 7c, 34, placa"
	ok, resultado, err := ComprobarQueSePuedeCrearSectorYVias(texto)

	assert.False(t, ok)
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "no contiene suficiente información para crear un sector", err.Error())
}

func TestElTextoNoContieneLaInformacinSuficienteParaCrearUnSectorPeroSiUbicacionMaps(t *testing.T) {
	texto := "Sector: Sector Avispa, https://ubiVias\n\nVias:\n Avispa, 6b, 15\nLepto, 7c, 34, placa"
	ok, resultado, err := ComprobarQueSePuedeCrearSectorYVias(texto)

	assert.False(t, ok)
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "el segundo valor del sector no debe comenzar con 'https', dede ser una descripción", err.Error())
}

func TestElTextoNoContieneLaInformacinSuficienteParaCrearVias(t *testing.T) {
	texto := "Sector: Sector Avispa, Desde que dejas el coche en el parking seguir el sendero a la derecha, https://ubiVias\n\nVias: \nAvispa, 6b"
	ok, resultado, err := ComprobarQueSePuedeCrearSectorYVias(texto)

	assert.False(t, ok)
	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, "no contiene suficiente información para crear una vía", err.Error())
}

func TestDificultadDeViaNoEsValida(t *testing.T) {
	vias := make(map[string]*models.Via)
	lineas := []string{
		"Avispa, 23b, 15",
		"Lepto, 7c, 34, placa",
	}
	vias, err := CrearVias(lineas, vias)

	assert.Error(t, err)
	assert.Nil(t, vias)
	assert.Equal(t, "la dificultad no cumple con los estandares", err.Error())
}

func TestLongitudDeViaNoEsValida(t *testing.T) {
	vias := make(map[string]*models.Via)
	lineas := []string{
		"Avispa, 6b, long:3",
		"Lepto, 7c, 34, placa",
	}
	vias, err := CrearVias(lineas, vias)

	assert.Error(t, err)
	assert.Nil(t, vias)
	assert.Equal(t, "longitud inválida", err.Error())
}

func TestLongitudDeViaNoPuedeSerNegativa(t *testing.T) {
	vias := make(map[string]*models.Via)
	lineas := []string{
		"Avispa, 6b, -1564",
	}
	vias, err := CrearVias(lineas, vias)

	assert.Error(t, err)
	assert.Nil(t, vias)
	assert.Equal(t, "la longitud no puede ser negativa", err.Error())
}

func TestLongitudDeViaMayorDe2000(t *testing.T) {
	vias := make(map[string]*models.Via)
	lineas := []string{
		"Avispa, 6b, 2500",
	}
	vias, err := CrearVias(lineas, vias)

	assert.Error(t, err)
	assert.Nil(t, vias)
	assert.Equal(t, "la longitud no puede ser mayor de 300", err.Error())
}
