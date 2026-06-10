from collections import deque
def averageOfLevels(root: Optional[TreeNode]) -> List[float]:
        result = []
        queue = deque([root])
        while queue:
            level_size = len(queue)
            level_sum = 0

            for _ in range(level_size):
                node = queue.popleft()
                level_sum += node.val
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            result.append(level_sum/level_size)
        return result
