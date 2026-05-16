def addTwoNumbers(l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        result = ListNode(0)
        current = result
        carry = 0

        while l1 or l2 or carry:
            value_1 = l1.val if l1 else 0
            value_2  = l2.val if l2 else 0

            total = value_1 + value_2 + carry
            carry = total // 10
            current.next = ListNode(total % 10)

            current = current.next
            if l1: l1 = l1.next
            if l2: l2 = l2.next
        return result.next
