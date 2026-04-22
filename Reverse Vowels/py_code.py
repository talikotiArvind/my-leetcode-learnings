def reverseVowels(s: str) -> str:
        vowel_list = ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']
        left, right = 0, len(s) - 1
        s = list(s)
        while left < right:
            if s[left] in vowel_list and s[right] in vowel_list:
                s[left], s[right] = s[right], s[left]
                left += 1
                right -= 1
                continue
            if s[left].lower() not in vowel_list:
                left += 1
            if s[right].lower() not in vowel_list:
                right -= 1
        return ''.join(s)
