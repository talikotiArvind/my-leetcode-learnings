from typing import List

class Solution:
    def nextPermutation(nums: List[int]) -> None:
        n = len(nums)

        # 1. Find first index from the right where nums[i] < nums[i+1]
        i = n - 2
        while i >= 0 and nums[i] >= nums[i + 1]:
            i -= 1

        # 2. If such a pivot exists, find the rightmost element > nums[i] and swap
        if i >= 0:
            j = n - 1
            while nums[j] <= nums[i]:
                j -= 1
            nums[i], nums[j] = nums[j], nums[i]

        # 3. Reverse the suffix after the pivot (it's in descending order)
        nums[i + 1:] = reversed(nums[i + 1:])
