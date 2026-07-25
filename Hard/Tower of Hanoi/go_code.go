package main

import "fmt"

func hanoi(n int, source, target, auxiliary string) {
	if n == 1 {
		fmt.Printf("Move disk 1 from %s to %s\n", source, target)
		return
	}
	hanoi(n-1, source, auxiliary, target)
	fmt.Printf("Move disk %d from %s to %s\n", n, source, target)
	hanoi(n-1, auxiliary, target, source)
}

func main() {
	n := 3
	fmt.Printf("Tower of Hanoi with %d disks:\n", n)
	hanoi(n, "A", "C", "B")
}
