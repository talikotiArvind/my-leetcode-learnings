package main

import "fmt"

func duplicateZeros(arr []int) {
	n := len(arr)
	zeros := 0
	for _, v := range arr {
		if v == 0 {
			zeros++
		}
	}

	i, j := n-1, n+zeros-1

	for i >= 0 {
		if arr[i] == 0 {
			if j < n {
				arr[j] = 0
			}
			j--
			if j < n {
				arr[j] = 0
			}
		} else {
			if j < n {
				arr[j] = arr[i]
			}
		}
		i--
		j--
	}
}

func main() {
	tests := [][]int{
		{1, 0, 2, 3, 0, 4, 5, 0},
		{1, 2, 3},
		{0, 0, 0, 0},
		{1, 0, 3},
	}

	for _, arr := range tests {
		original := make([]int, len(arr))
		copy(original, arr)
		duplicateZeros(arr)
		fmt.Printf("Input:  %v\nOutput: %v\n\n", original, arr)
	}
}
