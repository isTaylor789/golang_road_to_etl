package main

import "fmt"

// Customer representa un cliente.
type Customer struct {
	ID   int
	Name string
}

// PrintDetails imprime los detalles del cliente.
func (c Customer) PrintDetails() {
	fmt.Printf("Customer ID: %d, Name: %s\n", c.ID, c.Name)
}
