package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	// 1. Create a copy of the slice and sort it in descending order
	ranked := make([]int, len(score))
	copy(ranked, score)
	sort.Slice(ranked, func(i, j int) bool {
		return ranked[i] > ranked[j]
	})

	// 2. Build the rank lookup map
	rankMap := make(map[int]string)
	for i, s := range ranked {
		switch i {
		case 0:
			rankMap[s] = "Gold Medal"
		case 1:
			rankMap[s] = "Silver Medal"
		case 2:
			rankMap[s] = "Bronze Medal"
		default:
			rankMap[s] = strconv.Itoa(i + 1)
		}
	}

	// 3. Build the final output slice using the original score order
	result := make([]string, len(score))
	for i, s := range score {
		result[i] = rankMap[s]
	}

	return result
}

func main() {
	score := []int{10, 3, 8, 9, 4}
	fmt.Println(findRelativeRanks(score))
	// Output: [Gold Medal Bronze Medal 4 Silver Medal 5]
}
