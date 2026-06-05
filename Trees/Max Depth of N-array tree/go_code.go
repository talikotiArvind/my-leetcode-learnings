package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	for _, child := range root.Children {
		if d := maxDepth(child); d > max {
			max = d
		}
	}
	return 1 + max
}

func main() {
	// Tree structure:
	//         1
	//       / | \
	//      3  2  4
	//     / \   / \
	//    5   6 7   8
	//         |
	//         9
	// Expected max depth: 4

	node9 := &Node{Val: 9}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7, Children: []*Node{node9}}
	node8 := &Node{Val: 8}
	node3 := &Node{Val: 3, Children: []*Node{node5, node6}}
	node2 := &Node{Val: 2}
	node4 := &Node{Val: 4, Children: []*Node{node7, node8}}
	root := &Node{Val: 1, Children: []*Node{node3, node2, node4}}

	fmt.Println("Max depth:", maxDepth(root)) // 4

	// Edge cases
	fmt.Println("Nil root:", maxDepth(nil))                    // 0
	fmt.Println("Single node:", maxDepth(&Node{Val: 42}))      // 1
	fmt.Println("Two levels:", maxDepth(&Node{Val: 1, Children: []*Node{{Val: 2}}})) // 2
}
