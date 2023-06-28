package main

import "fmt"

func main() {
	var n int
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Print("Digite um número: ")
	fmt.Scanln(&n)
	arr[0] += n
	arr[len(arr)-1] += n
	fmt.Println("Array resultante:", arr)
}
