package main
import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(node *TreeNode) int {
    h := 0
    for node != nil {
        h ++
        node = node.Left
    }
    return h
}

func countNodes(root *TreeNode) int {
    if root == nil {return 0}
    leftHeight, rightHeight := height(root.Left), height(root.Right)
    
    if leftHeight == rightHeight {
        return (1 << leftHeight) + countNodes(root.Right)
    }
    return (1 << rightHeight) + countNodes(root.Left)
    
}

func main() {
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3,
		},
	}

	fmt.Println(countNodes(root))
}
