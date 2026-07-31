class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseKGroup(self, head, k):
        dummy = ListNode(0)
        dummy.next = head
        groupPrev = dummy

        while True:
            kth = groupPrev
            for _ in range(k):
                kth = kth.next
                if not kth:
                    return dummy.next

            groupNext = kth.next

            prev = groupNext
            curr = groupPrev.next

            while curr != groupNext:
                nxt = curr.next
                curr.next = prev
                prev = curr
                curr = nxt

            temp = groupPrev.next
            groupPrev.next = kth
            groupPrev = temp


def printList(head):
    while head:
        print(head.val, end=" ")
        head = head.next
    print()


if __name__ == "__main__":
    nodes = [ListNode(i) for i in range(1, 6)]

    for i in range(4):
        nodes[i].next = nodes[i + 1]

    head = Solution().reverseKGroup(nodes[0], 2)
    printList(head)