def arrangeCoins(n: int) -> int:
        low, high = 0, n
        coins_needed = 0
        while low <= high:
            mid = (low + high) // 2
            coins_needed = (mid)*(mid + 1) // 2
            if coins_needed == n:
                return mid
            elif coins_needed < n:
                low = mid + 1
            else:
                high = mid - 1
        return high
