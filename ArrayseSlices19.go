package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite o tamanho dos arrays: ")
	fmt.Scanln(&n)
	arr1 := make([]int, n)
	arr2 := make([]int, n)
	fmt.Println("Digite os elementos do primeiro array:")
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr1[i])
	}
	fmt.Println("Digite os elementos do segundo array:")
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr2[i])
	}
	arr3 := make([]int, n)
	for i := 0; i < n; i++ {
		arr3[i] = arr1[i] + arr2[i]
	}
	fmt.Println("Terceiro array:", arr3)
}
