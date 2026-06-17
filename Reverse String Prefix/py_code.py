def reversePrefix(s: str, k: int) -> str:
        data = s[:k]
        data = data[::-1]
        return data+s[k:]
