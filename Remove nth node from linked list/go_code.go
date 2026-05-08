// Go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    fast, slow := dummy, dummy
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}

// Test
func main() {
    fmt.Println(toSlice(removeNthFromEnd(build([]int{1,2,3,4,5}), 2)))  // [1 2 3 5]
}
