package main
import "fmt"

type TreeNode struct {
    Left *TreeNode
    Right *TreeNode
    Val int
}

func dfs(node *TreeNode, value int) bool{
    if node == nil {return true}
    
    if node.Val != value {return false}
    
    return dfs(node.Left, value) && dfs(node.Right, value)
}

func UnivalueBinary(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return dfs(root, root.Val)
}

func main() {
  root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   10,
			Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
			Right: &TreeNode{Val: 1, Left: nil, Right: nil},
		},
		Right: &TreeNode{Val: 1, Left: nil, Right: nil},
	}

	// Evaluate the tree
	result := UnivalueBinary(root)
	fmt.Printf("Is the tree univalued? %t\n", result)
}
