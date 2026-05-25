from typing import Optional

class Solution:
    def findMode(self, root: Optional[TreeNode]) -> list[int]:
        modes = []
        prev = [None]
        count = [0]
        max_count = [0]

        def inorder(node):
            if not node:
                return
            inorder(node.left)

            count[0] = count[0] + 1 if node.val == prev[0] else 1
            prev[0] = node.val

            if count[0] > max_count[0]:
                max_count[0] = count[0]
                modes.clear()
                modes.append(node.val)
            elif count[0] == max_count[0]:
                modes.append(node.val)

            inorder(node.right)

        inorder(root)
        return modes
