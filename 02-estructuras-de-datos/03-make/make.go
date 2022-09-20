package main

import "fmt"

func main() {
	/* Slicen vacio */
	numeros := make([]int, 0, 3)

	numeros = append(numeros, 400)

	fmt.Println("Slicen de numeros:", numeros, "Longitud:", len(numeros), "Capacidad:", cap(numeros))

}
