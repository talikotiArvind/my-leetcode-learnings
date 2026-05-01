def firstBadVersion(n: int) -> int:
        if n == 1:
            return n
        version = n
        left, right = 0, n
        while left < right:
            mid = (left + right) // 2
            if isBadVersion(mid):
                version = min(version, mid)
                right = mid
            else:
                left = mid + 1
        return version
        
