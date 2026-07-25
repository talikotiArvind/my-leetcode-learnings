package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	a := 0
	b := int(math.Sqrt(float64(c)))
	for a <= b {
		s := a*a + b*b
		if s == c {
			return true
		} else if s < c {
			a++
		} else {
			b--
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(3))
	fmt.Println(judgeSquareSum(4))
	fmt.Println(judgeSquareSum(2))
}
