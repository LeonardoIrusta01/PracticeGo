package main

import "fmt"

/* Con los tres puntos damos a entender que no se sabe la cantidad de valores que vamos a recibir */
func sumar(nombre string, numeros ...int) (string, int) {

	mensaje := fmt.Sprintf("La suma de %s es:", nombre)
	var total int

	for _, num := range numeros {
		total += num
	}

	return mensaje, total

}

func main() {
	mensaje, result := sumar("Leonardo", 10, 20, 40, 70, 60)

	fmt.Println(mensaje, result)
}
