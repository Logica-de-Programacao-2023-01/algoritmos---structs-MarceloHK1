package main

import (
	"fmt"
)

type pessoa struct {
	nome     string
	idade    int
	endereço struct {
		rua    string
		numero int
		cidade string
		estado string
	}
}

func dados(a pessoa) {
	fmt.Printf("endereço de: ", a.nome)
	fmt.Printf("%s,%d - %s, %s\n", a.endereço.rua, a.endereço.numero, a.endereço.cidade, a.endereço.estado)
}
