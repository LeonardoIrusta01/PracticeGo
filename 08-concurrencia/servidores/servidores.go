package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarSer(servidor string, canal chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		// fmt.Println(servidor, "No se encuentra disponible el servidor!")
		canal <- servidor + "No se encuentra disponible el servidor!"
	} else {
		// fmt.Println(servidor, "Esta funcionando correctamente")
		canal <- servidor + "Esta funcionando correctamente"
	}
}

func main() {
	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"http://oregoom.com",
		"http://www.udemy.com/",
		"http://www.youtube.com/",
		"http://www.facebook.com/",
		"http://www.google.com/",
	}

	for _, serv := range servidores {
		go revisarSer(serv, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	timepoTotal := time.Since(inicio)

	fmt.Println("Tiempo total:", timepoTotal)
}
