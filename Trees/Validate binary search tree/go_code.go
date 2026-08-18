package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var valid func(node *TreeNode, low, high int) bool
	valid = func(node *TreeNode, low, high int) bool {
		if node == nil {
			return true
		}
		if node.Val <= low || node.Val >= high {
			return false
		}
		return valid(node.Left, low, node.Val) && valid(node.Right, node.Val, high)
	}
	return valid(root, math.MinInt64, math.MaxInt64)
}

func main() {
	// [2,1,3] -> true
	root := &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}}
	fmt.Println(isValidBST(root)) // true

	// [5,1,4,null,null,3,6] -> false
	bad := &TreeNode{5, &TreeNode{Val: 1}, &TreeNode{4, &TreeNode{Val: 3}, &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(bad)) // false
}
