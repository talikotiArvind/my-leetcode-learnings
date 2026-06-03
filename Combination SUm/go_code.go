package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	var backtrack func(start, remaining int, current []int)
	backtrack = func(start, remaining int, current []int) {
		if remaining == 0 {
			combo := make([]int, len(current))
			copy(combo, current)
			result = append(result, combo)
			return
		}
		if remaining < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			current = append(current, candidates[i])
			backtrack(i, remaining-candidates[i], current)
			current = current[:len(current)-1]
		}
	}

	backtrack(0, target, []int{})
	return result
}

func main() {
	tests := []struct {
		candidates []int
		target     int
	}{
		{[]int{2, 3, 6, 7}, 7},
		{[]int{2, 3, 5}, 8},
		{[]int{2}, 1},
	}

	for _, tc := range tests {
		res := combinationSum(tc.candidates, tc.target)
		fmt.Printf("candidates=%v  target=%d\n", tc.candidates, tc.target)
		fmt.Printf("  → %v\n\n", res)
	}
}
