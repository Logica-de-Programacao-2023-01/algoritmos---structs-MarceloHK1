package main

import "fmt"

type filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliaçoes []int
}

func adicionar(a filme) []int {
	var nota int
	fmt.Print("fale uma nota: ", nota)
	fmt.Scan(&nota)

	a.avaliaçoes = append(a.avaliaçoes, nota)
	return a.avaliaçoes
}
func remover(a filme) []int {
	var posiçao int
	fmt.Print("digite uma posiçao: ", posiçao)
	fmt.Scan(&posiçao)

	a.avaliaçoes = append(a.avaliaçoes[:posiçao], a.avaliaçoes[posiçao+1:]...)
	return a.avaliaçoes
}
func mediadasavaliaçoes(a filme) float64 {
	soma := 0.0
	for _, i := range a.avaliaçoes {
		soma += float64(i)
	}
	medias := soma / float64(len(a.avaliaçoes))
	return medias
}
func infofilme(a filme) {
	soma := 0.0
	for _, i := range a.avaliaçoes {
		soma += float64(i)

		medias := soma / float64(len(a.avaliaçoes))
		fmt.Printf("nome: %s, idade: %d, media: %2f\n", a.avaliaçoes, a.avaliaçoes, medias)

	}
}
