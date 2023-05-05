package main

type funcionario struct {
	nome    string
	salario float64
	idade   int
}

func aumentar(a funcionario) float64 {
	novosalario := a.salario * 1.1
	return novosalario
}
func diminuir(a funcionario) float64 {
	novosalario := a.salario * 0.9
	return novosalario
}
func tempodeserviço(a funcionario) int {
	tempo := a.idade - 18
	return tempo
}
