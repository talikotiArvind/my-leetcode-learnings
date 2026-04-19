def isBalanced(root: Optional[TreeNode]) -> bool:
        def height_check(node):
            if not node:
                return 0
            
            left_height = height_check(node.left)
            if left_height == -1:
                return -1
            
            right_height = height_check(node.right)
            if right_height == -1:
                return -1

            if abs(left_height - right_height) > 1:
                return -1
            
            return 1 + max(left_height, right_height)
        
        return height_check(root) != -1
