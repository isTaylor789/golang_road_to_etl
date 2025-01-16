package main

import "fmt"

func main() {
	// Crear una troca usando el constructor
	ford := NewTroca("Raptor", 2024, "Ford Raptor", 4, true, 50500)

	// Imprimir detalles
	ford.PrintStruct()

	// Obtener el precio
	fmt.Printf("Precio inicial: %.2f\n", ford.GetPrice())

	// Cambiar el precio
	ford.SetPrice(48000)
	fmt.Printf("Nuevo precio: %.2f\n", ford.GetPrice())

	// Intentar establecer un precio negativo
	ford.SetPrice(-1000)
}
