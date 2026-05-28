def findTilt(root: Optional[TreeNode]) -> int:
        self.total_tilt = 0

        def subtree_sum(node):
            if not node:
                return 0
            left = subtree_sum(node.left)
            right = subtree_sum(node.right)
            self.total_tilt += abs(left - right)
            return left + right + node.val
        
        subtree_sum(root)
        return self.total_tilt
