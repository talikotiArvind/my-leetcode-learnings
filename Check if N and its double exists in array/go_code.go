package main

import "fmt"

func checkIfExist(arr []int) bool {
	for i, v := range arr {
		for j, w := range arr {
			if i != j && w == 2*v {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))  // true
	fmt.Println(checkIfExist([]int{3, 1, 7, 11}))  // false
}
