package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			// find the rightmost node of the left subtree's right spine
			pred := cur.Left
			for pred.Right != nil {
				pred = pred.Right
			}
			pred.Right = cur.Right // splice old right subtree after it
			cur.Right = cur.Left   // move left subtree to the right
			cur.Left = nil
		}
		cur = cur.Right
	}
}

func main() {
	//         1
	//        / \
	//       2   5
	//      / \   \
	//     3   4   6
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 5,
			Right: &TreeNode{Val: 6},
		},
	}

	flatten(root)

	var parts []string
	for n := root; n != nil; n = n.Right {
		parts = append(parts, fmt.Sprintf("%d", n.Val))
	}
	fmt.Println(strings.Join(parts, "->")) // 1->2->3->4->5->6
}
