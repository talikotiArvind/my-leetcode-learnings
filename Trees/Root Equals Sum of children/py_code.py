def checkTree(root: Optional[TreeNode]) -> bool:
        return root.val == root.left.val + root.right.val
