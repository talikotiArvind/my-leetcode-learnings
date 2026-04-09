def isSymmetric(root: Optional[TreeNode]) -> bool:
        def is_mirror(n1: TreeNode, n2: TreeNode):
            if not n1 and not n2:
                return True
            if not n1 or not n2:
                return False
            return n1.val == n2.val and is_mirror(n1.left, n2.right) and is_mirror(n1.right, n2.left)
        return is_mirror(root.left, root.right)
