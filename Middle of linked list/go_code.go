// Go
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

// Test
func main() {
    fmt.Println(toSlice(middleNode(build([]int{1,2,3,4,5}))))    // [3 4 5]
    fmt.Println(toSlice(middleNode(build([]int{1,2,3,4,5,6}))))  // [4 5 6]
}
