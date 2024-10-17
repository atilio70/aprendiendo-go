package main

import "fmt"

var funciones = map[string]func(int, int) int{
	"suma":       func(a, b int) int { return a + b },
	"resta":      func(a, b int) int { return a - b },
	"multiplica": func(a, b int) int { return a * b },
	"divide":     func(a, b int) int { return a / b },
}

func presentaResultado(operacion string, a int, b int) {
	f, exists := funciones[operacion]
	if !exists {
		fmt.Println("Operación no valida")
	}
	fmt.Println("Para a= ", a, " y b= ", b, "la ", operacion, "resultante es: ", f(a, b))
}
func main() {
	presentaResultado("suma", 5, 8)
	presentaResultado("resta", 12, 8)
	presentaResultado("multiplica", 3, 8)
	presentaResultado("divide", 8, 2)
}
