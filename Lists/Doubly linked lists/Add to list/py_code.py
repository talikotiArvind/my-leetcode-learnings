class Node:
    def __init__(self, val):
        self.val = val
        self.prev = None
        self.next = None


def insert_at_end(head, data):
    new_node = Node(data)

    # Empty list
    if head is None:
        return new_node

    current = head

    while current.next:
        current = current.next

    current.next = new_node
    new_node.prev = current

    return head


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

    n1.next = n2
    n2.prev = n1

    n2.next = n3
    n3.prev = n2

    print("Before Insertion:")
    print_list(n1)

    head = insert_at_end(n1, 4)

    print("After Insertion:")
    print_list(head)