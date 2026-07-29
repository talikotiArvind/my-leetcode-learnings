package main

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func InsertAtEnd(head *Node, data int) *Node {
	newNode := &Node{Val: data}

	// Empty list
	if head == nil {
		return newNode
	}

	current := head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	newNode.Prev = current

	return head
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

	n1.Next = n2
	n2.Prev = n1

	n2.Next = n3
	n3.Prev = n2

	fmt.Println("Before Insertion:")
	PrintList(n1)

	head := InsertAtEnd(n1, 4)

	fmt.Println("After Insertion:")
	PrintList(head)
}