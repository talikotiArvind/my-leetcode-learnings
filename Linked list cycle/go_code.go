package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	node3 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node1 := &ListNode{Val: 1}
	node3.Next = node2
	node2.Next = node1
	node1.Next = node3 // cycle back to node3

	fmt.Println(hasCycle(node3)) // Output: true
}