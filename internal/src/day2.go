package src

import (
	"math"
	"strconv"
	"strings"
)

func InvalidIdsSum(input string) int64 {
	ranges := strings.Split(input, ",")

	point := make(chan int64, len(ranges))
	for _, rng := range ranges {
		go worker(point, rng)
	}

	var total int64 = 0
	for range ranges {
		total += <-point
	}
	return total
}

func worker(point chan int64, rng string) {
	var total int64 = 0
	r := strings.Split(rng, "-")
	start, _ := strconv.ParseInt(strings.TrimSpace(r[0]), 10, 64)
	end, _ := strconv.ParseInt(strings.TrimSpace(r[1]), 10, 64)

	for num := start; num <= end; num++ {
		length, isValid := getLenEven(num)
		if !isValid {
			continue
		}
		divisor := int64(math.Pow10(length / 2))
		first := num / divisor
		second := num % divisor
		if first == second {
			total += num
		}
	}
	point <- total
}

func getLenEven(n int64) (int, bool) {
	length := 0
	temp := n
	for temp > 0 {
		temp /= 10
		length++
	}

	return length, length%2 == 0
}


// ______________________________________________

func InvalidIdsSum2(input string) int64 {
	ranges := strings.Split(input, ",")

	point := make(chan int64, len(ranges))
	for _, rng := range ranges {
		go worker2(point, rng)
	}

	var total int64 = 0
	for range ranges {
		total += <-point
	}
	return total
}

func worker2(point chan int64, rng string) {
	var total int64 = 0
	r := strings.Split(rng, "-")
	start, _ := strconv.ParseInt(strings.TrimSpace(r[0]), 10, 64)
	end, _ := strconv.ParseInt(strings.TrimSpace(r[1]), 10, 64)

	for num := start; num <= end; num++ {
		length := getLen(num)
		for size := 1; size <= length / 2; size++ {
			if length % size!= 0 {
                continue
            }
			if checkPattern(num, size, length/size) {
                total += num
                break
            }
		}
		
	}

	point <- total
}

func getLen(n int64) int {
	length := 0
	temp := n
	for temp > 0 {
		temp /= 10
		length++
	}

	return length
}

func checkPattern(num int64, size int, numBlocks int) bool {
    divisor := int64(math.Pow10(size))
    firstPart := num % divisor
    temp := num / divisor
    
    for i := 1; i < numBlocks; i++ {
        if temp % divisor != firstPart {
            return false
        }
        temp /= divisor
    }
    
    return true
}