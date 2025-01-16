package main

import (
	"fmt"
	"reflect"
)

type troca struct {
	name   string
	year   int
	model  string
	wheels int
	ia     bool
	price  float64
}

func (t troca) PrintStruct() {
	// Usa reflexión para iterar sobre los campos de la estructura
	val := reflect.ValueOf(t)
	typ := reflect.TypeOf(t)

	fmt.Println("Estructura completa:")
	fmt.Printf("Struct: %+v\n", t)

	text := "Los campos de la estructura son:\n"
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i) // Obtiene el tipo del campo
		value := val.Field(i) // Obtiene el valor del campo
		text += fmt.Sprintf("%s: %v\n", field.Name, value)
	}

	fmt.Println(text)
}

func (t *troca) ChangePrice(price float64) {
	t.price = price
	fmt.Println("New price:", price)
}
