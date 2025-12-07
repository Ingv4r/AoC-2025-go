package tests

import (
	"fmt"
	"testing"

	"AoC-2025-go/internal/src"
)

func TestGetZeros(t *testing.T) {
	input := "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"
	
	expectedPart1 := 3
	
	result1, err := src.GetZeros(input)
	if err != nil {
		t.Errorf("GetZeros() неожиданная ошибка: %v", err)
	}
	
	if result1 != expectedPart1 {
		t.Errorf("GetZeros() = %d, want %d", result1, expectedPart1)
	} else {
		fmt.Printf("✓ GetZeros правильно: %d\n", result1)
	}
}

func TestGetZeros2Part(t *testing.T) {
	input := "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"
	
	expectedPart2 := 6
	
	result2, err := src.GetZeros2Part(input)
	if err != nil {
		t.Errorf("GetZeros2Part() неожиданная ошибка: %v", err)
	}
	
	if result2 != expectedPart2 {
		t.Errorf("GetZeros2Part() = %d, want %d", result2, expectedPart2)
	} else {
		fmt.Printf("✓ GetZeros2Part правильно: %d\n", result2)
	}
}
