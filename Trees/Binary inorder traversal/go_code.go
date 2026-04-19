package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func BinaryInorderTraversal (root *TreeNode) [] int {
    result := [] int {}
    
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {return}
        dfs(node.Left)
        result = append(result, node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return result
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func main() {
    root := newNode(4)
    root.Left = newNode(2)
	root.Right = newNode(6)
	root.Left.Left = newNode(1)
	root.Left.Right = newNode(3)
	root.Right.Left = newNode(5)
	root.Right.Right = newNode(7)

	fmt.Println(BinaryInorderTraversal(root))
}
