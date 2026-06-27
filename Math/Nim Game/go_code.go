package main

import "fmt"

func canWinNim(n int) bool {
    return n%4 != 0
}

func main() {
    for _, n := range []int{1, 2, 3, 4, 5, 8} {
        fmt.Printf("n=%d -> %v\n", n, canWinNim(n))
    }
    // Output:
    // n=1 -> true
    // n=2 -> true
    // n=3 -> true
    // n=4 -> false
    // n=5 -> true
    // n=8 -> false
}
