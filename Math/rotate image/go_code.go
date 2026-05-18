package main
import "fmt"

func rotate(matrix [][]int) {
    n := len(matrix)
    
    for i:=0; i< n; i++ {
        for j := i; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i] ,matrix[i][j]
        }
    }
    
    for i :=0; i< n; i++ {
        left, right := 0, n-1
        for left < right {
            matrix[i][left], matrix[i][right] = matrix[i][right],matrix[i][left]
            left++
            right--
        }
    }
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Before:")
	printMatrix(matrix)

	rotate(matrix)

	fmt.Println("\nAfter 90° clockwise rotation:")
	printMatrix(matrix)
}
