def searchBST(root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        while root:
            if val == root.val:
                return root
            root = root.left if val < root.val else root.right
        return None
