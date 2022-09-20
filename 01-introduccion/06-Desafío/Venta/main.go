package main

import "fmt"

func venta(num1 float64) float64 {
	iva := (num1 * 18) / 100

	return num1 + iva
}

func main() {
	/* Seteo de variables */
	num1 := 0.0

	/* Entrada de datos */
	fmt.Print("Ingrese el valor de la venta: ")
	fmt.Scanln(&num1)

	/* Llamado a las funciones */
	ven := venta(num1)

	/* Muestra de datos por pantalla */
	fmt.Println("El precio de la venta es:", ven)

}
