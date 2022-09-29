package main

import "fmt"

func main() {
	nombres := [...]string{"Leonardo", "Mica", "Tomi", "Lucho"}

	/* Ciclo normal */
	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}

	/* For each
	* El primer valor del for each es para recuperar el indice.
	* El segundo valor del for each es para recuperar los elementos.
	* Si no queremos obtener los indices, simplemente podemos poner "_"
	 */
	for indice, elementos := range nombres {
		fmt.Println(indice, elementos)
	}
}
