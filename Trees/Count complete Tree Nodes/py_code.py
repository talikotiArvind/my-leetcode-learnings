def countNodes(root: Optional[TreeNode]) -> int:

        def height(node: TreeNode) -> int:
            h = 0
            while node:
                h += 1
                node = node.left
            return h
        count = 0
        if not root:
            return 0
        left_height, right_height = height(root.left), height(root.right)
        if left_height == right_height:
            return (1 << left_height) + countNodes(root.right)
        else:
            return (1 << right_height) + countNodes(root.left)
        
