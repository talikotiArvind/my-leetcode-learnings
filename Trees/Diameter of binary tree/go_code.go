package main
import "fmt"

type TreeNode struct {
    Left *TreeNode
    Right *TreeNode
    Val int
}

func diameterOfBinaryTree (head *TreeNode) int {
    diameter := 0
    
    var depth func (*TreeNode) int
    depth = func (node *TreeNode) int {
        if node == nil { return 0}
        left := depth(node.Left)
        right := depth(node.Right)
        
        if left + right > diameter {
            diameter = left + right 
        }
        if left > right {
            return 1 + left
        }
        return 1 + right
    }
    depth(head)
    return diameter
}


func main() {
    root := &TreeNode{Val: 1}
	root.Left       = &TreeNode{Val: 2}
	root.Right      = &TreeNode{Val: 3}
	root.Left.Left  = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(diameterOfBinaryTree(root)) // 3
}
