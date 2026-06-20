def sortArrayByParityII(nums: List[int]) -> List[int]:
        even, odd = 0, 1
        n = len(nums)
        while even < n:
            if nums[even] % 2 == 0:
                even += 2
            else:
                while nums[odd] % 2 == 1:
                    odd += 2
                nums[even], nums[odd] = nums[odd], nums[even]
        return nums
