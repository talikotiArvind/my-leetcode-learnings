package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func secondMinimum(root *TreeNode) int {
    answer, minVal := -1, root.Val
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil{return}
        if node.Val > minVal {
            if answer == -1 || node.Val < answer {
                answer = node.Val
            }
            return
        }
        dfs(node.Left)
        dfs(node.Right)
    }
    
    dfs(root)
    return answer
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
    fmt.Println(secondMinimum(root)) // 5
}
