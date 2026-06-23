package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	total := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		diff := timeSeries[i+1] - timeSeries[i]
		total += min(diff, duration)
	}
	total += duration
	return total
}

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))    // 4
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))    // 3
	fmt.Println(findPoisonedDuration([]int{1, 4, 9}, 3)) // 9
	fmt.Println(findPoisonedDuration([]int{}, 2))        // 0
}
