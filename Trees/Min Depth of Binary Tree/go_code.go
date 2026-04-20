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

func minDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type item struct {
		node  *TreeNode
		depth int
	}

	queue := []item{{root, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		node, depth := curr.node, curr.depth

		// First leaf reached = minimum depth
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queue = append(queue, item{node.Left, depth + 1})
		}
		if node.Right != nil {
			queue = append(queue, item{node.Right, depth + 1})
		}
	}

	return 0
}

func main() {
  fmt.Println("Try programiz.pro")
  root := &TreeNode{Val: 3}
  root.Left = &TreeNode{Val: 9}
  root.Right = &TreeNode{Val: 20}
  root.Right.Left = &TreeNode{Val: 15}
  root.Right.Right = &TreeNode{Val: 7}
  
  
  fmt.Println(minDepthDFS(root))
  fmt.Println(minDepthBFS(root))
}
