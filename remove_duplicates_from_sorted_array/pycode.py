def removeDuplicates(self, nums: List[int]) -> int:
        length = len(nums)
        j = 1

        for i in range(1, length):
            if nums[i] != nums[i-1]:
                nums[j] = nums[i]
                j += 1
        return j
