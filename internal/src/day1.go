package src

import (
	"strconv"
)

func absMode(n, d int) int {
	result := n % d
	if result < 0 { result += d }
    return result
}

func GetZeros(codes []string) (int, error) {

	current_pos := 50
	zeros := 0

	for _, line := range codes {
		direction := line[0:1]
		shiftStr := line[1:]
		shiftNum, err := strconv.Atoi(shiftStr)
		if err != nil {return -1, err}
		
		switch direction {
		case "L":
			current_pos -= shiftNum
		case "R":
			current_pos += shiftNum
		}
		if absMode(current_pos, 100) == 0 {
			zeros ++
		}
	}
	return zeros, nil
}
