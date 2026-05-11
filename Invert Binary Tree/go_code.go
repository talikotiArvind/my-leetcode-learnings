package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func invertTree (root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
    return root
}

func build(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)
		if node.Left != nil  { queue = append(queue, node.Left) }
		if node.Right != nil { queue = append(queue, node.Right) }
	}
	return result
}

// ── main ─────────────────────────────────────────────────────────────────────

func main() {
	// Case 1: Full tree
	//        4                  4
	//      /   \      →       /   \
	//     2     7            7     2
	//    / \   / \          / \   / \
	//   1   3 6   9        9   6 3   1
	t1 := build([]any{4, 2, 7, 1, 3, 6, 9})
	fmt.Println("Case 1:", levelOrder(invertTree(t1))) // [4 7 2 9 6 3 1]

	// Case 2: Three nodes
	//     2            2
	//    / \    →    / \
	//   1   3       3   1
	t2 := build([]any{2, 1, 3})
	fmt.Println("Case 2:", levelOrder(invertTree(t2))) // [2 3 1]

	// Case 3: Single node
	t3 := build([]any{1})
	fmt.Println("Case 3:", levelOrder(invertTree(t3))) // [1]

	// Case 4: Empty tree
	t4 := invertTree(nil)
	fmt.Println("Case 4:", t4) // <nil>

	// Case 5: Left-skewed tree
	//   1              1
	//  /       →        \
	// 2                  2
	//  \                /
	//   3              3
	t5 := build([]any{1, 2, nil, nil, 3})
	fmt.Println("Case 5:", levelOrder(invertTree(t5))) // [1 2 3]
}
