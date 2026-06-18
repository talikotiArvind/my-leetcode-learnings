package main

import "fmt"

func countNegatives(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	count := 0
	row := 0
	col := cols - 1

	for row < rows && col >= 0 {
		if grid[row][col] < 0 {
			count += (rows - row)
			col--
		} else {
			row++
		}
	}

	return count
}
