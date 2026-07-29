class Node:
    def __init__(self, val):
        self.val = val
        self.prev = None
        self.next = None


def delete_node(node):
    if node is None:
        return

    if node.prev:
        node.prev.next = node.next

    if node.next:
        node.next.prev = node.prev

    node.prev = None
    node.next = None


def print_list(head):
    current = head
    while current:
        print(current.val, end="")
        if current.next:
            print(" <-> ", end="")
        current = current.next
    print()


if __name__ == "__main__":

    n1 = Node(1)
    n2 = Node(2)
    n3 = Node(3)
    n4 = Node(4)

    n1.next = n2
    n2.prev = n1

    n2.next = n3
    n3.prev = n2

    n3.next = n4
    n4.prev = n3

    print("Before Deletion:")
    print_list(n1)

    delete_node(n3)

    print("After Deletion:")
    print_list(n1)