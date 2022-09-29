package main

import "fmt"

// func calcular(cons float64, des float64) float64 {
// 	/* Operaciones */
// 	montoDescuento := cons - des
// 	igv := montoDescuento * 0.19
// 	totalPagar := montoDescuento + igv

// 	return totalPagar
// }

// fmt.Println("Dato de descuento:", datosDescuento, "Total a pagar:", calcular(descuento, consumo))

func main() {
	/** App para restaurante
	* Descuentos de 10% hasta 100 de consumo
	* Descuentos de 20% hasta 200 de consumo
	* Descuentos de 30% mayor de 200 de consumo
	* igv (iva) 19%
	 */

	var consumo, descuento float64
	var datosDescuento string

	// Entrada de datos
	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {

		if consumo <= 100 {
			/* Descuento del 10% */
			datosDescuento = "10%"
			descuento = consumo * 0.10 // (consumo * 10) / 100

		} else if consumo > 100 && consumo <= 200 {
			/* Descuento del 20% */
			datosDescuento = "20%"
			descuento = consumo * 0.20 // (consumo * 20) / 100
		} else if consumo > 200 {
			/* Descuento del 30% */
			datosDescuento = "30%"
			descuento = consumo * 0.30 // (consumo * 30) / 100
		}

	} else {
		println("Error al ingresar el consumo")
	}

	/* Operaciones */
	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.19
	totalPagar := montoDescuento + igv

	/* Salida de datos */
	fmt.Println("-------- FACTURA DE CONSUMO --------")
	fmt.Println("Descuento que se aplica:", datosDescuento)
	fmt.Println("Consumo:", consumo)
	fmt.Println("Descuento:", descuento)
	fmt.Println("Monto de descuento:", montoDescuento)
	fmt.Println("IGV:", igv)
	fmt.Println("Total a pagar:", totalPagar)
	fmt.Println("-------- FACTURA DE CONSUMO --------")
}
