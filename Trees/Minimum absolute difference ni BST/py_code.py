def getMinimumDifference(root: Optional[TreeNode]) -> int:
        self.min_diff = float('inf')
        self.previous = None

        def inorder(node):
            if not node:
                return
            inorder(node.left)
            if self.previous is not None:
                self.min_diff = min(self.min_diff, node.val - self.previous)
            self.previous = node.val
            inorder(node.right)
        inorder(root)
        return self.min_diff
