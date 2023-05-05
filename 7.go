package main

import "fmt"

type animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

func mudarosom(a animal) {
	var novosom string
	fmt.Print("animal: ", a.nome)
	fmt.Print("o novo som é: ", novosom)
	fmt.Scan(&novosom)
}
func info(a animal) {
	fmt.Printf("nome: %s, especie: %s, idade: %d, som: %s", a.nome, a.especie, a.idade, a.som)
}
