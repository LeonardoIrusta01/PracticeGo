package main

import (
	"fmt"
	"math"
)

/* */
type Geometria interface {
	area() float64
	perimetro() float64
}

type Cuadrado struct {
	lado float64
}

func (cua *Cuadrado) area() float64 {
	return cua.lado * cua.lado
}
func (cua *Cuadrado) perimetro() float64 {
	return cua.lado * 4
}

type Circulo struct {
	radio float64
}

func (cir *Circulo) area() float64 {
	return math.Pi * (cir.radio * cir.radio)
}
func (cir *Circulo) perimetro() float64 {
	return 2 * math.Pi * cir.radio
}

func medidas(g Geometria) {
	fmt.Println(g)
	fmt.Print(g.area())
	fmt.Println(g.perimetro())
}
func main() {
	cua1 := Cuadrado{lado: 4}
	cir1 := Circulo{radio: 5}

	medidas(&cua1)
	medidas(&cir1)
}
