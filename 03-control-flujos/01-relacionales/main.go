package main

import "fmt"

func main() {
	a := 4
	b := 5

	var r bool

	/* Igualdad */
	r = a == b

	/* %d para integers y %t para booleanos*/
	fmt.Printf("%d es igual que %d? %t \n ", a, b, r)

	/* Distinto */
	r = a != b
	fmt.Printf("%d es distinto de %d? %t \n ", a, b, r)

	/* Mayor que */
	r = a > b
	fmt.Printf("%d es mayor que %d? %t \n ", a, b, r)

	/* Menor que */
	r = a < b
	fmt.Printf("%d es menor que %d? %t \n ", a, b, r)

	/* Mayor o igual */
	r = a >= b
	fmt.Printf("%d es mayor o igual que %d? %t \n ", a, b, r)

	/* Menor o igual */
	r = a <= b
	fmt.Printf("%d es menor o igual que %d? %t \n ", a, b, r)
}
