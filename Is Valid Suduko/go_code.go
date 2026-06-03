package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	// 9 rows, 9 cols, 9 boxes — each tracks seen digits
	var rows, cols, boxes [9][9]bool

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			digit := board[r][c] - '1' // '1'-'9' → 0-8
			boxIdx := (r/3)*3 + (c / 3)

			if rows[r][digit] || cols[c][digit] || boxes[boxIdx][digit] {
				return false
			}

			rows[r][digit] = true
			cols[c][digit] = true
			boxes[boxIdx][digit] = true
		}
	}

	return true
}

func main() {
	valid := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	// duplicate '8' in top-left 3x3 box
	invalid := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println("Valid board:  ", isValidSudoku(valid))   // true
	fmt.Println("Invalid board:", isValidSudoku(invalid)) // false
}
