package main

import "fmt"

func main() {
	// array
	var numeros [5]int

	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println("Nuevo arreglo", numeros)
	fmt.Println("Valor unico", numeros[1])

	/* Arrays con valores */
	nombre := [3]string{"Alex", "Leonardo", "Irusta"}

	fmt.Println("Arrays con nombres:", nombre)

	/* Arrays sin longitud definida*/

	colores := [...]string{
		"Rojo",
		"Verde",
		"Azul",
		"Amarillo",
	}

	fmt.Println("Arreglo de colores:", colores, "Y su longitud es de:", len(colores))

	/* Indices definidos */
	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euro",
	}

	fmt.Println("Arreglo de monedas:", monedas, "Y su longitud es de:", len(monedas))

	/* Sub array */
	subMoneda := monedas[0:3]
	subMonedaMax := monedas[:3]
	subMonedaMin := monedas[3:]

	fmt.Println("Arreglo de subMonedas:", subMoneda)
	fmt.Println("Arreglo de subMonedaMax:", subMonedaMax)
	fmt.Println("Arreglo de subMonedaMin:", subMonedaMin)
}
