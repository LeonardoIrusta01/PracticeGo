package mensajes

import "fmt"

func Hello() {
	fmt.Println("Hola desde el paquete mensaje")
}

const mensaje = "Hola desde mi constante"

func funcionPrivada() {
	fmt.Println("Hola desde la funcion privada")
}

func Imprimir() {
	fmt.Println(mensaje)
	funcionPrivada()
}
