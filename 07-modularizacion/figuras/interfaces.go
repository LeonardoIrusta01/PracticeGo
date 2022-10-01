package figuras

import "fmt"

type Geometria interface {
	area() float64
	perimetro() float64
}

func Medidas(g Geometria) {
	fmt.Println("Medidas:", g)
	fmt.Print("Area:", g.area())
	fmt.Println("\n Perimetro:", g.perimetro())
}
