package main

import "fmt"

/* Struc persona (class) */
type Persona struct {
	nombre string
	edad   int
}

/* Métodos
* La variables sirve para almacenar los datos al cual donde estamos apuntando con el sistema de punteros
* Sin el "*" no vamos a poder acceder a los valores de referencia de nuestro struc y no nos mande las copias
 */
func (p *Persona) printeo() {
	fmt.Printf("Nombre: %s \nEdad: %d \n", p.nombre, p.edad)
}

func (p *Persona) edit(nombre string) {
	p.nombre = nombre
}

/* Herencia
* Se pueden heredar métodos y estructuras,
* pero no se puede instanciar de la misma manera si fuera propia.
 */
type Empleado struct {
	Persona
	sueldo float64
}

func main() {

	p1 := Persona{"Leonardo", 21}
	p1.printeo()

	/* Modificar valores */
	p1.edit("Gabriel")
	p1.printeo()

	p2 := Persona{
		nombre: "Mica",
		edad:   18,
	}
	p2.printeo()

	empl1 := Empleado{
		sueldo: 500,
	}
	empl1.nombre = "Pedro"
	empl1.edad = 30
	empl1.printeo()
	fmt.Println(empl1)
}
