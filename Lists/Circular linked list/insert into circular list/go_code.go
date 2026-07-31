package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func insert(head *Node, insertVal int) *Node {
	newNode := &Node{Val: insertVal}

	if head == nil {
		newNode.Next = newNode
		return newNode
	}

	cur := head

	for {
		if cur.Val <= insertVal && insertVal <= cur.Next.Val {
			break
		}

		if cur.Val > cur.Next.Val {
			if insertVal >= cur.Val || insertVal <= cur.Next.Val {
				break
			}
		}

		cur = cur.Next

		if cur == head {
			break
		}
	}

	newNode.Next = cur.Next
	cur.Next = newNode

	return head
}

func printCircular(head *Node, count int) {
	cur := head
	for i := 0; i < count; i++ {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func main() {
	n1 := &Node{Val: 3}
	n2 := &Node{Val: 4}
	n3 := &Node{Val: 1}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n1

	head := insert(n1, 2)

	printCircular(head, 4)
}