class Node:
    def __init__(self, val=None, next=None):
        self.val = val
        self.next = next


class Solution:
    def insert(self, head: 'Node', insertVal: int) -> 'Node':

        newNode = Node(insertVal)

        if head is None:
            newNode.next = newNode
            return newNode

        cur = head

        while True:

            # Normal increasing order
            if cur.val <= insertVal <= cur.next.val:
                break

            # Maximum -> Minimum point
            if cur.val > cur.next.val:
                if insertVal >= cur.val or insertVal <= cur.next.val:
                    break

            cur = cur.next

            if cur == head:
                break

        newNode.next = cur.next
        cur.next = newNode

        return head


def printCircular(head, count):
    cur = head
    for _ in range(count):
        print(cur.val, end=" ")
        cur = cur.next
    print()


if __name__ == "__main__":
    n1 = Node(3)
    n2 = Node(4)
    n3 = Node(1)

    n1.next = n2
    n2.next = n3
    n3.next = n1

    head = Solution().insert(n1, 2)

    printCircular(head, 4)