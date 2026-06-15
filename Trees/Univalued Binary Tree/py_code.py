def isUnivalTree(root: Optional[TreeNode]) -> bool:
        value = root.val
        def search(node, value):
            if not node:
                return True
            if node.val != value:
                return False
            return search(node.left, value) and search(node.right, value)
        return search(root, root.val)
