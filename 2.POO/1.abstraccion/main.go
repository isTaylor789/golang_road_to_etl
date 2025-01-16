package main

func main() {
	ford := &troca{
		"Raptor",
		2024,
		"Ford Raptor",
		4,
		true,
		50500,
	}

	cli("Structs")
	ford.PrintStruct()
	ford.ChangePrice(50000.99)

}

func cli(str string) {
	print("\n--------------- ", str, " -----------------------\n")
}
