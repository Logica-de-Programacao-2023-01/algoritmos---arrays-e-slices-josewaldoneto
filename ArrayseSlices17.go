package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	soma := 0
	for i := 0; i < len(arr); i += 2 {
		soma += arr[i]
	}
	fmt.Printf("Soma dos elementos nas posições pares: %d\n", soma)
}
