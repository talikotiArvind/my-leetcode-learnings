def findIndices(nums: List[int], indexDifference: int, valueDifference: int) -> List[int]:
        min_idx = max_idx = 0
        for j in range(indexDifference, len(nums)):
            i = j - indexDifference  # newly eligible
            if nums[i] < nums[min_idx]:
                min_idx = i
            if nums[i] > nums[max_idx]:
                max_idx = i
            if nums[max_idx] - nums[j] >= valueDifference:
                return [max_idx, j]
            if nums[j] - nums[min_idx] >= valueDifference:
                return [min_idx, j]
        return [-1, -1]
