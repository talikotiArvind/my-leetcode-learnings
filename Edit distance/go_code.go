package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {

	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {

		for j := 1; j <= n; j++ {

			if word1[i-1] == word2[j-1] {

				dp[i][j] = dp[i-1][j-1]

			} else {

				deleteOp := dp[i-1][j]
				insertOp := dp[i][j-1]
				replaceOp := dp[i-1][j-1]

				dp[i][j] = 1 + min(
					deleteOp,
					min(insertOp, replaceOp),
				)
			}
		}
	}

	return dp[m][n]
}

func main() {

	word1 := "horse"
	word2 := "ros"

	fmt.Println(minDistance(word1, word2))
}