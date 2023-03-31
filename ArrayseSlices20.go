package main

import "fmt"

// Crie um programa que leia um slice de inteiros e verifique se ele está ordenado em ordem crescente.

func main() {
	var tamanho, x int
	var slice []int

	fmt.Print("Qual tamanho do slice?")
	fmt.Scan(&tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)
		slice = append(slice, x)
	}

	crescente := true
	for i := 1; i < len(slice); i++ {
		atual := slice[i]
		anterior := slice[i-1]
		if atual < anterior {
			fmt.Println("Não está em ordem crescente.")
			fmt.Println(slice)
			crescente = false
			break
		}
		anterior = atual
	}
	if crescente {
		fmt.Println("Está em ordem crescente.")
		fmt.Println(slice)
	}
}
