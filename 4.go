package main

import "fmt"

type playlist struct {
	nome    string
	musicas []musica
}
type musica struct {
	titulo  string
	artista string
	duraçao float64
}

func dadosdamusica(a playlist, b musica) {
	tempototal := 0.0
	for _, i := range a.musicas {
		tempototal += i.duraçao
	}
	fmt.Printf("nome da playlist: ", a.nome)
	fmt.Printf("duraçao total da playlist: ", tempototal)
	fmt.Printf("titulo: %s, artista: %s, duraçao: %d", b.titulo, b.artista, b.duraçao)
}
