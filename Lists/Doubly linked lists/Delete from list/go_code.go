package main

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func DeleteNode(node *Node) {
	if node == nil {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Prev = nil
	node.Next = nil
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

	fmt.Println("Before Deletion:")
	PrintList(n1)

	DeleteNode(n3)

	fmt.Println("After Deletion:")
	PrintList(n1)
}