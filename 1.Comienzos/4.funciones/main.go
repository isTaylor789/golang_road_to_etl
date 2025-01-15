package main

import (
	"fmt"
	"strings"
)

func main() {

	// funciones tipadas
	cli("Funciones tipadas")
	sumaResultado := suma(5, 4)
	fmt.Println("Resultado", sumaResultado)

	// return multiples valores
	cli("Multiples valores de retorno")
	upper, lower := sentence("Paco")
	fmt.Println("Resultado: ", upper, lower)

	// funciones anonimas
	cli("Funciones anonimas")
	anon()

	// multiples entradas de vars
	cli("Multiples entradas")
	fmt.Println(mulVars(1, 2, 3, 4, 5, 6, 7, 8, 9))

}

func cli(str string) {
	print("\n--------------- ", str, " -----------------------\n")
}

// func tipada
func suma(a, b int) int {
	return a + b
}

// func multiples
func sentence(str string) (string, string) {
	upper := strings.ToUpper(str)
	lower := strings.ToLower(str)

	return upper, lower
}

// funciones anonimas
func anon() {

	x := func() {
		fmt.Println("Hola soy una funcion anonima")
	}
	x() // ejecutar una func anonima

	func() {
		fmt.Println("Hola soy una funcion anonima AUTO EJECUTANDOSE")
	}()

}

// multiples entradas
func mulVars(nums ...int) int {
	count := 0

	for _, v := range nums {
		count += v
	}

	return count
}
