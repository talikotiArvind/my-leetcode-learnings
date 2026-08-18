package main

import "fmt"

func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(hammingWeight(11))         // 3
	fmt.Println(hammingWeight(128))        // 1
	fmt.Println(hammingWeight(2147483645)) // 30
}
