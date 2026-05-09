package main
import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {
		result = result<<1 | head.Val
		head = head.Next
	}
	return result
}

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
	fmt.Println(getDecimalValue(build([]int{1, 0, 1, 0}))) // 10
	fmt.Println(getDecimalValue(build([]int{1, 1, 1})))     // 7
	fmt.Println(getDecimalValue(build([]int{0})))           // 0
	fmt.Println(getDecimalValue(build([]int{1})))           // 1
	fmt.Println(getDecimalValue(build([]int{1, 0, 0, 0}))) // 8
}
