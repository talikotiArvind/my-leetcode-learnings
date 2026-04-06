package main

import "fmt"

func containsDuplicates(nums []int) bool {
	mapper := make(map[int]bool)

	for _, num := range nums {
		if mapper[num] {
			return true
		}
		mapper[num] = true
	}

	return false

}

func main() {
	data := []int{1, 2, 3, 5}
	output := containsDuplicates(data)
	fmt.Println(output)
}
