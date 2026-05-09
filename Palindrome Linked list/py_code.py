def isPalindrome(head: Optional[ListNode]) -> bool:
        slow, fast = head, head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next

        prev, current = None, slow
        while current:
            next_node = current.next
            current.next = prev
            prev = current
            current = next_node

        left, right = head, prev
        while right:
            if left.val != right.val:
                return False
            left, right = left.next, right.next
        return True
        
