package main

import "fmt"

// Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.

func main() {
	arr := [4]float64{26.8, 18.2, 80.6, 34.1}
	var produto float64 = 1
	for _, x := range arr {
		produto *= x
	}
	fmt.Print(produto)
}
