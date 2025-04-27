package manhattan

import "fmt"

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
