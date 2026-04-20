# DFS approach
def minDepth(root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        
        if not root.left:
            return 1 + minDepth(root.right)
        
        if not root.right:
            return 1 + minDepth(root.left)
        
        return 1 + min(minDepth(root.left), minDepth(root.right))

# BFS (best approach)
from collections import deque
def minDepth(root):
        if not root:
            return 0

        queue = deque([(root, 1)])

        while queue:
            node, depth = queue.popleft()

            if not node.left and not node.right:
                return depth
            
            if node.left:
                queue.append((node.left, depth + 1))
            if node.right:
                queue.append((node.right, depth + 1))
