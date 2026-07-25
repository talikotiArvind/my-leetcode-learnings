import bisect

def findRadius(houses: List[int], heaters: List[int]) -> int:
    heaters.sort()
    result = 0
    for house in houses:
        iter_val = bisect.bisect_left(heaters, house)
        left = house - heaters[iter_val - 1] if iter_val > 0 else float('inf')
        right = heaters[iter_val] - house if iter_val < len(heaters) else float('inf')
        result = max(result, min(left, right))
    return result
  
