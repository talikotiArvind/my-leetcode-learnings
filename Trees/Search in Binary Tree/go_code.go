package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func searchBST(root *TreeNode, target int) *TreeNode {
    for root != nil {
        if root.Val == target {
            return root
        }
        if target < root.Val {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return nil
}

func main() {
    // Tree: [2,2,5,null,null,5,7]
    root := &TreeNode{2,
        &TreeNode{2, nil, nil},
        &TreeNode{5,
            &TreeNode{5, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    fmt.Println(searchBST(root, 5))
}
