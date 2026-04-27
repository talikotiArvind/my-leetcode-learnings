package main
import "fmt"
import "math"

func validSquares(nums int) bool {
    root := int(math.Sqrt(float64(nums)))
    return root*root == nums
}

func main() {
    fmt.Println(validSquares(16))
    fmt.Println(validSquares(15))
}
