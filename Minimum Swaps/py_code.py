def minimumSwaps(nums: list[int]) -> int:
        k = nums.count(0)
        region = len(nums) - k
        return sum(1 for i in range(region) if nums[i] == 0)
