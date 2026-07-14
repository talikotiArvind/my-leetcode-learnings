class Node:
    def __init__(self, val):
        self.val = val
        self.next = None


def detect_cycle(head):
    """Return the node where the cycle begins, or None if the list is acyclic."""
    slow = fast = head
    while fast and fast.next:               # phase 1: find a meeting point
        slow = slow.next
        fast = fast.next.next
        if slow is fast:                    # pointers met -> a cycle exists
            ptr = head                      # phase 2: find the cycle entry
            while ptr is not slow:
                ptr = ptr.next
                slow = slow.next
            return ptr
    return None                             # fast fell off the end -> no cycle


def main():
    a, b, c, d = Node(3), Node(2), Node(0), Node(-4)
    a.next, b.next, c.next, d.next = b, c, d, b   # tail links back to b (value 2)
    start = detect_cycle(a)
    print(f"cyclic list  : has cycle = {start is not None}, enters at value {start.val}")

    x, y, z = Node(1), Node(2), Node(3)
    x.next, y.next = y, z
    print(f"acyclic list : has cycle = {detect_cycle(x) is not None}")


if __name__ == "__main__":
    main()
