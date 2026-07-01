def flatten(root):
    cur = root
    while cur:
        if cur.left:
            # rightmost node of the left subtree's right spine
            pred = cur.left
            while pred.right:
                pred = pred.right
            pred.right = cur.right   # splice old right subtree after it
            cur.right = cur.left     # move left subtree to the right
            cur.left = None
        cur = cur.right
