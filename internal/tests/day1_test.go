package tests

import (
	"testing"
	"AoC-2025-go/internal/src"
)

func TestGetZeros(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
		wantErr  bool
	}{
		{
			name: "Тест 1",
			input: []string{
				"L68",
				"L30",
				"R48",
				"L5",
				"R60",
				"L55",
				"L1",
				"L99",
				"R14",
				"L82",
			},
			expected: 3,
			wantErr:  false,
		},
		{
			name: "Тест 2",
			input: []string{
				"L25",
				"L25",
				"R25",
				"L25",
				"R25",
				"L25",
			},
			expected: 3,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := src.GetZeros(tt.input)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetZeros() ожидалась ошибка, но не получена")
				}
				return
			}

			if err != nil {
				t.Errorf("GetZeros() неожиданная ошибка: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("GetZeros() = %v, ожидалось %v", result, tt.expected)
			}
		})
	}
}