package main

import figuras "github.com/LeonardoIrusta01/figurasgo"

func main() {
	cua1 := figuras.Cuadrado{Lado: 10}
	figuras.Medidas(&cua1)
	/*
		cua1 := figuras.Cuadrado{Lado: 4}
		circ1 := figuras.Circulo{Radio: 5}

		figuras.Medidas(&cua1)
		figuras.Medidas(&circ1)
	*/

	/*
		p1 := models.Persona{}
		p1.Constructor("Leonardo", 21)

		fmt.Println(p1)
		p1.SetNombre("Gabriel")
		p1.SetEdad(23)
		fmt.Println(p1)
	*/

}
