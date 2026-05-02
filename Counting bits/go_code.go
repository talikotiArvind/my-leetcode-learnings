package main

import "fmt"

func countBits(n int) []int {
    ans := make([]int, n+1)
    for i := 1; i <= n; i++ {
        ans[i] = ans[i>>1] + i&1
    }
    return ans
}

func main() {
    fmt.Println(countBits(2))  // [0 1 1]
    fmt.Println(countBits(5))  // [0 1 1 2 1 2]
}
