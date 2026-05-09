def removeElements(head: Optional[ListNode], val: int) -> Optional[ListNode]:
        dummy = ListNode(0)
        dummy.next = head
        current = dummy
        while current.next:
            if current.next.val != val:
                current = current.next
            else:
                current.next = current.next.next
        return dummy.next
