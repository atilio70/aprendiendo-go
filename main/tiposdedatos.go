package main

import "fmt"

func condicionales() {
	//var nombrePersona string = "Leonardo"
	//var apellidoPersona = "Roggero"
	//segundoNombre := "Jose"

	//fmt.Println("Hola mundo. Soy " + nombrePersona + " " + segundoNombre + " " + apellidoPersona)

	// parte  numerica
	//var asoActual int16 = 2024
	//edad := 56

	//fmt.Println("El año actual es ", asoActual)
	//fmt.Println("Tengo ", edad)

	//arreglos y slices y rangos

	//var listaFrutas = [4]string{"pera", "Naranja"}
	//fmt.Println(listaFrutas[0])

	//listaPaises := [3]string{"Peru", "Brasil", "Venezuela"}
	//listaPaises := []string{"Peru", "Brasil", "Venezuela"}
	//fmt.Println(listaPaises)
	//listaPaises = append(listaPaises, "Chile")
	//fmt.Println(listaPaises)

	//listaPaises2 := listaPaises[:2]
	//fmt.Println(listaPaises2)

	//mapas,bucles y condicionales

	//	var miMapa = map[string]string{
	//		"Colombia":  "Bogota",
	//		"Venezuela": "Caracas",
	//		"Brasil":    "Brasilia",
	//		"Chile":     "Santiago",
	//	}

	//	fmt.Println("La capital de Venezuela es:", miMapa["Venezuela"])
	//	miMapa["Argentina"] = "Buenos Aires"
	//	fmt.Println("El mapa de paises es ", miMapa)
	//	delete(miMapa, "Venezuela")
	//	fmt.Println("El mapa de paises es ", miMapa)

	var edad int = 13
	var permiso bool = true

	if edad < 18 && permiso {
		fmt.Println("El menor de edad puede viajar solo")
	} else if edad < 18 && !permiso {
		fmt.Println("El menor no puede viajar solo")
	} else {
		fmt.Println("La persona puede viajar")
	}

	fmt.Println("Fin del programa")
}
