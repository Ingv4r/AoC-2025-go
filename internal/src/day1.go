package src

import (
	"strconv"
	"strings"
)

func GetZeros(input string) (int, error) {
	code := strings.Split(strings.TrimSpace(input), "\n")
	currentPos := 50
	zeros := 0

	for _, line := range code {
		direction := line[0:1]
		shiftStr := line[1:]
		shiftNum, err := strconv.Atoi(shiftStr)
		if err != nil {
			return -1, err
		}

		switch direction {
		case "L":
			currentPos -= shiftNum
		case "R":
			currentPos += shiftNum
		}
		if absMode(currentPos, 100) == 0 {
			zeros++
		}
	}
	return zeros, nil
}

func GetZeros2Part(input string) (int, error) {
	code := strings.Split(strings.TrimSpace(input), "\n")
	currentPos := 50
	zeros := 0

	for _, line := range code {
		direction := line[0:1]
		shiftStr := line[1:]
		shiftNum, err := strconv.Atoi(shiftStr)
		if err != nil {
			return -1, err
		}

		oldPos := currentPos
		switch direction {
		case "R":
			currentPos += shiftNum
			zeros += floorDiv(currentPos, 100) - floorDiv(oldPos, 100)
		case "L":
			currentPos -= shiftNum
			zeros += floorDiv(oldPos-1, 100) - floorDiv(currentPos-1, 100)
		}
	}
	return zeros, nil
}

func absMode(n, d int) int {
	result := n % d
	if result < 0 {
		result += d
	}
	return result
}

func floorDiv(a, b int) int {
	q := a / b
	if a < 0 && a%b != 0 {
		q--
	}
	return q
}
