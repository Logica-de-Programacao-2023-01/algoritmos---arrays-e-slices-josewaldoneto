package main

import "fmt"

func main() {
	matriz := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var linha, coluna int
	fmt.Print("Digite o índice da linha (0 ou 1): ")
	fmt.Scanln(&linha)
	fmt.Print("Digite o índice da coluna (0, 1 ou 2): ")
	fmt.Scanln(&coluna)

	valor := matriz[linha][coluna]
	fmt.Printf("Valor na posição [%d][%d]: %d\n", linha, coluna, valor)
}
