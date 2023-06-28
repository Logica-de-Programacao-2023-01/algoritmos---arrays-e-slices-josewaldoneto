package main

import "fmt"

func main() {
	slice := make([]string, 8)

	for i := 0; i < len(slice); i++ {
		fmt.Printf("Digite o valor %d: ", i+1)
		fmt.Scanln(&slice[i])
	}

	var valor string
	fmt.Print("Digite o valor a ser removido: ")
	fmt.Scanln(&valor)

	novoSlice := removeValor(slice, valor)

	fmt.Println("Slice resultante:", novoSlice)
}

func removeValor(slice []string, valor string) []string {
	novoSlice := make([]string, 0)

	for _, elemento := range slice {
		if elemento != valor {
			novoSlice = append(novoSlice, elemento)
		}
	}

	return novoSlice
}
