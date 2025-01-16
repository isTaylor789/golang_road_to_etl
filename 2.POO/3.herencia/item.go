package main

import "fmt"

// Item representa un producto o servicio en la factura.
type Item struct {
	Description string
	Price       float64
	Quantity    int
}

// Total calcula el total para un ítem.
func (i Item) Total() float64 {
	return i.Price * float64(i.Quantity)
}

// PrintDetails imprime los detalles del ítem.
func (i Item) PrintDetails() {
	fmt.Printf("Item: %s, Price: %.2f, Quantity: %d, Total: %.2f\n",
		i.Description, i.Price, i.Quantity, i.Total())
}
