package tests

import (
	"fmt"
	"testing"

	"AoC-2025-go/internal/src"
)

func TestInvalidIdsSum(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`
	
	expectedPart1 := int64(1227775554)
	
	result1 := src.InvalidIdsSum(input)
	if result1 != expectedPart1 {
		t.Errorf("InvalidIdsSum() = %d, want %d", result1, expectedPart1)
	} else {
		fmt.Printf("✓ InvalidIdsSum правильно: %d\n", result1)
	}
}

func TestInvalidIdsSum2(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`
	
	expectedPart2 := int64(4174379265)
	
	result2 := src.InvalidIdsSum2(input)
	if result2 != expectedPart2 {
		t.Errorf("InvalidIdsSum2() = %d, want %d", result2, expectedPart2)
	} else {
		fmt.Printf("✓ InvalidIdsSum2 правильно: %d\n", result2)
	}
}