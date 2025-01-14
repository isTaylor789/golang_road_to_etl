package main

import "fmt"

func main() {
	// declaracion de variables
	var dog string
	dog = "perro"

	// delcaracion de variables en una sola linea
	var cat, tiger string = "gato", "tigre"

	// tipado dinamico(solo una vez, si se intenta reescribir es invalido)
	//	con declaracion implicita
	// cat, tiger := "gato", "tigre"

	fmt.Println("Animales: ", dog, " ", cat, " ", tiger)

	// declaracion de funciones
	constantes()

	dataTypes()

	forecasting()

	operadores()
}

func constantes() {
	const pi = 3.1416

	fmt.Println("Constantes: \n", pi)
}

/*---------------- Comentarios --------------------*/

// con go doc --all generamos la documentacion

// Tipos de datos
func dataTypes() {
	// tipos de datos,  string, int, etc
	var a bool = true

	// ver las caracteristicas de la variable
	fmt.Printf("Caracteristicas -> Tipo: %T, valor: %v \n", a, a)
}

// Tipos de datos
func forecasting() {
	// sumar dos tipos de datos no tan compatibles
	var b uint8 = 120
	var a uint16 = 200

	// no altera el tipo de dato de la var b
	c := uint16(b) + a

	// ver las caracteristicas de la variable
	fmt.Printf("Forecasting -> Tipo: %T, valor: %v \n", c, c)
}

func operadores() {
	var a, b, c, d int
	var f, g bool

	//incremento
	a += 1

	// comparacion
	b = 12
	f = b > a

	//incremento y suma
	c++
	c += 2

	// logicos
	d = 12*23 - 1
	g = 12 > 1 && b < d

	fmt.Println("Resultaddos de op. logicos: ", a, b, c, d, f, g)
}
