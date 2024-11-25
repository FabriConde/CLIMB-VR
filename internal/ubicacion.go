package internal

type Ubicacion struct {
	coordenadas Coordenadas
	descripcion string
}

type Coordenadas struct {
	latitud  float64
	longitud float64
}
