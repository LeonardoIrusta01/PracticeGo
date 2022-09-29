package main

import "fmt"

func main() {
	/* Bucle infinito */
	// for {
	// 	fmt.Println("Hola mundo")
	// }

	/* Bucle tipo while */
	numeros := 12312455
	cont := 0
	for numeros > 0 {
		numeros /= 10
		cont++
	}
	fmt.Println("Cantidad de digitos es:", cont)

	/** Bucle for
	* primer valor es donde puedo inicializar una variable.
	* segundo valor es donde planteo la condicion.
	* tercer valor es donde incremento el contador.
	 */

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
