package main

import (
	"fmt"
	"os"
)

func main() {

	/* Funcion Recover */
	defer func() {
		/* La funcion recover se encarga de controlar todos los panic() de la aplicación */
		if error := recover(); error != nil {
			fmt.Println("Ups!, al parecer el programana no pudo finalizar")
		}
	}()
	// file, error := os.Open("hola.txt")

	if file, error := os.Open("hoa.txt"); error != nil {
		panic("No fue posible obtener el archivo!")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo!")
			file.Close()
		}()
		/* Generamos un slice para almacenar los datos del archivo */
		contenido := make([]byte, 254)

		/* File.Read los usamos para almacenar los datos dentro de contenido, donde
		nos devuelve la longitud de lo que va a almacenar y un error si llega a ocurrir algo. */
		long, _ := file.Read(contenido)

		/* Convertimos el slice del contenido a string */
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)

	}
	// file.Close()
}
