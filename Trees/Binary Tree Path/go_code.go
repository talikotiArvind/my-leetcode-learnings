package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}
		if node.Left != nil {
			dfs(node.Left, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			dfs(node.Right, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	dfs(root, strconv.Itoa(root.Val))
	return res
}

func main() {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}

	paths := binaryTreePaths(root)
	fmt.Println(paths) // [1->2->5 1->3]
}
