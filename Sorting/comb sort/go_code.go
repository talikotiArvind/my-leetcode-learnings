package main

import "fmt"

func combSort(arr []int) {
	n := len(arr)

	gap := n
	shrink := 1.3
	sorted := false

	for !sorted {
		gap = int(float64(gap) / shrink)

		if gap <= 1 {
			gap = 1
			sorted = true
		}

		for i := 0; i+gap < n; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				sorted = false
			}
		}
	}
}

func main() {
	arr := []int{8, 4, 1, 6, 3, 7, 2}

	combSort(arr)

	fmt.Println(arr)
}