package main

import "fmt"

// Approach 1: Bit manipulation with mask — O(1), no loops/recursion
// 0x55555555 = 0101...0101, catches a single set bit at even positions only
func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n&0x55555555 != 0
}

// Approach 2: Power-of-two check + modulo trick (4 ≡ 1 mod 3) — O(1)
func isPowerOfFourMod(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}

// Approach 3: Iterative — easy to read
func isPowerOfFourIter(n int) bool {
	if n < 1 {
		return false
	}
	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

func main() {
	tests := []int{1, 4, 5, 8, 16, 0, -4, 64, 32}
	for _, n := range tests {
		fmt.Printf("n = %3d | mask: %-5v | mod: %-5v | iter: %-5v\n",
			n, isPowerOfFour(n), isPowerOfFourMod(n), isPowerOfFourIter(n))
	}
}
