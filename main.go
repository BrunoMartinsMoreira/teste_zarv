package main

import (
	"fmt"
	"teste_zarv/manhattan"
)

func manhattanDistance(matrix [][]int) (int, error) {
	return manhattan.Distance(matrix)
}

func run() error {
	matrix := [][]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
	}

	distance, err := manhattanDistance(matrix)
	if err != nil {
		return err
	}

	fmt.Printf("A distância de Manhattan entre os dois pontos é: %d\n", distance)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("Erro: %v\n", err)
	}
}
