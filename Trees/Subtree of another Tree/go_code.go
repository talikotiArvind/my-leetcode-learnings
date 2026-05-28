package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val &&
		isSameTree(s.Left, t.Left) &&
		isSameTree(s.Right, t.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func main() {
	// Example 1:
	//     3              4
	//    / \            / \
	//   4   5          1   2
	//  / \
	// 1   2
	// Expected: true

	subRoot1 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	root1 := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}
	fmt.Println(isSubtree(root1, subRoot1)) // true
	subRoot2 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	root2 := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 0},
			},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}
	fmt.Println(isSubtree(root2, subRoot2)) // false

	subRoot3 := &TreeNode{Val: 2}
	root3 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isSubtree(root3, subRoot3)) // true
}
