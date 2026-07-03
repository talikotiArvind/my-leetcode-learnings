def evaluateTree(root: TreeNode) -> bool:
    # Leaf: 0 -> False, 1 -> True
    if not root.left and not root.right:
        return root.val == 1
    left = evaluateTree(root.left)
    right = evaluateTree(root.right)
    if root.val == 2:      # OR
        return left or right
    return left and right  # val == 3 -> AND
