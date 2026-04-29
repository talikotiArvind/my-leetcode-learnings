package main
import "fmt"

type TreeNode struct{
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) bool {
    if root == nil {return false}
    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum
    }
    remaining := targetSum - root.Val
    return pathSum(root.Left, remaining) || pathSum(root.Right, remaining)
}

func main() {
  root := &TreeNode{Val: 5}
  root.Left = &TreeNode{Val: 4}
  root.Right = &TreeNode{Val: 8}
  root.Left.Left = &TreeNode{Val: 11}
  root.Left.Left.Left = &TreeNode{Val: 7}
  root.Left.Left.Right = &TreeNode{Val: 2}

  fmt.Println(pathSum(root, 22))
  fmt.Println(pathSum(root, 10))

}
