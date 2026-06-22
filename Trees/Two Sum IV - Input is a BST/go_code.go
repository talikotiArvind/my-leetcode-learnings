package main
import "fmt"

type TreeNode struct {
    Left *TreeNode
    Right *TreeNode
    Val int
}

func twoSumBST(root *TreeNode, k int) bool {
    seen := make(map[int]struct{})
    
    var dfs func (node *TreeNode) bool
    dfs = func (node *TreeNode) bool {
        if node == nil {
            return false
        }
        if _, ok := seen[k - node.Val]; ok {
            return true
        }
        seen[node.Val] = struct{}{}
        return dfs(node.Left) || dfs(node.Right)
    }
    return dfs(root)
}

func main() {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}},
	}
	fmt.Println(twoSumBST(root, 9))  // true
	fmt.Println(twoSumBST(root, 28)) // false
	fmt.Println(twoSumBST(root, 4))  // false
}
