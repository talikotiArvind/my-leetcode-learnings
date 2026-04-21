# XOR solution (fastest)
def findTheDifference(s: str, t: str) -> str:
        result = 0
        for char in s+t:
            result ^= ord(char)
        return chr(result)

# sum of differences of ASCII values
def findTheDifference(s: str, t: str) -> str:
        return chr(sum(ord(c) for c in t) - sum(ord(c) for c in s))
