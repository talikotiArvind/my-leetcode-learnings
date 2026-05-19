from typing import List
import math

def minEatingSpeed(piles: List[int], h: int) -> int:
    lo, hi = 1, max(piles)  # speed range: 1 to largest pile

    while lo < hi:
        mid = (lo + hi) // 2
        hours = sum(math.ceil(p / mid) for p in piles)

        if hours <= h:
            hi = mid        # mid works, try slower
        else:
            lo = mid + 1    # too slow, go faster

    return lo

print(minEatingSpeed([3, 6, 7, 11], 8))   # Output: 4
print(minEatingSpeed([30, 11, 23, 4, 20], 5))  # Output: 30
print(minEatingSpeed([1, 1, 1, 999999999], 10)) # Output: 142857143