package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	groupPrev := dummy

	for {
		kth := groupPrev
		for i := 0; i < k; i++ {
			kth = kth.Next
			if kth == nil {
				return dummy.Next
			}
		}

		groupNext := kth.Next

		prev := groupNext
		curr := groupPrev.Next

		for curr != groupNext {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		temp := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = temp
	}
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	nodes := make([]*ListNode, 5)

	for i := 0; i < 5; i++ {
		nodes[i] = &ListNode{Val: i + 1}
	}

	for i := 0; i < 4; i++ {
		nodes[i].Next = nodes[i+1]
	}

	head := reverseKGroup(nodes[0], 2)
	printList(head)
}