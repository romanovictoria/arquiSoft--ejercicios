package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()

}
func play() {

	numAleatorio := rand.Intn(100)
	var numIngresado int

	for cantIntentos := 0; cantIntentos < 10; cantIntentos++ {

		fmt.Printf("Ingrese un número (intentos restantes %d): ", 10-cantIntentos)
		fmt.Scanln(&numIngresado)

		if numAleatorio == numIngresado {
			fmt.Println("Ganaste!")
			displaymenu()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("No acertaste.El numero es mayor ")
		} else {
			fmt.Println("No acertaste.El numero es menor ")
		}
	}
	fmt.Println("No ganaste”.El numero es: ", numAleatorio)
	displaymenu()

}

func displaymenu() {

	var opcion string

	fmt.Println("Desea jugar nuevamente? (s/n): ")
	fmt.Scanln(&opcion)

	switch opcion {

	case "s":
		play()

	case "n":
		fmt.Println("Gracias por jugar")

	default:
		fmt.Println("Opcion incorrecta. Intente de nuevo: ")
		displaymenu()
	}
}
