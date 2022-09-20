package main

import "fmt"

func suma(num1, num2 int) int {
	return num1 + num2
}

func main() {
	num1 := 0
	num2 := 0

	fmt.Print("Ingrese el primer numero: ")
	fmt.Scanln(&num1)

	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scanln(&num2)

	fmt.Println(suma(num1, num2))
}
