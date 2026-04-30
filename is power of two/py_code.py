class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        if n == 1:
            return True
        if n < 0:
            return False
        return f"{n:b}".count('1') == 1
