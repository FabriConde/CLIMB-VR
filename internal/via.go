package internal

type DificultadEscalada float
const (
	I 	DificultadEscalada = 1.0
	II		= 2.0
	III  	= 3.0
	IV		= 4.0
	V 		= 5.0
	6a		= 6.0
	6aPlus	= 6.20
	6b		= 6.35
	6bPlus	= 6.50
	6c		= 6.70
	6cPlus	= 6.90
	7a		= 7.0
	7aPlus	= 7.20
	7b		= 7.35
	7bPlus	= 7.50
	7c		= 7.70
	7cPlus	= 7.90
	8a		= 8.0
	8aPlus	= 8.20
	8b		= 8.35
	8bPlus	= 8.50
	8c		= 8.70
	8cPlus	= 8.90
	9a		= 9.0
	9aPlus	= 9.20
	9b		= 9.35
	9bPlus	= 9.50
	9c		= 9.75
	9cPlus	= 10
)

type Via struct {
	nombre         string
	dificultad     DificultadEscalada
	longitud       float32

	infoExtra      string
}
