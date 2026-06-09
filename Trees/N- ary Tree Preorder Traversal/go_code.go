package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return result
}

func main() {
	//        1
	//      / | \
	//     3  2  4
	//    / \
	//   5   6

	root := &Node{Val: 1, Children: []*Node{
		{Val: 3, Children: []*Node{
			{Val: 5},
			{Val: 6},
		}},
		{Val: 2},
		{Val: 4},
	}}

	fmt.Println(preorder(root)) // [1 3 5 6 2 4]
}
