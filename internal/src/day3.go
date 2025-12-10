package src

import (
	"strings"
)

func MinVoltage(input string) int {
    lines := strings.Split(strings.TrimSpace(input), "\n")
    total := 0
    
    for _, line := range lines {
        currentMax := int(line[len(line)-1] - '0')
        lineBest := 0
        
        for i := len(line) - 2; i >= 0; i-- {
            digit := int(line[i] - '0')
            num := digit*10 + currentMax
            
            if num > lineBest {
                lineBest = num
            }
            
            if digit > currentMax {
                currentMax = digit
            }
        }
        
        total += lineBest
    }
    
    return total
}


func MaxVoltage(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
    total := 0
    
    for _, line := range lines {
		pos := 0
		number := 0

		for m := range 12 {
			endPos := len(line) - (12 - m - 1)

			maxDigit, maxDigitPos := findMaxDigitPos(line, pos, endPos)

			number = number*10 + int(maxDigit-'0')
			pos = maxDigitPos + 1
		}
		total += number
	}

	return total
}

func findMaxDigitPos(line string, pos int, endPos int) (byte, int) {
	maxDigit := byte('0')
	maxDigitPos := pos
	for i := pos; i < endPos; i++ {
		if line[i] > maxDigit {
			maxDigit = line[i]
			maxDigitPos = i
		}
	}
	return maxDigit, maxDigitPos
}