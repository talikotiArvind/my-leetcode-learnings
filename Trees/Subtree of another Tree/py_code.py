def isSubtree(root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        def isSameTree(s, t):
            if not s and not t:   return True
            if not s or not t:    return False
            return s.val == t.val and isSameTree(s.left, t.left) and isSameTree(s.right, t.right)
        
        def dfs(node):
            if not node:
                return False
            if isSameTree(node, subRoot):
                return True
            return dfs(node.left) or dfs(node.right)
        
        return dfs(root)
