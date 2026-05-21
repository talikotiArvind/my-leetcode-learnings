def binaryTreePaths(root: Optional[TreeNode]) -> List[str]:
        if not root:
            return []
        
        result = []

        def path_ride(node, path):
            if not node.left and not node.right:
                result.append(path)
                return
            if node.left:
                path_ride(node.left, path + "->" + str(node.left.val))
            if node.right:
                path_ride(node.right, path + "->" + str(node.right.val))

        path_ride(root, str(root.val))
        return result
        
