package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	minDiff := int(^uint(0) >> 1) // MaxInt
	prev := -1

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != -1 && node.Val-prev < minDiff {
			minDiff = node.Val - prev
		}
		prev = node.Val
		inorder(node.Right)
	}

	inorder(root)
	return minDiff
}

func main() {
	// Tree 1:
	//      4
	//     / \
	//    2   6
	//   / \
	//  1   3
	// Inorder: 1 2 3 4 6
	// Diffs:    1 1 1 2  → min = 1
	root1 := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println("Tree 1:", getMinimumDifference(root1)) // 1

	// Tree 2:
	//      1
	//       \
	//        5
	//       / \
	//      3   6
	// Inorder: 1 3 5 6
	// Diffs:    2 2 1  → min = 1
	root2 := &TreeNode{Val: 1,
		Right: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 6},
		},
	}
	fmt.Println("Tree 2:", getMinimumDifference(root2)) // 1

	// Tree 3: large gap except one close pair
	//        100
	//       /   \
	//      50   200
	//             \
	//             201
	// Inorder: 50 100 200 201
	// Diffs:    50  100   1  → min = 1
	root3 := &TreeNode{Val: 100,
		Left: &TreeNode{Val: 50},
		Right: &TreeNode{Val: 200,
			Right: &TreeNode{Val: 201},
		},
	}
	fmt.Println("Tree 3:", getMinimumDifference(root3)) // 1

	// Edge case: two nodes only
	//    5
	//     \
	//      8
	// min = 3
	root4 := &TreeNode{Val: 5, Right: &TreeNode{Val: 8}}
	fmt.Println("Two nodes:", getMinimumDifference(root4)) // 3
}
