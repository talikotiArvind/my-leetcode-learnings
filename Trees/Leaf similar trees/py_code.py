def leafSimilar(root1: Optional[TreeNode], root2: Optional[TreeNode]) -> bool:
        def leaves(node1):
            if not node1:
                return
            if not node1.left and not node1.right:
                yield node1.val
                return
            yield from leaves(node1.left)
            yield from leaves(node1.right)
        return list(leaves(root1)) == list(leaves(root2))
