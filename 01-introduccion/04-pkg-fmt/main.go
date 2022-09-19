package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	/* El fmt.Print sirve para imprimir por pantalla pero no deja ni espacios ni saltos de linea, es decir imprime todo junto*/
	fmt.Print(hola)

	/* El fmt.Println sirve para imprimir con un salto de linea*/
	fmt.Println(mundo)

	var nombre string = "Leonardo"
	var edad int = 21

	/* Usamos %s para string y %d para enteros*/
	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)

	/* Usamos el %v cuando no sabemos que datos vamos a imprimir por pantalla*/
	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)

	/* El fmt.SprintF formatea todo lo que querramos almacenar dentro de una variable.*/
	mensaje := fmt.Sprintf("Hola, %s y su edad es %d", nombre, edad)
	fmt.Println(mensaje)

	/* Con %T nos permite saber el tipo de datos que querramos imprimir por pantalla */
	fmt.Printf("nombre: %T \n", nombre)

	/* Con fmt.Print me permite decirle al usario que ingrese un valor por la terminal */
	fmt.Print("Ingrese otro nombre: ")
	/* Con fmt.Scanln puedo capturar un valor que no contenga esapcios */
	fmt.Scanln(&nombre)

	fmt.Println("Otro nombre: ", nombre)
}
