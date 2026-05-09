
package main
import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func buildList(vals []int) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    for _, v := range vals {
        curr.Next = &ListNode{Val: v}
        curr = curr.Next
    }
    return dummy.Next
}

func deleteElementsfromList (head *ListNode, value int) *ListNode {
    dummy := &ListNode{Next: head}
    current := dummy
    
    for current.Next != nil {
        if current.Next.Val != value {
            current = current.Next
        } else {
            current.Next = current.Next.Next
        }
    }
    return dummy.Next
}

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
    head := buildList([]int{1, 2, 6, 3, 6})
    result := deleteElementsfromList(head, 3)
    printList(result)
