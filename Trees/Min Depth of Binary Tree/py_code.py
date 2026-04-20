# DFS approach
def minDepth(root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        
        if not root.left:
            return 1 + minDepth(root.right)
        
        if not root.right:
            return 1 + minDepth(root.left)
        
        return 1 + min(minDepth(root.left), minDepth(root.right))
