package main

import (
	"fmt"
)

func main() {

	// punteros
	num := 23

	cli("Punteros")
	punteros(num)
	unReference(num)
	pointerChange(&num)

	//arrays
	arr := [3]int{1, 2, 3}
	arrSlice := []int{1, 2, 4, 4}

	cli("Arrays")
	arrays(arr)
	slices(arrSlice)
	arrayOps(arrSlice)

	// crear mapas
	cli("Maps")

	maps()

	//Structs
	cli("Structs")

	structs()

}

func cli(str string) {
	print("--------------- ", str, " -----------------------")
}

func punteros(num int) {
	println("\nDireccion del puntero %v", &num)
}

func unReference(num int) {
	unRef := &num
	println("Unreference %s", *unRef)
}

func pointerChange(num *int) {
	change := num

	// cambiar el numero en memoria
	*change = 12

	println("Nuevo valor del puntero: ", *change)
}

// arrays
func arrays(array [3]int) {
	fmt.Println("\nArrays: ", array)
}

func slices(slice []int) {
	fmt.Println("Slice: ", slice)
}

func arrayOps(array []int) {

	// extraer desde el indice hasta un punto del array - 1
	extrac := array[1:3]

	// go por defecto escala el array si se usan metodos como append
	extrac = append(extrac, 5, 5, 5, 5)

	// info del array
	fmt.Println("Array original: ", array)
	fmt.Println("Size: ", len(array))
	fmt.Println("Secuencia extraida: ", extrac)
	fmt.Println("Capacidad del segundo array: ", cap(extrac))
}

// maps
func maps() {

	animals := map[string]string{
		"cat":     "meow",
		"wolf":    "laughs in furry",
		"chicken": "pio",
		"crab":    "money, money",
	}

	fmt.Println("\n", animals)

	// eliminar un una feature del maps
	delete(animals, "wolf")
	fmt.Println("Map post delete: ", animals)

	// crear una nueva caracteristica en el mapa
	animals["gorilla"] = "ugh ugh"
	fmt.Println(animals)

}

func structs() {

	type car struct {
		name   string
		model  string
		wheels int
		ia     bool
		price  float32
	}

	troca := car{
		name:   "Hilux 2024",
		model:  "Hilux",
		wheels: 4,
		ia:     false,
		price:  50000,
	}
	fmt.Printf("\n%+v\n", troca)

	// declaracion implicita
	fordsito := car{
		"Raptor 2023",
		"Raptor",
		4,
		false,
		51000,
	}
	fmt.Printf("\n%+v\n", fordsito)

	// modificar struct
	p := &fordsito
	p.ia = true
	fmt.Printf("\n%+v\n", p)

}
