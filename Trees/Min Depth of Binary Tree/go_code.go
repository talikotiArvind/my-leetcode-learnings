package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func minDepthDFS(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    if root.Left == nil {
        return 1 + minDepthDFS(root.Right)
    }
    
    if root.Right == nil {
        return 1 + minDepthDFS(root.Left)
    }
    
    left := minDepthDFS(root.Left)
    right := minDepthDFS(root.Right)
    
    if  left < right {
        return 1 + left
    }
    return 1 + right
}

func main() {
  fmt.Println("Try programiz.pro")
  root := &TreeNode{Val: 3}
  root.Left = &TreeNode{Val: 9}
  root.Right = &TreeNode{Val: 20}
  root.Right.Left = &TreeNode{Val: 15}
  root.Right.Right = &TreeNode{Val: 7}
   
 fmt.Println(minDepthDFS(root))
}
