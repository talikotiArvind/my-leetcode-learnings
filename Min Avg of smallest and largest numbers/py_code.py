def minimumAverage(nums: List[int]) -> float:
        nums.sort()
        left, right = 0, len(nums) - 1
        output = []
        while left < right:
            output.append((nums[left] + nums[right]) / 2)
            left += 1
            right -= 1
        return min(output)
