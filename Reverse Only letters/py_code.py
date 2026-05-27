def reverseOnlyLetters(s: str) -> str:
        chars = list(s)
        l, r = 0, len(chars) - 1

        while l < r:
            if not chars[l].isalpha():
                l += 1
            elif not chars[r].isalpha():
                r -= 1
            else:
                chars[l], chars[r] = chars[r], chars[l]
                l += 1
                r -= 1
        return ''.join(chars)
