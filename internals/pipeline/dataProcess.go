package pipeline

import (
	"math"
)

func GenerateNumber(n int, num chan int) {
	for i := range n {
		num <- i+1
	}
	close(num)
}

func EvenNumberFilter(num, evenNum chan int) {
	for val := range num {
		if val % 2 == 0 {
			evenNum <- val
		}
	}
	close(evenNum)
}

func SquaringValue(evenNum, squaredNum chan int) {
	for val := range evenNum {
		squaredNum <- int(math.Pow(float64(val), 2))
	}
	close(squaredNum)
}