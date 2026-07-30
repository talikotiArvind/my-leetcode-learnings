package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area))) // largest possible width
	for area%w != 0 {                  // walk down to nearest divisor
		w--
	}
	return []int{area / w, w} // [L, W], guaranteed L >= W
}

func main() {
	examples := []int{4, 37, 122122}
	for _, area := range examples {
		res := constructRectangle(area)
		fmt.Printf("area = %d -> L = %d, W = %d\n", area, res[0], res[1])
	}
}
