from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Base case: 0 or 1 node is already sorted
        if not head or not head.next:
            return head
        
        # Step 1: Split the list into two halves using slow/fast pointers
        prev, slow, fast = None, head, head
        while fast and fast.next:
            prev = slow
            slow = slow.next
            fast = fast.next.next
            
        # Disconnect the left half from the right half
        if prev:
            prev.next = None
            
        # Step 2: Recursively sort each half
        left = self.sortList(head)
        right = self.sortList(slow)
        
        # Step 3: Merge the sorted halves
        return self.merge(left, right)
    
    def merge(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)
        curr = dummy
        
        # Standard two-pointer merge process
        while l1 and l2:
            if l1.val < l2.val:
                curr.next = l1
                l1 = l1.next
            else:
                curr.next = l2
                l2 = l2.next
            curr = curr.next
            
        # Link any remaining nodes
        curr.next = l1 if l1 else l2
            
        return dummy.next
