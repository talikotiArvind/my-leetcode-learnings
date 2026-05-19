package main

import "fmt"

type Node struct {
    Val    int
    Next   *Node
    Random *Node
}

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    cloneMap := make(map[*Node]*Node)

    // Pass 1: create clones
    curr := head
    for curr != nil {
        cloneMap[curr] = &Node{Val: curr.Val}
        curr = curr.Next
    }

    // Pass 2: wire pointers
    curr = head
    for curr != nil {
        if curr.Next != nil {
            cloneMap[curr].Next = cloneMap[curr.Next]
        }
        if curr.Random != nil {
            cloneMap[curr].Random = cloneMap[curr.Random]
        }
        curr = curr.Next
    }

    return cloneMap[head]
}

func main() {
    // Build: [7->nil, 13->0, 11->4, 10->2, 1->0]
    nodes := make([]*Node, 5)
    vals := []int{7, 13, 11, 10, 1}
    for i, v := range vals {
        nodes[i] = &Node{Val: v}
    }
    for i := 0; i < 4; i++ {
        nodes[i].Next = nodes[i+1]
    }
    nodes[0].Random = nil
    nodes[1].Random = nodes[0]
    nodes[2].Random = nodes[4]
    nodes[3].Random = nodes[2]
    nodes[4].Random = nodes[0]

    copied := copyRandomList(nodes[0])

    for copied != nil {
        r := -1
        if copied.Random != nil {
            r = copied.Random.Val
        }
        fmt.Printf("val=%d random=%d\n", copied.Val, r)
        copied = copied.Next
    }
}