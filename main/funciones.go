package main

import "fmt"

func suma(a, b int) int {

	return a + b
}

func presentaResultado(a, b int) {
	fmt.Println("La suma de ", a, " y ", b, " es:", suma(a, b))
}

func main() {
	presentaResultado(5, 8)
}
