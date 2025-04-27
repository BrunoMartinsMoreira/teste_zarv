package main

import (
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		name        string
		matrix      [][]int
		want        int
		wantErr     bool
		description string
	}{
		{
			name: "Matriz 4x4 válida",
			matrix: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 1},
				{0, 0, 0, 0},
			},
			want:        3,
			wantErr:     false,
			description: "Deve calcular corretamente a distância na matriz 4x4",
		},
		{
			name: "Matriz 3x3 válida",
			matrix: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			want:        4,
			wantErr:     false,
			description: "Deve calcular corretamente a distância na matriz 3x3",
		},
		{
			name: "Matriz inválida - apenas um ponto",
			matrix: [][]int{
				{1, 0, 0},
				{0, 0, 0},
			},
			want:        0,
			wantErr:     true,
			description: "Deve retornar erro quando há apenas um ponto",
		},
		{
			name:        "Matriz vazia",
			matrix:      [][]int{},
			want:        0,
			wantErr:     true,
			description: "Deve retornar erro para matriz vazia",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := manhattanDistance(tt.matrix)

			if (err != nil) != tt.wantErr {
				t.Errorf("[%s] manhattanDistance() error = %v, wantErr %v\nDescrição: %s",
					tt.name, err, tt.wantErr, tt.description)
				return
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("[%s] manhattanDistance() = %d, want %d\nDescrição: %s",
					tt.name, got, tt.want, tt.description)
			}
		})
	}
}

func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		t.Errorf("run() retornou erro inesperado: %v", err)
	}
}
