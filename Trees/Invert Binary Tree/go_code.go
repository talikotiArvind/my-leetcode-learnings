
package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func invertBinaryTree(root *TreeNode) *TreeNode {
    if root == nil {return nil}
    root.Left, root.Right = root.Right, root.Left
    invertBinaryTree(root.Left)
    invertBinaryTree(root.Right)
    return root
    
}

func printInorder(node *TreeNode) {
	if node == nil {
		return
	}
	printInorder(node.Left)
	fmt.Printf("%d ", node.Val)
	printInorder(node.Right)
}

func main() {
    root := &TreeNode{Val: 15}
  	root.Left = &TreeNode{Val: 19}
  	root.Right = &TreeNode{Val: 25}
  	root.Left.Right = &TreeNode{Val: 99}
  	
  
  	fmt.Print("Inorder before inversion: ")
  	printInorder(root) // Expected: 19 99 15 25
  	fmt.Println()
  	invertBinaryTree(root)
  	fmt.Print("Inorder after inversion:  ")
  	printInorder(root) // Expected: 25 15 99 19
  	fmt.Println()
}
