def sortArrayByParity(nums: List[int]) -> List[int]:
        left, j = 0, 0
        while left < len(nums):
            if nums[left] % 2 == 0:
                nums[j], nums[left] = nums[left], nums[j]
                j += 1
            left += 1
        return nums
