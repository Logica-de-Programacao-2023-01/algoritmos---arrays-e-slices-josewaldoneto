package main

import "fmt"

// Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.

func main() {
	var arr = [3]int{80, 127, 97}
	var soma int
	for _, x := range arr {
		soma += x
	}
	fmt.Println(soma)
}
