class Solution:
    def climbStairs(self, n: int) -> int:
        i, j = 1, 1
        for _ in range(n-1):
            i, j = j, i+j
        return j
        
