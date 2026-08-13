package main

import "fmt"

func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	firstRowZero, firstColZero := false, false

	for c := 0; c < cols; c++ {
		if matrix[0][c] == 0 {
			firstRowZero = true
		}
	}
	for r := 0; r < rows; r++ {
		if matrix[r][0] == 0 {
			firstColZero = true
		}
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if matrix[r][c] == 0 {
				matrix[r][0] = 0
				matrix[0][c] = 0
			}
		}
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if firstRowZero {
		for c := 0; c < cols; c++ {
			matrix[0][c] = 0
		}
	}
	if firstColZero {
		for r := 0; r < rows; r++ {
			matrix[r][0] = 0
		}
	}
}

func main() {
	m1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(m1)
	fmt.Println(m1) // [[1 0 1] [0 0 0] [1 0 1]]

	m2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(m2)
	fmt.Println(m2) // [[0 0 0 0] [0 4 5 0] [0 3 1 0]]
}