package main

import "fmt"

func SumarHasta(valor int) int {
	NumeroAnterior := 0
	for i := 0; i < valor; i++ {
		NumeroAnterior += i
	}
	return NumeroAnterior
}

func main() {
	var resultado int = SumarHasta(10)
	fmt.Println(resultado)
	Ejemplo(55)
}
