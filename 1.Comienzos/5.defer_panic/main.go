package main

import (
	"fmt"
	"os"
)

func main() {

	// crear un archivo
	cli("Crear un archivo")
	createFile()
	fmt.Println("Creacion de archivos")

	// funcion panic
	cli("Panic")
	panico(12, 3)

}

func cli(str string) {
	fmt.Println("--------------- ", str, " -----------------------")
}

func createFile() {

	file, err := os.Create("best.txt")
	if err != nil {
		fmt.Printf("Error al crear el archivo %v", err)
		return
	}
	defer file.Close() //se ejecuta al final de la funcion, modo cola

	_, err = file.Write([]byte("fastfetch best command in linux"))
	if err != nil {
		fmt.Printf("Error al escribir el archivo %v", err)
		return
	}
}

func panico(numerador, divisor int) {

	// metodo recover para seguir con la ejecucion
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperandose despues del fallo", r)
		}
	}()

	if divisor == 0 {
		panic("Error, no se puede dividir entre 0!!!!!!!!")
	}

	result := 0.0

	result = float64(numerador) / float64(divisor)

	fmt.Println("Resultado de la division: ", result)
}
