package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, x := range nums {
		counts[x]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range counts {
		buckets[freq] = append(buckets[freq], num)
	}

	result := []int{}
	for freq := len(buckets) - 1; freq > 0 && len(result) < k; freq-- {
		for _, num := range buckets[freq] {
			result = append(result, num)
			if len(result) == k {
				break
			}
		}
	}
	return result
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // [1 2]
	fmt.Println(topKFrequent([]int{1}, 1))                // [1]
}