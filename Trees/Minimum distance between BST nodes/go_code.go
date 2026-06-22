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

func minDiffInBST(root *TreeNode) int {
	prev := -1
	minDiff := math.MaxInt32

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != -1 {
			if diff := node.Val - prev; diff < minDiff {
				minDiff = diff
			}
		}
		prev = node.Val
		inorder(node.Right)
	}

	inorder(root)
	return minDiff
}

func main() {
	root := &TreeNode{
		Val:  4,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println(minDiffInBST(root)) // 1

	root2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 3}}}
	fmt.Println(minDiffInBST(root2)) // 2
}
