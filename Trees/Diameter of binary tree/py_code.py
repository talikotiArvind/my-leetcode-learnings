def diameterOfBinaryTree(root: Optional[TreeNode]) -> int:
        self.diameter = 0

        def depth_finder(node):
            if not node:
                return 0
            left = depth_finder(node.left)
            right = depth_finder(node.right)

            self.diameter = max(self.diameter, left + right)
            return 1 + max(left, right)
        depth_finder(root)
        return self.diameter
