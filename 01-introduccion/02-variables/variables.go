package main

import "fmt"

func main() {

	/* Declaracion de variables, primero seteo el tipo de dato que voy a usar y luego le asigno dicho valor. */
	var nombre1 string
	nombre1 = "Leonardo"

	/* Seteo el tipo de dato y a su vez le asigno el valor que va a tener */
	var nombre2 string = "Gabriel"

	/* Forma corta de declarar variables*/
	edad := 21

	fmt.Println("Mi nombre es:", nombre1, nombre2, "y mi edad es:", edad)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	const pi = 3.14
	fmt.Println("Valor de pi:", pi)
}
