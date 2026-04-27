func guessNumber(n int) int {
    low, high := 0, n
    for {
        mid := (low + high) / 2
        switch guess(mid) {
            case -1:
            high = mid - 1
            case 1:
            low = mid + 1
            default:
            return mid
        }
    }
}
