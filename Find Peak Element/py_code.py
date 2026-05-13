def findPeakElement(nums: List[int]) -> int:
        low, high = 0, len(nums) - 1
        max = 0
        result = None
        while low < high:
            mid = (low + high) // 2
            if nums[mid] < nums[mid + 1]:
                low = mid + 1
            else:
                high = mid
        return low
