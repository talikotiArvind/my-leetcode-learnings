package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	var result []int
	stack := []*Node{root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, curr.Val)
		for _, child := range curr.Children {
			stack = append(stack, child)
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func main() {
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n3 := &Node{Val: 3, Children: []*Node{n5, n6}}
	n2 := &Node{Val: 2}
	n4 := &Node{Val: 4}
	root := &Node{Val: 1, Children: []*Node{n3, n2, n4}}

	fmt.Println(postorder(root))
}
