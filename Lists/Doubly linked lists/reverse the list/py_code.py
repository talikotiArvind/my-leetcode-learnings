class Node:
    def __init__(self, val):
        self.val = val
        self.prev = None
        self.next = None


def reverse(head):
    current = head
    new_head = None

    while current:
        current.prev, current.next = current.next, current.prev
        new_head = current
        current = current.prev

    return new_head


def print_list(head):
    current = head
    while current:
        print(current.val, end="")
        if current.next:
            print(" <-> ", end="")
        current = current.next
    print()


if __name__ == "__main__":
    # Create: 1 <-> 2 <-> 3 <-> 4

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

    print("Original List:")
    print_list(n1)

    head = reverse(n1)

    print("Reversed List:")
    print_list(head)