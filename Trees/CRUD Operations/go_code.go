package tree

import "errors"

// DefaultRead defines traversal order
type DefaultRead string

const (
	InOrder   DefaultRead = "InOrder"
	PreOrder  DefaultRead = "PreOrder"
	PostOrder DefaultRead = "PostOrder"
)

// Node is the smallest unit of the tree
type Node struct {
	Value any
	Left  *Node
	Right *Node
}

// BinarySearchTree holds the root of the tree
type BinarySearchTree struct {
	Root *Node
}

// Insert adds a new value into the BST
func (bst *BinarySearchTree) Insert(value any) {
	newNode := &Node{Value: value}
	if bst.Root == nil {
		bst.Root = newNode
		return
	}

	current := bst.Root
	for {
		if lessThan(value, current.Value) {
			if current.Left == nil {
				current.Left = newNode
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				return
			}
			current = current.Right
		}
	}
}

// Search returns true if value exists in the BST
func (bst *BinarySearchTree) Search(value any) bool {
	current := bst.Root
	for current != nil {
		if equal(value, current.Value) {
			return true
		} else if lessThan(value, current.Value) {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return false
}

// Delete removes a value from the BST
func (bst *BinarySearchTree) Delete(value any) {
	bst.Root = deleteRecursive(bst.Root, value)
}

func deleteRecursive(node *Node, value any) *Node {
	if node == nil {
		return nil
	}
	if lessThan(value, node.Value) {
		node.Left = deleteRecursive(node.Left, value)
	} else if greaterThan(value, node.Value) {
		node.Right = deleteRecursive(node.Right, value)
	} else {
		// Case 1: leaf
		if node.Left == nil && node.Right == nil {
			return nil
		}
		// Case 2: one child
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		// Case 3: two children — replace with in-order successor
		successor := findMin(node.Right)
		node.Value = successor.Value
		node.Right = deleteRecursive(node.Right, successor.Value)
	}
	return node
}

func findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// Update replaces oldValue with newValue
func (bst *BinarySearchTree) Update(oldValue, newValue any) bool {
	if !bst.Search(oldValue) {
		return false
	}
	bst.Delete(oldValue)
	bst.Insert(newValue)
	return true
}

// Read traverses the tree in the given order and returns a slice of values
func (bst *BinarySearchTree) Read(order DefaultRead) ([]any, error) {
	switch order {
	case InOrder:
		return inorder(bst.Root, []any{}), nil
	case PreOrder:
		return preorder(bst.Root, []any{}), nil
	case PostOrder:
		return postorder(bst.Root, []any{}), nil
	default:
		return nil, errors.New("unknown traversal order: " + string(order))
	}
}

func inorder(node *Node, result []any) []any {
	if node == nil {
		return result
	}
	result = inorder(node.Left, result)
	result = append(result, node.Value)
	result = inorder(node.Right, result)
	return result
}

func preorder(node *Node, result []any) []any {
	if node == nil {
		return result
	}
	result = append(result, node.Value)
	result = preorder(node.Left, result)
	result = preorder(node.Right, result)
	return result
}

func postorder(node *Node, result []any) []any {
	if node == nil {
		return result
	}
	result = postorder(node.Left, result)
	result = postorder(node.Right, result)
	result = append(result, node.Value)
	return result
}

// --- Comparison helpers (Go has no generics operator overloading) ---

func lessThan(a, b any) bool {
	switch av := a.(type) {
	case int:
		return av < b.(int)
	case float64:
		return av < b.(float64)
	case string:
		return av < b.(string)
	}
	return false
}

func greaterThan(a, b any) bool {
	return lessThan(b, a)
}

func equal(a, b any) bool {
	return a == b
}
