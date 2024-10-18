package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

/*
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
*/

type Alumno struct {
	Nombre string
	Notas  []float64
}

func leerCSV(nombreArchivo string) ([]Alumno, error) {
	archivo, err := os.Open(nombreArchivo)

	if err != nil {
		return nil, err
	}

	lector := csv.NewReader(archivo)
	lector.Comma = ';'

	registros, err := lector.ReadAll()
	if err != nil {
		return nil, err
	}

	var alumnos []Alumno

	for _, registro := range registros {
		nombre := registro[0]
		var notas []float64
		for _, notaStr := range registro[1:] {
			nota, err := strconv.ParseFloat(notaStr, 64)
			if err != nil {
				return nil, err
			}
			notas = append(notas, nota)
		}
		alumnos = append(alumnos, Alumno{Nombre: nombre, Notas: notas})
	}
	return alumnos, nil
}

func calcularPromedio(notas []float64) float64 {
	var suma float64
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}

func calcularMediaAritmetrica(alumnos []Alumno) float64 {
	var suma float64
	var cantidadNotas int
	for _, alumno := range alumnos {
		for _, nota := range alumno.Notas {
			suma += nota
			cantidadNotas++
		}
	}
	return suma / float64(cantidadNotas)
}

func escribirCSV(nombreArchivo string, alumnos []Alumno) error {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	escritor := csv.NewWriter(archivo)
	defer escritor.Flush()

	for _, alumno := range alumnos {
		promedio := calcularPromedio(alumno.Notas)
		registro := []string{alumno.Nombre, fmt.Sprintf("%.2f", promedio)}
		if err := escritor.Write(registro); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	/*	presentaResultado("suma", 5, 8)
		presentaResultado("resta", 12, 8)
		presentaResultado("multiplica", 3, 8)
		presentaResultado("divide", 8, 2)
	*/
	alumnos, err := leerCSV("notas.csv")
	if err != nil {
		fmt.Println("Error al leer el archivo CSV", err)
		return
	}
	mediaAritmetrica := calcularMediaAritmetrica(alumnos)
	fmt.Printf("La media aritmetrica de todas las notas es:%2.f\n", mediaAritmetrica)

	err = escribirCSV("promedios.csv", alumnos)
	if err != nil {
		fmt.Println("Error al escribir el archivo CSV:", err)
		return
	}
	fmt.Println("Archivo 'Promedios.csv' generado con exito")
}
