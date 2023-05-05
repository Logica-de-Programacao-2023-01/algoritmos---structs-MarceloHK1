package main

type viagem struct {
	origem  string
	destino string
	data    int
	preço   float64
}

func maiscara(a []viagem) viagem {
	cara := a[0]

	for _, i := range a {
		if i.preço > cara.preço {
			cara = i
		}
	}
	return cara
}
