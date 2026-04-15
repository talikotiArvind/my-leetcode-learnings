def findMaxConsecutiveOnes(nums: List[int]) -> int:
        count = 0
        max_count = 0
        for i in nums:
            if i == 1:
                count += 1
            else:
                max_count = max(max_count, count)
                count = 0
        return max_count if max_count > count else count
