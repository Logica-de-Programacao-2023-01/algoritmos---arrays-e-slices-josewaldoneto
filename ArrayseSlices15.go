package main

import "fmt"

func main() {
	arr := [10]float64{1.2, 3.4, 5.6, 7.8, 9.0, 1.1, 2.2, 3.3, 4.4, 5.5}
	var slice []float64
	for _, v := range arr {
		if v > 5 {
			slice = append(slice, v)
		}
	}
	fmt.Println("Novo Slice:", slice)
}
