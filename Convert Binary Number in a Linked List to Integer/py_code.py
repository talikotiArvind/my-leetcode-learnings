def getDecimalValue(head: Optional[ListNode]) -> int:
        output = ""
        current = head
        while current:
            output += f"{current.val}"
            current = current.next
        return int(output, 2)
