package main
import "fmt"


type TreeNode struct {
    Left *TreeNode
    Right *TreeNode
    Val int
}

func mergeBinaryTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil && root2 == nil {return nil}
    if root1 == nil {return root2}
    if root2 == nil {return root1}
    
    node:= &TreeNode{Val: root1.Val + root2.Val}
    node.Left = mergeBinaryTrees(root1.Left, root2.Left)
    node.Right = mergeBinaryTrees(root1.Right, root2.Right)
    
    return node
}

// helper to print tree level-order for verification
func printInorder(node *TreeNode) {
	if node == nil {
		return
	}
	printInorder(node.Left)
	fmt.Printf("%d ", node.Val)
	printInorder(node.Right)
}

func main() {
	// Example 1:
	// Tree 1:       Tree 2:        Expected Merged:
	//     1              2              3
	//    / \            / \            / \
	//   3   2          1   3          4   5
	//  /                \   \        / \   \
	// 5                  4   7      5   4   7

	root1 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 2},
	}
	root2 := &TreeNode{Val: 2,
		Left: &TreeNode{Val: 1,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 3,
			Right: &TreeNode{Val: 7},
		},
	}
	merged1 := mergeBinaryTrees(root1, root2)
	fmt.Print("Example 1 inorder: ")
	printInorder(merged1) // 5 4 4 3 5 7
	fmt.Println()

	// Example 2:
	// Tree 1:    Tree 2:    Expected:
	//     1          nil        1
	// One tree is nil — return the other as-is

	root3 := &TreeNode{Val: 1}
	merged2 := mergeBinaryTrees(root3, nil)
	fmt.Print("Example 2 inorder: ")
	printInorder(merged2) // 1
	fmt.Println()

	// Example 3:
	// Tree 1:    Tree 2:    Expected:
	//     2          3          5
	//    /          / \        / \
	//   1          1   4      2   4
	//    \                     \
	//     3                     3

	root5 := &TreeNode{Val: 2,
		Left: &TreeNode{Val: 1,
			Right: &TreeNode{Val: 3},
		},
	}
	root6 := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4},
	}
	merged3 := mergeBinaryTrees(root5, root6)
	fmt.Print("Example 3 inorder: ")
	printInorder(merged3) // 2 3 5 4
	fmt.Println()
}
