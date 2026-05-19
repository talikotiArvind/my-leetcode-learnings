from typing import Optional

class Node:
    def __init__(self, x, next=None, random=None):
        self.val = x
        self.next = next
        self.random = random

def copyRandomList(head: Optional[Node]) -> Optional[Node]:
    if not head:
        return None

    # Map: original node -> its clone
    clone_map = {}

    # Pass 1: create all clone nodes
    curr = head
    while curr:
        clone_map[curr] = Node(curr.val)
        curr = curr.next

    # Pass 2: wire next and random pointers
    curr = head
    while curr:
        if curr.next:
            clone_map[curr].next = clone_map[curr.next]
        if curr.random:
            clone_map[curr].random = clone_map[curr.random]
        curr = curr.next

    return clone_map[head]

# Build helper
def build(data):
    if not data:
        return None
    nodes = [Node(v) for v, _ in data]
    for i, (_, r) in enumerate(data):
        nodes[i].next = nodes[i + 1] if i + 1 < len(nodes) else None
        nodes[i].random = nodes[r] if r is not None else None
    return nodes[0]

def to_list(head):
    nodes, result = [], []
    curr = head
    while curr:
        nodes.append(curr)
        curr = curr.next
    index = {n: i for i, n in enumerate(nodes)}
    for n in nodes:
        result.append([n.val, index[n.random] if n.random else None])
    return result

head = build([[7,None],[13,0],[11,4],[10,2],[1,0]])
print(to_list(copyRandomList(head)))
# Output: [[7, None], [13, 0], [11, 4], [10, 2], [1, 0]]