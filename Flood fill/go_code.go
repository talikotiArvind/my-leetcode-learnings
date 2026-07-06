package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows, cols := len(image), len(image[0])
	start := image[sr][sc]
	if start == color {
		return image
	}

	queue := [][2]int{{sr, sc}}
	image[sr][sc] = color
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		r, c := cell[0], cell[1]
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && image[nr][nc] == start {
				image[nr][nc] = color
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}

	return image
}

func main() {
	fmt.Println(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	// [[2 2 2] [2 2 0] [2 0 1]]
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0))
	// [[0 0 0] [0 0 0]]
}
