package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

func main() {
	check := func(got, want int) {
		if got != want {
			panic(fmt.Sprintf("got %d, want %d", got, want))
		}
	}
	check(getSum(1, 2), 3)
	check(getSum(-2, 3), 1)
	check(getSum(-12, 8), -4)
	check(getSum(0, 0), 0)
	check(getSum(-5, -7), -12)
	fmt.Println("All tests passed.")
}
