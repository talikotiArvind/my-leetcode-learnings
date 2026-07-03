package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	// Leaf: 0 -> false, 1 -> true
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}
	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)
	if root.Val == 2 { // OR
		return left || right
	}
	return left && right // Val == 3 -> AND
}

func main() {
	//          OR(2)
	//         /     \
	//     True(1)   AND(3)
	//               /    \
	//          False(0)  True(1)
	//
	// AND(false, true) = false ; OR(true, false) = true
	root := &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}}

	fmt.Println(evaluateTree(root)) // true
}
