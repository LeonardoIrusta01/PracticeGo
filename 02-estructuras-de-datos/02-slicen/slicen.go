package main

import "fmt"

func main() {
	/* Slicen */
	numeros := []int{1, 2, 3}

	fmt.Println("Slicen de numeros:", numeros, "Y su longitud es de:", len(numeros))

	/* Agregar datos */

	numeros = append(numeros, 4) // El append en su primer parametro recibe el Slicen a manipular y en su segundo parametro los valores a agregar.
	numeros = append(numeros, 5)

	fmt.Println("Slicen de numeros:", numeros, "Y su longitud es de:", len(numeros))

	/* Sub Slicen */

	subSlicen := numeros[:2]

	numeros[0] = 100

	fmt.Println("subSlicen:", subSlicen, "Y su longitud es de:", len(subSlicen))
	fmt.Println("Papa Slicen:", numeros)

	/* Capacidad */
	meses := []string{
		"Enero",
		"Febrero",
		"Marzo",
	}

	fmt.Printf("Longitud: %v, Capacidad: %v, Memoria: %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abirl")

	fmt.Printf("Longitud: %v, Capacidad: %v, Memoria: %p \n", len(meses), cap(meses), meses)
}
