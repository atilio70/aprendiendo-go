package main

import "fmt"

func presentaSuma(suma func(int, int) int, a, b int) {
	fmt.Println("La suma de ", a, " y ", b, " es:", suma(a, b))
}

func presentaResta(resta func(int, int) int, a, b int) {
	fmt.Println("La resta de ", a, " y ", b, " es:", resta(a, b))
}

func presentaResultado(operacion string, a int, b int) {
	suma := func(a, b int) int {
		return a + b
	}
	resta := func(a, b int) int {
		return a - b
	}

	resultado := 0

	if operacion == "sumar" {
		resultado = suma(a, b)
	} else if operacion == "restar" {
		resultado = resta(a, b)
	}
	fmt.Println("La ", operacion, " de ", a, " y ", b, " es:", resultado)

}

func main() {

	presentaResultado(suma, 5, 8)
}
