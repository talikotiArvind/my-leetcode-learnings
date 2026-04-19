package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    return check(root) != -1
}

func check (node *TreeNode) int {
    if node == nil {return 0}
    
    leftHeight := check(node.Left)
    if leftHeight == -1 {return -1}
    
    rightHeight := check(node.Right)
    if rightHeight == -1 {return -1}
    
    if abs(leftHeight - rightHeight) > 1 {return -1}
    
    return 1 + max(leftHeight, rightHeight)
}

func abs(x int) int {
    if x < 0 {return -x}
    return x
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func main() {

	balanced := newNode(3)
	balanced.Left = newNode(9)
	balanced.Right = newNode(20)
	balanced.Right.Left = newNode(15)
	balanced.Right.Right = newNode(7)
	balanced.Right.Right.Left = newNode(22)

	fmt.Println("Balanced tree:")
	fmt.Println(isBalanced(balanced))
}
