package main

import (
	"fmt"
)

type aluno struct {
	nome  string
	idade int
	notas []float64
}

func add(a aluno) []float64 {
	var nota float64
	fmt.Print("fale uma nota: ", nota)
	fmt.Scan(&nota)

	a.notas = append(a.notas, nota)
	return a.notas
}
func tirar(a aluno) []float64 {
	var posiçao int
	fmt.Print("digite uma posiçao: ", posiçao)
	fmt.Scan(&posiçao)

	a.notas = append(a.notas[:posiçao], a.notas[posiçao+1:]...)
	return a.notas
}
func media(a aluno) float64 {
	soma := 0.0
	for _, i := range a.notas {
		soma += i
	}
	medias := soma / float64(len(a.notas))
	return medias
}
func infoaluno(a aluno) {
	soma := 0.0
	for _, i := range a.notas {
		soma += i
	}
	medias := soma / float64(len(a.notas))
	fmt.Printf("nome: %s, idade: %d, media: %2f\n", a.nome, a.idade, medias)

}
