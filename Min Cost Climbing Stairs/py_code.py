def minCostClimbingStairs(cost: List[int]) -> int:
  first, second = cost[0], cost[1]
  for i in range(2, len(cost)):
      current = cost[i] + min(first, second)
      first = second
      second = current
  return min(first, second)
