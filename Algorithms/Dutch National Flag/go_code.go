package main

import "fmt"

// dutchFlagSort sorts an array of 0s, 1s and 2s in a single pass, in place.
func dutchFlagSort(arr []int) {
	low, mid, high := 0, 0, len(arr)-1
	for mid <= high {
		switch arr[mid] {
		case 0:
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		case 1:
			mid++
		default: // 2
			arr[mid], arr[high] = arr[high], arr[mid]
			high-- // note: mid is NOT advanced here
		}
	}
}

func main() {
	flag := []int{2, 0, 2, 1, 1, 0, 0, 2, 1}
	fmt.Printf("before : %v\n", flag)
	dutchFlagSort(flag)
	fmt.Printf("after  : %v\n", flag)
}
