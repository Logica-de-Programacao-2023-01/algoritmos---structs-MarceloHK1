package main

import "math"

type circulo struct {
	raio float64
}

func areadocirculo(a circulo) float64 {
	pi := math.Pi
	area := pi * a.raio * a.raio
	return area
}
