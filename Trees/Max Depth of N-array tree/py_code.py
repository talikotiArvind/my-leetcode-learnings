def maxDepth(root) -> int:
        if not root:
            return 0
        
        def dfs(node, length):
            if not node:
                return 0
            result = 1
            for child in node.children:
                result = max(result, 1 + dfs(child, length+1))
            return result
        return dfs(root, 0)
