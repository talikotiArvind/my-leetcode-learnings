package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// Step 1: Find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: Reverse second half
	var prev *ListNode
	curr := slow
	for curr != nil {
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}

	// Step 3: Compare
	left, right := head, prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

// --- Helpers ---
func build(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func main() {
	fmt.Println(isPalindrome(build([]int{1, 2, 2, 1})))    // true
	fmt.Println(isPalindrome(build([]int{1, 2, 3, 2, 1}))) // true
	fmt.Println(isPalindrome(build([]int{1, 2, 3})))        // false
	fmt.Println(isPalindrome(build([]int{1})))              // true
	fmt.Println(isPalindrome(build([]int{1, 2})))           // false
}
