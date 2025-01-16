package main

import "fmt"

// Invoice representa una factura.
type Invoice struct {
	Customer Customer
	Items    []Item
}

// Total calcula el total de la factura.
func (inv Invoice) Total() float64 {
	total := 0.0
	for _, item := range inv.Items {
		total += item.Total()
	}
	return total
}

// PrintDetails imprime los detalles de la factura.
func (inv Invoice) PrintDetails() {
	fmt.Println("Invoice Details:")
	inv.Customer.PrintDetails()
	for _, item := range inv.Items {
		item.PrintDetails()
	}
	fmt.Printf("Total Invoice Amount: %.2f\n", inv.Total())
}
