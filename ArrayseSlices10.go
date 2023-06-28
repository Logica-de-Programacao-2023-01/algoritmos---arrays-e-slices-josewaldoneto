package main

import "fmt"

func main() {
	slice := []int{10, 5, 7, 2, 9, 4, 3, 8, 1, 6}

	minimo, maximo := slice[0], slice[0]

	for _, valor := range slice {
		if valor < minimo {
			minimo = valor
		}
		if valor > maximo {
			maximo = valor
		}
	}

	fmt.Println("Valor mínimo:", minimo)
	fmt.Println("Valor máximo:", maximo)
}
