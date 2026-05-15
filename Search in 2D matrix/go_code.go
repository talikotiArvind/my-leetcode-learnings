// Online Go compiler to run Golang program online
// Print "Start small. Ship something." message

package main
import "fmt"

func searchMatrix (nums [][]int, target int) bool {
    m, n := len(nums), len(nums[0])
    low, high := 0, m*n-1
    for low <= high {
        mid := (low + high) / 2
        value := nums[mid / n][mid % n]
        if value == target {
            return true
        } else if value < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(searchMatrix(matrix, 3))  // true
	fmt.Println(searchMatrix(matrix, 13)) // false
}
