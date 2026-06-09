def preorder(root: 'Node') -> List[int]:
        if not root:
            return []
        result = [root.val]
        for child in root.children:
            result += preorder(child)
        return result
