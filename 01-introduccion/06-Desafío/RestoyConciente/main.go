package main

import "fmt"

func cociente(num1, num2 int) int {
	return num1 / num2
}

func resto(num1, num2 int) int {
	return num1 / num2
}

func main() {
	/* Seteo de variables */
	var num1 int = 0.0
	var num2 int = 0.0

	/* Entrada de datos */
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scanln(&num1)

	fmt.Print("Ingrese el primer numero: ")
	fmt.Scanln(&num2)

	/* Llamado a las funciones */
	coc := cociente(num1, num2)
	res := resto(num1, num2)

	/* Muestra de datos por pantalla */
	fmt.Println("El cociente es:", coc)
	fmt.Println("El resto es:", res)

}
