package main

import "fmt"

func main() {

	// sentencia if
	cli("if, if-else.if, if-else")
	numType(12)

	// switch
	cli("switch")
	switchCase("cow")

	//for
	cli("for")
	sumN(100)

	// for con arrays
	cli("Arrays")
	escalar(3)

	// for para maps
	cli("Map")
	mapPrint()

	// for para strings
	cli("Strings")
	strPrint("Hallo Word!!!")

}

func cli(str string) {
	print("\n--------------- ", str, " -----------------------\n")
}

// sentencia if
func numType(num float32) {
	fmt.Println("Numero: ", num)

	if 0 > num {
		fmt.Println("Numero negativo")
	} else if num > 0 {
		fmt.Println("Numero positivo")
	} else {
		fmt.Println("El numero es cero")
	}
}

// switch
func switchCase(message string) {
	fmt.Println("Sentencia: ", message)

	switch message {
	case "cow":
		fmt.Println("muu")
	case "cat":
		fmt.Println("meow")
	default:
		fmt.Println("Is a animal")
	}

}

// for
func sumN(num int) {
	count := 0

	for i := 1; i <= num; i++ {
		count += i
	}

	fmt.Println("Resultado de la sumatoria: ", count)

	// otra forma
	// i := 1
	// for i <= num {
	//	count += i
	//	i++
	// }

}

// Escalar un array
func escalar(nm int) {

	num := []float32{200, 12.5, 99, 34.12314}
	fmt.Println("Array original: ", num)
	fmt.Println("Escalar: ", nm)

	for i := range num {
		num[i] *= float32(nm)
	}
	fmt.Println("Resultado: ", num)
}

// Imprimir un mapa
func mapPrint() {
	animals := map[string]string{
		"panda": "🐼",
		"tiger": "🐯",
		"cat":   "😺",
		"cow":   "🐮",
	}

	for k, v := range animals {
		fmt.Println("key: ", k, " value: ", v)
	}
}

func strPrint(str string) {

	for _, v := range str {
		fmt.Println(string(v))
	}
}
