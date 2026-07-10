package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for _, coin := range coins {
		for a := coin; a <= amount; a++ {
			if dp[a-coin]+1 < dp[a] {
				dp[a] = dp[a-coin] + 1
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11)) // 3
	fmt.Println(coinChange([]int{2}, 3))        // -1
	fmt.Println(coinChange([]int{1}, 0))        // 0
}
