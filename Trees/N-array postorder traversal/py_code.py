def postorder(root: 'Node') -> list[int]:
    if not root:
        return []
    
    result = []
    stack = [root]
    
    while stack:
        curr = stack.pop()
        result.append(curr.val)
        for child in curr.children:
            stack.append(child)
            
    return result[::-1]
