package main

import "fmt"

func main() {
	/* Not */
	fmt.Println(!true)

	/* And */
	fmt.Println("True & True:", true && true)
	fmt.Println("False & True:", false && true)
	fmt.Println("False & False:", false && false)

	/* Or */
	fmt.Println("True || True:", true || true)
	fmt.Println("False || True:", false || true)

	/* ------------------------------------------------------ */
	b := 2

	r := b == 2 && !(4 > 2)

	fmt.Println(r)
}
