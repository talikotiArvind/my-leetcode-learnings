package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return result
	}

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++

		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--

		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}

		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}

	return result
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// [1 2 3 6 9 8 7 4 5]
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	// [1 2 3 4 8 12 11 10 9 5 6 7]
}