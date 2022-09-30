package main

import "fmt"

/* Le colocamos un asterico (que funciona como puntero) para que tome el valor de la referencia en memoria */
func modificarValor(cadena *string) {
	fmt.Printf("%p \n", cadena)
	*cadena = "Hola desde la funcion"
}

func main() {
	cadena := "Hola Mundo de Go"
	fmt.Printf("%p \n", &cadena)
	println("Antes de la función:", cadena)

	/* Por ahora estoy pasando la copia. */
	// modificarValor(cadena)

	/* Para poder modificar el valor inicial, pasaremos la referencia de la memoria */
	modificarValor(&cadena) // Por ahora estoy pasando la copia.

	fmt.Println("Despues de la función:", cadena)
}
