package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {
	arrayCadena := strings.Split(cadena, "")
	arrayInvertida := make([]string, 0)

	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}

	// fmt.Println("Array spliteado", arrayCadena)
	// fmt.Println("Array invertido", arrayInvertida)

	return strings.Join(arrayInvertida, "")
}

func esPalindormo(palabra string) bool {
	/** Paquete Strings
	* Con el .Replace si o si debo acalarar al final todas las cosas que quiero remplazar
	* Con el .ReplaceAll remplazara absolutamente todo, es decir, no hace falta aclarar cuantos debe reemplazar
	 */
	palabra = strings.ToLower(palabra)

	// palabra = strings.Replace(palabra, "Z", "S", 2)
	palabra = strings.ReplaceAll(palabra, " ", "")

	palabraInvertida := reverse(palabra)

	return palabra == palabraInvertida
}

func main() {
	var dato string

	fmt.Print("Ingrese una palabra: ")
	fmt.Scanln(&dato)

	if esPalindormo(dato) {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

}
