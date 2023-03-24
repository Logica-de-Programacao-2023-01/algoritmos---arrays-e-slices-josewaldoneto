package main

import "fmt"

// Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante.

func main() {
	var slic = []int{1, 2, 3, 4, 5}
	slic = append(slic[:2], slic[3:]...)
	fmt.Print(slic)
}
