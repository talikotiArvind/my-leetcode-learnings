package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var dfs func(node *TreeNode, isLeft bool) int
	dfs = func(node *TreeNode, isLeft bool) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			if isLeft {
				return node.Val
			}
			return 0
		}
		return dfs(node.Left, true) + dfs(node.Right, false)
	}
	return dfs(root, false)
}

func main() {
	// [3, 9, 20, null, null, 15, 7] → 24
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(sumOfLeftLeaves(root)) // 24

	// [1] → 0
	fmt.Println(sumOfLeftLeaves(&TreeNode{Val: 1})) // 0

	// [1, 2] → 2
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	fmt.Println(sumOfLeftLeaves(root2)) // 2
}
