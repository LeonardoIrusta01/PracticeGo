package main

import "fmt"

/* Listas de Tareas */
type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTasks(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTasks(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

/* Tareas */
type Task struct {
	name     string
	desc     string
	complete bool
}

func (t *Task) toprint() {
	fmt.Printf("Name: %s \nDescription: %s \nComplete: %t\n", t.name, t.desc, t.complete)
}

func (t *Task) tocomplete() {
	t.complete = true
}

func main() {
	t1 := Task{
		name:     "Curso de Go",
		desc:     "Completar curso de Go este mes",
		complete: false,
	}

	t2 := Task{
		name:     "Curso de HTML",
		desc:     "Completar curso de HTML esta semana",
		complete: true,
	}
	// t1.toprint()
	// t2.toprint()

	/* Creamos la lista de tareas para luego agregarlas */
	list := TaskList{}
	list.appendTasks(&t1)
	list.appendTasks(&t2)

	// fmt.Println(list)

	t3 := Task{
		name:     "Curso de CSS",
		desc:     "Completar curso de CSS esta semana",
		complete: false,
	}

	list.appendTasks(&t3)

	list.removeTasks(1)
	/* Recorremos nuestro slice para poder imprimir los datos del struc */
	for i, task := range list.tasks {
		fmt.Println(i, task.name)
	}

}
