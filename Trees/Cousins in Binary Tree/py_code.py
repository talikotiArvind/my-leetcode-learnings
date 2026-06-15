from collections import deque

def isCousins(root: TreeNode, x: int, y: int) -> bool:
    if not root:
        return False
        
    # Queue stores pairs of (current_node, parent_node)
    queue = deque([(root, None)])
    
    while queue:
        level_size = len(queue)
        found_x = False
        found_y = False
        parent_x = None
        parent_y = None
        
        # Process all nodes at the current depth level
        for _ in range(level_size):
            node, parent = queue.popleft()
            
            if node.val == x:
                found_x = True
                parent_x = parent
            if node.val == y:
                found_y = True
                parent_y = parent
                
            # Add children to the queue for the next level
            if node.left:
                queue.append((node.left, node))
            if node.right:
                queue.append((node.right, node))
        
        # If both are found at the current level, verify their parents differ
        if found_x and found_y:
            return parent_x != parent_y
            
        # If only one is found at this level, they cannot be cousins
        if found_x or found_y:
            return False
            
    return False
