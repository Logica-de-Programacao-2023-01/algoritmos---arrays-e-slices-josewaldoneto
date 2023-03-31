package main

import "fmt"

// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos.
// Em seguida, imprima o Slice e a soma dos valores armazenados.

func main() {
	var soma, tamanho, x int
	var slic []int

	fmt.Print("Qual tamanho do slice?")
	fmt.Scan(&tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)
		slic = append(slic, x)
		soma += x
	}

	fmt.Println("O slice é de: ", slic)
	fmt.Println("A soma é de ", soma)
}
