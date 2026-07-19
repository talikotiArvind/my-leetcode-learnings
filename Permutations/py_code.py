def permute(nums: List[int]) -> List[List[int]]:
        res = []

        def backtrack(start: int):
            if start == len(nums):
                res.append(nums[:])
                return
            for i in range(start, len(nums)):
                nums[start], nums[i] = nums[i], nums[start]   # fix nums[start]
                backtrack(start + 1)
                nums[start], nums[i] = nums[i], nums[start]   # swap back

        backtrack(0)
        return res
