package main
import "fmt"

func islandPerimeter(grid [][]int) int {
    rows, columns := len(grid), len(grid[0])
    result := 0
    
    for r:= 0; r < rows; r++ {
        for c:=0; c < columns; c++ {
            if grid[r][c] == 1 {
                result = result + 4
                if r > 0 && grid[r-1][c] == 1{
                    result = result - 2
                }
                if c > 0 && grid[r][c-1] == 1{
                    result = result - 2
                }
            }
        }
    }
    return result
}

func main() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}

	result := islandPerimeter(grid)
	fmt.Println("Perimeter:", result) // 16
}
