package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

/* Estructura del estudiante (Relacion 1 a 1)
* Solo puede tener un usuario
* Y un usuario debe ser de un solo estudiante
 */
type Student struct {
	user   User
	codigo string
}

/* El curso va a tener varios videos */
type Curse struct {
	title  string
	videos []Video // Creamos un arreglo de tipo Video, es decir que apunta a la estructura de Video
}

/* Y un video solo puede pertenecer a un curso */
type Video struct {
	title string
	curso Curse // Le especificamos que va a pertenecer a dicho curso
}

func main() {
	/* ----------------------------------------------------------------------------------------------------------- */
	/* Relacion de uno a uno*/
	leo := User{
		nombre: "Leonardo",
		email:  "leonardo@gmail.com",
		activo: true,
	}

	mica := User{
		nombre: "Mica",
		email:  "Mica@gmail.com",
		activo: true,
	}

	leoStudent := Student{
		user:   leo,
		codigo: "351asdftr",
	}

	fmt.Println(mica)
	fmt.Println(leoStudent.user.nombre)
	/* ----------------------------------------------------------------------------------------------------------- */
	/*Relacion de uno a muchos */

	/* Creamos los cursos */
	video1 := Video{
		title: "01-Introduccion",
	}
	video2 := Video{
		title: "02-Instalacion",
	}

	/* Creamos el curso */
	curse1 := Curse{
		title:  "Curso de Go",
		videos: []Video{video1, video2}, // Aclaramos que videos van a pertenecer a este curso
	}

	/* A los videos les definimos a que curso va a estar realcionado */
	video1.curso = curse1
	video2.curso = curse1

	fmt.Println(video1.curso.title)

	for _, video := range curse1.videos {
		fmt.Println(video.title)
	}

}
