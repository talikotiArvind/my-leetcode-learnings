package main
import "fmt"

func climbStairs (n int) int {
    i, j := 1, 1
    for x:= 0; x < n - 1; x++ {
        i, j = j, i+j
    }
    return j
}

func main() {
    fmt.Println(climbStairs(5))
}
