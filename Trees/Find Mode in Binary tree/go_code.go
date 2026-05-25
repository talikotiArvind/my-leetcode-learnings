package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var modes []int
	var prev *int
	count, maxCount := 0, 0

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		if prev != nil && *prev == node.Val {
			count++
		} else {
			count = 1
		}
		prev = &node.Val

		if count > maxCount {
			maxCount = count
			modes = []int{node.Val}
		} else if count == maxCount {
			modes = append(modes, node.Val)
		}

		inorder(node.Right)
	}

	inorder(root)
	return modes
}

func main() {
	// Tree: [1, null, 2, 2]
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 2}
	fmt.Println(findMode(root)) // [2]

	// Tree: [1, 1, 2, 1]  — mode is 1
	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 2}
	fmt.Println(findMode(root2)) // [2]

	// Single node
	root3 := &TreeNode{Val: 42}
	fmt.Println(findMode(root3)) // [42]

	// Multiple modes: [1, null, 2] both appear once
	root4 := &TreeNode{Val: 1}
	root4.Right = &TreeNode{Val: 2}
	fmt.Println(findMode(root4)) // [1 2]
}
