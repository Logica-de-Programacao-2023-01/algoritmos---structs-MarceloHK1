package main

type triangulo struct {
	base   float64
	altura float64
}

func areadotri(a triangulo) float64 {
	area := (a.base * a.altura) / 2.0
	return area
}
