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

func TestDistance(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int
		wantErr  bool
	}{
		{
			name: "Matriz válida com distância horizontal",
			matrix: [][]int{
				{1, 0, 0, 1},
				{0, 0, 0, 0},
			},
			expected: 3,
			wantErr:  false,
		},
		{
			name: "Matriz válida com distância vertical",
			matrix: [][]int{
				{1, 0},
				{0, 0},
				{0, 1},
			},
			expected: 3,
			wantErr:  false,
		},
		{
			name: "Matriz válida com distância diagonal",
			matrix: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			expected: 4,
			wantErr:  false,
		},
		{
			name:     "Matriz vazia",
			matrix:   [][]int{},
			expected: 0,
			wantErr:  true,
		},
		{
			name: "Matriz muito grande (>100 linhas)",
			matrix: func() [][]int {
				m := make([][]int, 101)
				for i := range m {
					m[i] = make([]int, 1)
				}
				return m
			}(),
			expected: 0,
			wantErr:  true,
		},
		{
			name: "Matriz com apenas um ponto",
			matrix: [][]int{
				{1, 0, 0},
				{0, 0, 0},
			},
			expected: 0,
			wantErr:  true,
		},
		{
			name: "Matriz com mais de dois pontos",
			matrix: [][]int{
				{1, 0, 1},
				{0, 1, 0},
			},
			expected: 2,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Distance(tt.matrix)

			if (err != nil) != tt.wantErr {
				t.Errorf("Distance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.expected {
				t.Errorf("Distance() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDistanceIntegration(t *testing.T) {
	tests := []struct {
		name        string
		matrix      [][]int
		expected    int
		wantErr     bool
		description string
	}{
		{
			name: "Integração - Fluxo completo com sucesso",
			matrix: [][]int{
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 1},
			},
			expected:    4,
			wantErr:     false,
			description: "Testa o fluxo completo passando por findCoordinates e calculateDistance",
		},
		{
			name: "Integração - Matriz inválida por tamanho",
			matrix: func() [][]int {
				m := make([][]int, 101)
				for i := range m {
					m[i] = make([]int, 1)
					m[i][0] = 1
				}
				return m
			}(),
			expected:    0,
			wantErr:     true,
			description: "Deve falhar logo na validação inicial do tamanho da matriz",
		},
		{
			name: "Integração - Matriz inválida por quantidade de pontos",
			matrix: [][]int{
				{0, 1, 0},
				{0, 0, 0},
			},
			expected:    0,
			wantErr:     true,
			description: "Deve falhar na função findCoordinates por ter apenas um ponto",
		},
		{
			name: "Integração - Distância zero entre pontos",
			matrix: [][]int{
				{1, 1, 0},
				{0, 0, 0},
			},
			expected:    1,
			wantErr:     false,
			description: "Testa o fluxo completo com pontos próximos",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distance, err := Distance(tt.matrix)

			if (err != nil) != tt.wantErr {
				t.Errorf("[%s] Distance() error = %v, wantErr %v\nDescrição: %s",
					tt.name, err, tt.wantErr, tt.description)
				return
			}

			if !tt.wantErr && distance != tt.expected {
				t.Errorf("[%s] Distance() = %v, want %v\nDescrição: %s",
					tt.name, distance, tt.expected, tt.description)
			}
		})
	}
}
