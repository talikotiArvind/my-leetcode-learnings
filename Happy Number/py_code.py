class Solution:
    def isHappy(self, n: int) -> bool:
        found = set()
        while n != 1:
            n = sum(int(d) ** 2 for d in str(n))
            if n in found:
                return False
            found.add(n)
        return True
