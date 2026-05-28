package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	tilt := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		tilt += abs(left - right)
		return left + right + node.Val
	}
	dfs(root)
	return tilt
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	root1 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(findTilt(root1)) // 1

	root2 := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 9,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(findTilt(root2)) // 15
}
