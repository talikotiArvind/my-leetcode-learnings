package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leaves(node *TreeNode, out *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*out = append(*out, node.Val)
		return
	}
	leaves(node.Left, out)
	leaves(node.Right, out)
}

func leafSimilar(root1, root2 *TreeNode) bool {
	var l1, l2 []int
	leaves(root1, &l1)
	leaves(root2, &l2)
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func main() {
	// Tree 1: leaves -> [6, 7, 4, 9, 8]
	root1 := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}},
		Right: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8}}}

	// Tree 2: same leaf sequence -> [6, 7, 4, 9, 8]
	root2 := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}},
		Right: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8}}}

	fmt.Println(leafSimilar(root1, root2)) // true
}
