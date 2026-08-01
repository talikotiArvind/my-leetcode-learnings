package main

import "fmt"

func pigeonholeSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	minVal := arr[0]
	maxVal := arr[0]

	// Find minimum and maximum
	for _, num := range arr {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	// Create pigeonholes
	size := maxVal - minVal + 1
	holes := make([]int, size)

	// Count each number
	for _, num := range arr {
		holes[num-minVal]++
	}

	// Put numbers back in sorted order
	index := 0

	for i := 0; i < size; i++ {
		for holes[i] > 0 {
			arr[index] = i + minVal
			index++
			holes[i]--
		}
	}
}

func main() {
	arr := []int{8, 3, 2, 7, 4, 6, 8}

	fmt.Println("Before:", arr)

	pigeonholeSort(arr)

	fmt.Println("After: ", arr)
}