package main

import "fmt"

func maxCoins(nums []int) int {

	balloons := make([]int, len(nums)+2)

	balloons[0] = 1
	balloons[len(balloons)-1] = 1

	copy(balloons[1:], nums)

	n := len(balloons)

	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Interval length
	for length := 2; length < n; length++ {

		for left := 0; left+length < n; left++ {

			right := left + length

			for k := left + 1; k < right; k++ {

				coins := dp[left][k] +
					balloons[left]*balloons[k]*balloons[right] +
					dp[k][right]

				if coins > dp[left][right] {
					dp[left][right] = coins
				}
			}
		}
	}

	return dp[0][n-1]
}

func main() {

	nums := []int{3, 1, 5, 8}

	result := maxCoins(nums)

	fmt.Println(result)
}