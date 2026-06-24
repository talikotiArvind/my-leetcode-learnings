def sortArrayByParityII(nums: List[int]) -> List[int]:
        i, j, n = 0, 1, len(nums)
        while i < n:
            if nums[i] % 2 == 0:
                i += 2
            else:
                while nums[j] % 2 == 1:
                    j += 2
                nums[i], nums[j] = nums[j], nums[i]

        return nums
