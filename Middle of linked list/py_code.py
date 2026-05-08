# Python
def middleNode(head):
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    return slow

print(to_list(middleNode(build([1,2,3,4,5]))))    # [3,4,5]
print(to_list(middleNode(build([1,2,3,4,5,6]))))  # [4,5,6]
