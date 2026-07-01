def reorderList(head):
    if not head:
        return
    nodes = []
    cur = head
    while cur:
        nodes.append(cur)
        cur = cur.next

    i, j = 0, len(nodes) - 1
    while i < j:
        nodes[i].next = nodes[j]   # append first
        i += 1
        if i == j:                 # even count: landed on the shared middle
            break
        nodes[j].next = nodes[i]   # append last
        j -= 1
    nodes[i].next = None           # terminate at the true middle
