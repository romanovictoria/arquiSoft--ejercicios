package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Tarea struct {
	nombre      string
	descripcion string
	completada  bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (l *ListaTareas) agregarTarea(t Tarea) {

	l.tareas = append(l.tareas, t)
	fmt.Println("Tarea agregada correctamente")

}

func (l *ListaTareas) marcarCompletado(indice int) {

	l.tareas[indice].completada = true

}

func (l *ListaTareas) eliminarTarea(indice int) {

	l.tareas = append(l.tareas[:indice], l.tareas[indice+1:]...)

}

func (l *ListaTareas) editarTarea(indice int, t Tarea) {

	l.tareas[indice] = t

}

func (l *ListaTareas) MostrarTareas() {

	fmt.Println("Listado de tareas: ")
	fmt.Println("==============")

	for ind, t := range l.tareas {

		fmt.Printf(" %d. %s, %s, - completado %t\n", ind, t.nombre, t.descripcion, t.completada)
		fmt.Println("=============")

	}
}

func (lista *ListaTareas) obtenerIndice(action string) (int, error) {

	fmt.Printf("Ingresar indice de la tarea que desea %s: \n", action)
	var indice int
	fmt.Scanln(&indice)

	if indice < 0 || indice >= len(lista.tareas) {
		return -1, errors.New("Indice fuera de rango")
	}
	return indice, nil
}

func displaymenuSelectOption() int {
	fmt.Print(
		"Seleccione una opcion: \n",
		" 1.Agregar tarea \n",
		" 2.Marcar tarea como completada \n",
		" 3.Editar Tarea \n",
		" 4.Eliminar Tarea \n",
		" 5. Salir \n",
	)
	var opcion int
	fmt.Scanln(&opcion)
	return opcion
}

func crearTarea(action string) Tarea {
	fmt.Printf("Ingrese nombre de la tarea que desea %s: \n", action)
	leer := bufio.NewReader(os.Stdin)
	nombre, _ := leer.ReadString('\n')

	fmt.Printf("Ingrese descripcion de la tarea que desea %s: \n ", action)
	descripcion, _ := leer.ReadString('\n')

	return Tarea{nombre: nombre, descripcion: descripcion}

}

func main() {
	lista := ListaTareas{}
	for {
		opcion := displaymenuSelectOption()
		switch opcion {
		case 1:
			//agregar tarea
			t := crearTarea("agregar")
			lista.agregarTarea(t)

		case 2:
			//marcar tarea como completada
			indice, err := lista.obtenerIndice("marcar como completada")
			if err != nil {
				fmt.Println("error: ", err)
				break
			}
			lista.marcarCompletado(indice)

		case 3:
			//editar una tarea
			indice, err := lista.obtenerIndice("editar")
			if err != nil {
				fmt.Println("error: ", err)
				break
			}
			t := crearTarea("editar")
			lista.editarTarea(indice, t)

		case 4:
			//Eliminar tarea
			indice, err := lista.obtenerIndice("eliminar")
			if err != nil {
				fmt.Println("error: ", err)
				break
			}
			lista.eliminarTarea(indice)
		}
		lista.MostrarTareas()
	}
}
