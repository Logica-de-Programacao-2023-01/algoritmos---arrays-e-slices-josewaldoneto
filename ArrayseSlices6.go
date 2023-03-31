package main

import "fmt"

// Crie um Array de inteiros com 10 elementos.
// Em seguida, solicite ao usuário que informe um valor e verifique se esse valor está presente no Array.
// Imprima o resultado da busca.

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var buscar int

	fmt.Println("Digite um número: ")
	fmt.Scan(&buscar)
	var encontrado bool

	for _, x := range arr {
		if x == buscar {
			fmt.Println("O elemento está no array.")
			encontrado = true
			break
		}
	}
	if !encontrado {
		fmt.Println("O elemento não está no array.")
	}
}
