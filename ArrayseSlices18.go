package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scanln(&n)

	primeirosPrimos := gerarPrimos(n)

	fmt.Println("Array com os", n, "primeiros números primos:", primeirosPrimos)
}

func gerarPrimos(n int) []int {
	primos := make([]int, 0)
	num := 2

	for len(primos) < n {
		if ehPrimo(num) {
			primos = append(primos, num)
		}
		num++
	}

	return primos
}

func ehPrimo(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
