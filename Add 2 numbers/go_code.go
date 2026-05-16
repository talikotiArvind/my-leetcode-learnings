
package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    result := &ListNode{}
    curr := result
    for l1 != nil || l2 != nil || carry != 0 {
        v1, v2 := 0, 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }
        
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }
        
        total := v1 + v2 + carry
        carry = total / 10
        curr.Next = &ListNode{Val: total % 10}
        curr = curr.Next
    }
    return result.Next
}

func build(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

func toSlice(node *ListNode) []int {
	result := []int{}
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func main() {
	l1 := build([]int{2, 4, 3}) // 342
	l2 := build([]int{5, 6, 4}) // 465
	fmt.Println(toSlice(addTwoNumbers(l1, l2))) // [7 0 8] → 807

	l1 = build([]int{9, 9, 9, 9, 9, 9, 9}) // 9999999
	l2 = build([]int{9, 9, 9, 9})           // 9999
	fmt.Println(toSlice(addTwoNumbers(l1, l2))) // [8 9 9 9 0 0 0 1] → 10009998
}
