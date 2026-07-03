package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}

func main() {
	// Tree:    10
	//         /  \
	//        4    6      -> 4 + 6 == 10
	root1 := &TreeNode{Val: 10, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}
	fmt.Println(checkTree(root1)) // true

	// Tree:     5
	//          / \
	//         3   1       -> 3 + 1 != 5
	root2 := &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}
	fmt.Println(checkTree(root2)) // false
}
