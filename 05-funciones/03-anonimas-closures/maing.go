package main

import (
	"fmt"
	"strings"
)

/* Closures */
func repeat(n int) func(cadena string) string {

	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	/* Funcion anonima */
	// func() {
	// 	fmt.Println("Hola desde la funcion anonima")
	// }()
	myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola, %s desde la funcion anonima", nombre)
	}

	fmt.Println(myfunc("Leonardo"))
	/* ---------------------------------------------------------------------------------------------------- */

	/* Primero ejecutamos y llamamos a la funcion padre y le damos los valores a usar */
	repeat3 := repeat(3)
	/* Luego llamamos la funcion hija y le pasamos los datosa usar */
	fmt.Println(repeat3("Hola"))
	fmt.Println(repeat3("Mundo"))

	repeat5 := repeat(5)
	fmt.Println(repeat5("Leonardo"))
	fmt.Println(repeat5("Irusta"))
}
