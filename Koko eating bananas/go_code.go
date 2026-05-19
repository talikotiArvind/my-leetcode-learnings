package main

import (
    "fmt"
    "math"
)

func minEatingSpeed(piles []int, h int) int {
    lo, hi := 1, 0
    for _, p := range piles {
        if p > hi {
            hi = p
        }
    }

    for lo < hi {
        mid := (lo + hi) / 2
        hours := 0
        for _, p := range piles {
            hours += int(math.Ceil(float64(p) / float64(mid)))
        }

        if hours <= h {
            hi = mid     // mid works, try slower
        } else {
            lo = mid + 1 // too slow, go faster
        }
    }
    return lo
}

func main() {
    fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))        // Output: 4
    fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))  // Output: 30
    fmt.Println(minEatingSpeed([]int{1, 1, 1, 999999999}, 10)) // Output: 142857143
}