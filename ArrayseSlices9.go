package main

import "fmt"

func main() {
	array := [6]float64{}

	var numero float64
	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	for i := 0; i < len(array); i++ {
		array[i] = numero
	}

	fmt.Println("Array resultante:", array)
}
