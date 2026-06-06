package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, maxSoFar int) int
	dfs = func(node *TreeNode, maxSoFar int) int {
		if node == nil {
			return 0
		}
		isGood := 0
		if node.Val >= maxSoFar {
			isGood = 1
		}
		newMax := maxSoFar
		if node.Val > newMax {
			newMax = node.Val
		}
		return isGood + dfs(node.Left, newMax) + dfs(node.Right, newMax)
	}
	return dfs(root, -1<<31)
}

func main() {
	//        3
	//       / \
	//      1   4
	//     /   / \
	//    3   1   5
	//
	// Good nodes: 3 (root), 3 (left-left), 4, 5  → answer = 4

	root := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 5},
		},
	}
	fmt.Println(goodNodes(root)) // 4

	// Edge: single node is always good
	single := &TreeNode{Val: 7}
	fmt.Println(goodNodes(single)) // 1

	// Path: 5 -> 3 -> 4 -> 3 (only 5 and 4 are good? No: 5->3 fails, 3 is root=good, 4>3=good, 3<4=bad)
	//        5
	//         \
	//          3
	//           \
	//            4
	//             \
	//              3
	chain := &TreeNode{Val: 5,
		Right: &TreeNode{Val: 3,
			Right: &TreeNode{Val: 4,
				Right: &TreeNode{Val: 3},
			},
		},
	}
	fmt.Println(goodNodes(chain)) // 2  (nodes 5 and 4)
}
