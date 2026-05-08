# Python
def removeNthFromEnd(head, n):
    dummy = ListNode(0, head)
    fast = slow = dummy
    for _ in range(n + 1):
        fast = fast.next
    while fast:
        fast = fast.next
        slow = slow.next
    slow.next = slow.next.next
    return dummy.next

print(to_list(removeNthFromEnd(build([1,2,3,4,5]), 2)))  # [1,2,3,5]
