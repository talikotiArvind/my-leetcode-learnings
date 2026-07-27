package main

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// Base case: if list is empty or has only one node
	if head == nil || head.Next == nil {
		return head
	}

	// Split the list into two halves
	var prev *ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Cut the link between the first and second half
	if prev != nil {
		prev.Next = nil
	}

	// Recursively sort both halves
	left := sortList(head)
	right := sortList(slow)

	// Merge the sorted halves
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	// Append remaining nodes
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}

// Helper function to convert slice to LinkedList
func createLinkedList(arr []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, val := range arr {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper function to print LinkedList
func printLinkedList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d", curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println()
}

// Main Execution Example
func main() {
	testData := []int{4, 2, 1, 3}
	fmt.Println("Original List:")
	head := createLinkedList(testData)
	printLinkedList(head)

	// Sort
	sortedHead := sortList(head)

	fmt.Println("\nSorted List:")
	printLinkedList(sortedHead)
}
