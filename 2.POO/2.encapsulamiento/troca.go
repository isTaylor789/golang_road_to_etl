package main

import (
	"fmt"
	"reflect"
)

// troca representa una camioneta
type troca struct {
	name   string  // privado: accesible solo dentro del paquete
	year   int     // privado
	model  string  // privado
	wheels int     // privado
	ia     bool    // privado
	price  float64 // privado
}

// NewTroca es un constructor que crea una troca con valores predeterminados.
func NewTroca(name string, year int, model string, wheels int, ia bool, price float64) *troca {
	return &troca{
		name:   name,
		year:   year,
		model:  model,
		wheels: wheels,
		ia:     ia,
		price:  price,
	}
}

// GetPrice devuelve el precio de la troca.
func (t *troca) GetPrice() float64 {
	return t.price
}

// SetPrice establece un nuevo precio para la troca.
func (t *troca) SetPrice(price float64) {
	if price < 0 {
		fmt.Println("Error: El precio no puede ser negativo.")
		return
	}
	t.price = price
}

// PrintStruct imprime los detalles de la troca.
func (t *troca) PrintStruct() {
	val := reflect.ValueOf(*t)
	typ := reflect.TypeOf(*t)

	fmt.Println("Detalles de la troca:")
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i) // Tipo del campo
		value := val.Field(i) // Valor del campo
		fmt.Printf("%s: %v\n", field.Name, value)
	}
}
