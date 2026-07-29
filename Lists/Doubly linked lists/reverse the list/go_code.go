package main

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func Reverse(head *Node) *Node {
	current := head
	var newHead *Node

	for current != nil {
		current.Prev, current.Next = current.Next, current.Prev
		newHead = current
		current = current.Prev
	}

	return newHead
}

func PrintList(head *Node) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" <-> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {

	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}

	n1.Next = n2
	n2.Prev = n1

	n2.Next = n3
	n3.Prev = n2

	n3.Next = n4
	n4.Prev = n3

	fmt.Println("Original List:")
	PrintList(n1)

	head := Reverse(n1)

	fmt.Println("Reversed List:")
	PrintList(head)
}