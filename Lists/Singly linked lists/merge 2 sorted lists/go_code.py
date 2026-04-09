// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	output := &ListNode{}
	current := output

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			current.Next = list2
			list2 = list2.Next
		} else {
			current.Next = list1
			list1 = list1.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	return output.Next
}

// Helper: create linked list from slice
func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

// Helper: print linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example input
	list1 := buildList([]int{1, 2, 4})
	list2 := buildList([]int{1, 3, 4})

	fmt.Print("List 1: ")
	printList(list1)

	fmt.Print("List 2: ")
	printList(list2)

	// Merge
	merged := mergeTwoLists(list1, list2)

	fmt.Print("Merged: ")
	printList(merged)
}
