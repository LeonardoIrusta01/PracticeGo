package main

import "fmt"

/* Interface */
type Animal interface {
	move() string
}

type Dog struct{}

func (*Dog) move() string {
	return "Soy un perro y camino"
}

type Fish struct{}

func (*Fish) move() string {
	return "Soy un pez y nado"
}

type Bird struct{}

func (*Bird) move() string {
	return "Soy un pájaro y vuelo"
}

/* Funcion que va a dar uso de la interface */
func moveAnimal(animal Animal) {
	fmt.Println(animal.move())
}

func main() {
	dog1 := Dog{}
	moveAnimal(&dog1)

	fish1 := Fish{}
	moveAnimal(&fish1)

	bird1 := Bird{}
	moveAnimal(&bird1)
}
