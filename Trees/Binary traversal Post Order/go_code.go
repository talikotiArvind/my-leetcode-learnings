
package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func postOrderTraversal(root *TreeNode) [] int {
    var output [] int
    
    var traverse func(*TreeNode)
    
    traverse = func(node *TreeNode) {
        if node == nil {return}
        traverse(node.Left)
        traverse(node.Right)
        output = append(output, node.Val)
    }
    traverse(root)
    return output
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}

	result := postOrderTraversal(root)

	fmt.Println("Postorder Traversal:", result)
	// Expected Output: [5 2 3 1]
}
