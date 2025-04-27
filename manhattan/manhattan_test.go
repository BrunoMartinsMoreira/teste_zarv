package manhattan

import (
	"reflect"
	"testing"
)

func TestFindCoordinates(t *testing.T) {
	tests := []struct {
		name    string
		matrix  [][]int
		want    []Coordinates
		wantErr bool
	}{
		{
			name: "Matriz válida com dois pontos",
			matrix: [][]int{
				{0, 1, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			want: []Coordinates{
				{line: 0, collumn: 1},
				{line: 2, collumn: 2},
			},
			wantErr: false,
		},
		{
			name: "Matriz com apenas um ponto",
			matrix: [][]int{
				{0, 1, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Matriz com três pontos",
			matrix: [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: []Coordinates{
				{line: 0, collumn: 1},
				{line: 1, collumn: 1},
			},
			wantErr: false,
		},
		{
			name: "Matriz vazia",
			matrix: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findCoordinates(tt.matrix)

			if (err != nil) != tt.wantErr {
				t.Errorf("findCoordinates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateDistance(t *testing.T) {
	tests := []struct {
		name     string
		coord1   Coordinates
		coord2   Coordinates
		expected int
	}{
		{
			name:     "Distância na horizontal",
			coord1:   Coordinates{line: 0, collumn: 0},
			coord2:   Coordinates{line: 0, collumn: 3},
			expected: 3,
		},
		{
			name:     "Distância na vertical",
			coord1:   Coordinates{line: 0, collumn: 0},
			coord2:   Coordinates{line: 4, collumn: 0},
			expected: 4,
		},
		{
			name:     "Distância diagonal",
			coord1:   Coordinates{line: 0, collumn: 0},
			coord2:   Coordinates{line: 2, collumn: 2},
			expected: 4,
		},
		{
			name:     "Distância com coordenadas negativas",
			coord1:   Coordinates{line: -1, collumn: -1},
			coord2:   Coordinates{line: 1, collumn: 1},
			expected: 4,
		},
		{
			name:     "Mesmas coordenadas",
			coord1:   Coordinates{line: 1, collumn: 1},
			coord2:   Coordinates{line: 1, collumn: 1},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateDistance(tt.coord1, tt.coord2)
			if result != tt.expected {
				t.Errorf("calculateDistance() = %v, want %v", result, tt.expected)
			}
		})
	}
}
