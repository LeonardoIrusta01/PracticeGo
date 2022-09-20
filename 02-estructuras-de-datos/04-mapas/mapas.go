package main

import "fmt"

func main() {
	/* Dentro del los corchetes del map, se le setea el tipo de dato que se va a usar */
	dias := make(map[int]string)

	fmt.Println("Mapa:", dias)

	/* Agregar datos */
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"

	fmt.Println("Dias:", dias)

	dias[7] = "SABADO"

	fmt.Println("Dias:", dias)

	/* Elimnar */
	delete(dias, 1)
	fmt.Println("Dias:", dias, "Longitud:", len(dias))

	/* Nueva mapa */
	estudiantes := make(map[string][]int)

	estudiantes["Alex"] = []int{13, 15, 16}
	estudiantes["Leo"] = []int{16, 19, 21}

	fmt.Println("Estudiantes:", estudiantes)
	fmt.Println("Estudiante:", estudiantes["Leo"][0])
}
