package main

import "fmt"

func uniquePaths(m int, n int) int {
	row := make([]int, n)
	for j := range row {
		row[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			row[j] += row[j-1]
		}
	}
	return row[n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7)) // 28
	fmt.Println(uniquePaths(3, 2)) // 3
	fmt.Println(uniquePaths(1, 1)) // 1
}