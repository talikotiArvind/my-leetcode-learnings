package main

import "fmt"

func cycleSort(arr []int) {
	n := len(arr)

	for cycleStart := 0; cycleStart < n-1; cycleStart++ {
		item := arr[cycleStart]

		// Find correct position
		pos := cycleStart

		for i := cycleStart + 1; i < n; i++ {
			if arr[i] < item {
				pos++
			}
		}

		if pos == cycleStart {
			continue
		}

		// Skip duplicates
		for item == arr[pos] {
			pos++
		}

		arr[pos], item = item, arr[pos]

		// Rotate rest of the cycle
		for pos != cycleStart {
			pos = cycleStart

			for i := cycleStart + 1; i < n; i++ {
				if arr[i] < item {
					pos++
				}
			}

			for item == arr[pos] {
				pos++
			}

			arr[pos], item = item, arr[pos]
		}
	}
}

func main() {
	arr := []int{5, 3, 2, 4, 1}

	cycleSort(arr)

	fmt.Println(arr)
}