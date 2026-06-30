package main

import "fmt"

// O(√n) — sum proper divisors by pairing i with num/i
func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	total := 1 // 1 is always a proper divisor for num > 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			total += i
			if i != num/i { // don't double-count the square root
				total += num / i
			}
		}
	}
	return total == num
}

func main() {
	tests := []int{6, 28, 496, 8128, 33550336, 1, 7, 100}
	for _, n := range tests {
		fmt.Printf("checkPerfectNumber(%8d) = %v\n", n, checkPerfectNumber(n))
	}
}
