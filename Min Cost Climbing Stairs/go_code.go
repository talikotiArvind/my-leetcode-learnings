func minCostClimbingStairs(cost []int) int {
    first := cost[0]
    second := cost[1]
    
    for i := 2; i < len(cost); i++ {
        current := cost[i] + min(first, second)
        first = second
        second = current
    }
    
    return min(first, second)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
