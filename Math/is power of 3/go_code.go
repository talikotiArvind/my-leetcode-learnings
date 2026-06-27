package main

import "fmt"

func isPowerOfThree(n int) bool {
    if n < 1 {
        return false
    }
    for n%3 == 0 {
        n /= 3
    }
    return n == 1
}

func main() {
    for _, n := range []int{27, 0, 9, 45, 1, 243} {
        fmt.Printf("n=%d -> %v\n", n, isPowerOfThree(n))
    }
    // n=27  -> true
    // n=0   -> false
    // n=9   -> true
    // n=45  -> false  (45 = 9*5)
    // n=1   -> true   (3^0)
    // n=243 -> true
}
