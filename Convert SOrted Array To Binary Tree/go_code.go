package main

import "fmt"

// TreeNode defines the structure for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// sortedArrayToBST is the main function that converts the array.
func sortedArrayToBST(nums []int) *TreeNode {
    return buildBST(nums, 0, len(nums)-1)
}

// buildBST is the recursive helper function.
func buildBST(nums []int, left int, right int) *TreeNode {
    if left > right {
        return nil
    }

    // Calculate middle index (safe from overflow)
    mid := left + (right-left)/2
    
    // Create current root and recursively build subtrees
    root := &TreeNode{Val: nums[mid]}
    root.Left = buildBST(nums, left, mid-1)
    root.Right = buildBST(nums, mid+1, right)

    return root
}

// Example usage
func main() {
    // 1. Define a sorted input array
    nums := []int{-10, -3, 0, 5, 9}
    
    // 2. Convert to BST
    root := sortedArrayToBST(nums)
    
    // 3. Output check: Level-order print (simplified)
    fmt.Println("Root value:", root.Val)           // Should be 0
    fmt.Println("Left child:", root.Left.Val)     // Should be -3
    fmt.Println("Right child:", root.Right.Val)    // Should be 9
}
