def reverseStr(s: str, k: int) -> str:
        chars = list(s)
        for i in range(0, len(chars), 2*k):
            chars[i:i+k] = chars[i:i+k][::-1]
        return ''.join(chars)
