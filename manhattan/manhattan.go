package manhattan

import (
	"fmt"
	"math"
)

type Coordinates struct {
	line    int
	collumn int
}

func findCoordinates(matrix [][]int) ([]Coordinates, error) {
	coordinates := make([]Coordinates, 0, 2)

searchLoop:
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 1 {
				coordinates = append(coordinates, Coordinates{line: i, collumn: j})
				if len(coordinates) == 2 {
					break searchLoop
				}
			}
		}
	}

	if len(coordinates) != 2 {
		return nil, fmt.Errorf("a matriz deve conter exatamente dois pontos com valor 1")
	}

	return coordinates, nil
}

func calculateDistance(coord1, coord2 Coordinates) int {
	return int(math.Abs(float64(coord1.line-coord2.line))) +
		int(math.Abs(float64(coord1.collumn-coord2.collumn)))
}

func Distance(matrix [][]int) (int, error) {
	if len(matrix) == 0 || len(matrix) > 100 {
		return 0, fmt.Errorf("a matriz deve ter entre 1 e 100 linhas")
	}

	if len(matrix[0]) == 0 || len(matrix[0]) > 100 {
		return 0, fmt.Errorf("a matriz deve ter entre 1 e 100 colunas")
	}

	coordinates, err := findCoordinates(matrix)
	if err != nil {
		return 0, err
	}

	return calculateDistance(coordinates[0], coordinates[1]), nil
}
