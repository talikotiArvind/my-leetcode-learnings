def longestPalindrome(s: str) -> str:
        from collections import Counter
        
        counts = Counter(s)
        length = 0
        odd_found = False

        for freq in counts.values():
            length += freq // 2 * 2   # take all even pairs
            if freq % 2 == 1:
                odd_found = True      # one odd character can go in center

        return length + (1 if odd_found else 0)
