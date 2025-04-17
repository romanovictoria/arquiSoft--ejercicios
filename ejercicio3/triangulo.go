package main

import (
	"fmt"
	"math"
)

func main() {

	var lado1, lado2 float64

	//entrada de datos
	fmt.Print("Ingrese datos del lado1: ")
	fmt.Scanln(&lado1)
	fmt.Print("Ingrese datos del lado2: ")
	fmt.Scanln(&lado2)

	//area
	area := (lado1 * lado2) / 2
	fmt.Printf("El area es: %.2f \n", area)

	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa
	fmt.Printf("Perimetro: %.2f", perimetro)

}
