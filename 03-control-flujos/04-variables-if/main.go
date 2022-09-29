package main

import (
	"fmt"
)

func main() {
	if nombre, edad := "Leonardo", 34; nombre == "Leonardo" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Printf("Nombre: %s -  Edad: %d \n", nombre, edad)
	}

	/* Obtener valor de mapa */
	users := make(map[string]string)

	users["Leonardo"] = "leo@gmail.com"
	users["Mica"] = "mica@gmail.com"

	if correo, ok := users["Leonardo"]; ok {
		fmt.Println("El correo es:", correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}

}
