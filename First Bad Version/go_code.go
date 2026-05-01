func firstBadVersion(n int) int {
    left, right := 0, n
    for ;left <= right; {
        mid := left + (right - left) / 2
        if isBadVersion(mid) {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left
}
