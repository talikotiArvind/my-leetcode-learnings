package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := 0, len(letters)

	for lo < hi {
		mid := (lo + hi) / 2
		if letters[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return letters[lo%len(letters)]
}

func main() {
	fmt.Printf("%c\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd')) // f
	fmt.Printf("%c\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j')) // c
}
