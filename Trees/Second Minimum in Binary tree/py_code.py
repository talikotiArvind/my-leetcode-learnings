def findSecondMinimumValue(root: Optional[TreeNode]) -> int:
        self.answer = float('inf')
        min_val = root.val

        def dfs(node):
            if not node:
                return 
            if min_val < node.val < self.answer:
                self.answer = node.val
            
            if node.val == min_val:
                dfs(node.left)
                dfs(node.right)
        
        dfs(root)
        return self.answer if self.answer < float('inf') else -1
