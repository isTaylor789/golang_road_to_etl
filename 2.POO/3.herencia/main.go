package main

func main() {
	// Crear un cliente
	customer := Customer{
		ID:   1,
		Name: "Juan Pérez",
	}

	// Crear algunos ítems
	items := []Item{
		{"Laptop", 1500.00, 1},
		{"Mouse", 25.50, 2},
		{"Keyboard", 45.00, 1},
	}

	// Crear una factura
	invoice := Invoice{
		Customer: customer,
		Items:    items,
	}

	// Imprimir los detalles de la factura
	invoice.PrintDetails()
}
