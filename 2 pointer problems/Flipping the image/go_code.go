package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		lo, hi := 0, len(row)-1
		for lo < hi {
			row[lo], row[hi] = row[hi]^1, row[lo]^1
			lo++
			hi--
		}
		if lo == hi {
			row[lo] ^= 1 // odd-length middle
		}
	}
	return image
}

func main() {
	img := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(flipAndInvertImage(img))
	// [[1 0 0] [0 1 0] [1 1 1]]
}
