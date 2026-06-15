package main

import "fmt"

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NodeInfo stores a node along with its parent tracking context.
type NodeInfo struct {
	Node   *TreeNode
	Parent *TreeNode
}

// isCousins checks if nodes x and y are cousins.
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	// Queue for BFS level-order traversal
	queue := []NodeInfo{{Node: root, Parent: nil}}

	for len(queue) > 0 {
		levelSize := len(queue)
		var parentX, parentY *TreeNode
		foundX, foundY := false, false

		// Process all nodes at the current depth layer
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:] // Pop operation

			if curr.Node.Val == x {
				foundX = true
				parentX = curr.Parent
			}
			if curr.Node.Val == y {
				foundY = true
				parentY = curr.Parent
			}

			// Add children to the queue for the next tree level
			if curr.Node.Left != nil {
				queue = append(queue, NodeInfo{Node: curr.Node.Left, Parent: curr.Node})
			}
			if curr.Node.Right != nil {
				queue = append(queue, NodeInfo{Node: curr.Node.Right, Parent: curr.Node})
			}
		}

		// If both are found at this depth level, verify they have different parents
		if foundX && foundY {
			return parentX != parentY
		}

		// If only one is found at this depth level, they cannot be cousins
		if foundX || foundY {
			return false
		}
	}

	return false
}

func main() {
	// Create a test binary tree:
	//        1
	//       / \
	//      2   3
	//     /     \
	//    4       5
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{Val: 5, Left: nil, Right: nil},
		},
	}

	// Node 4 and Node 5 are both at depth 3, but have different parents (2 and 3)
	x, y := 4, 5
	result := isCousins(root, x, y)
	fmt.Printf("Are nodes %d and %d cousins? %t\n", x, y, result)
	// Output: Are nodes 4 and 5 cousins? true

	// Node 2 and Node 3 are siblings (same parent), so they should return false
	x2, y2 := 2, 3
	result2 := isCousins(root, x2, y2)
	fmt.Printf("Are nodes %d and %d cousins? %t\n", x2, y2, result2)
	// Output: Are nodes 2 and 3 cousins? false
}
