package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var result int

	for _, valor := range valores {
		num, e := strconv.Atoi(valor)

		if e != nil {
			fmt.Println(e)
			fmt.Println("Error: tiene que ingresar un numero entero")
			fmt.Println("Solo debe usar el operador +!!")
		} else {
			result += num
		}
	}

	return result

}

func main() {
	var expresion string
	var result int

	fmt.Print("Ingrese una expreccion: ")
	fmt.Scanln(&expresion)

	result = sumar(expresion)

	fmt.Printf("La suma total es => %d\n", result)
}
