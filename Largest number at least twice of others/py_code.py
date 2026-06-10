def dominantIndex(nums: List[int]) -> int:
        max_value = max(nums)
        nums2 = sorted(nums)
        for i in nums2[:-1]:
            if (i * 2) > max_value:
                return -1
        return nums.index(max_value)
