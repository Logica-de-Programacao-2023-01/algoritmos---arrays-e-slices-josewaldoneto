package main

import "fmt"

func main() {
	slice := make([]int, 0, 5)

	for len(slice) < 5 {
		var numero int
		fmt.Print("Digite um número inteiro: ")
		fmt.Scanln(&numero)

		if !contains(slice, numero) {
			slice = append(slice, numero)
		}
	}

	fmt.Println("Slice resultante:", slice)
}

func contains(slice []int, numero int) bool {
	for _, valor := range slice {
		if valor == numero {
			return true
		}
	}
	return false
}
