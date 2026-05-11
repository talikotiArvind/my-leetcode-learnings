package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode (headA, headB *ListNode) *ListNode {
    a, b := headA, headB
    for a != b {
        if a != nil {
            a = a.Next
        } else {
            a = headB
        }
        
        if b != nil {
            b = b.Next
        } else {
            b = headA
        }
    }
    return a
}

func build(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func tail(node *ListNode) *ListNode {
	for node.Next != nil {
		node = node.Next
	}
	return node
}

func main() {
	// Case 1: Intersection at node 8
	// A: 4 → 1 → 8 → 4 → 5
	// B: 5 → 6 → 1 → 8 → 4 → 5
	shared := build([]int{8, 4, 5})
	listA := build([]int{4, 1})
	listB := build([]int{5, 6, 1})
	tail(listA).Next = shared
	tail(listB).Next = shared

	res := getIntersectionNode(listA, listB)
	fmt.Println("Case 1 - Intersection at:", res.Val) // 8
}
