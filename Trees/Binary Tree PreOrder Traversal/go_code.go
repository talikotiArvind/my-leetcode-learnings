package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func preOrderTraversal(root *TreeNode) [] int {
    var output [] int
    
    var traverse func(*TreeNode)
    
    traverse = func(node *TreeNode) {
        if node == nil {return}
        output = append(output, node.Val)
        traverse(node.Left)
        traverse(node.Right)
    }
    traverse(root)
    return output
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}

	// Get the traversal result
	result := preOrderTraversal(root)

	// Output the result
	fmt.Println("Preorder Traversal:", result)
	// Expected Output: [1 2 5 3]
}
