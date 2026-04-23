def summaryRanges(nums: List[int]) -> List[str]:
        output = []
        i = 0
        while i < len(nums):
            start = nums[i]
            while i + 1 < len(nums) and nums[i + 1] == nums[i] + 1:
                i += 1
            if nums[i] == start:
                output.append(str(start))
            else:
                output.append(f"{start}->{nums[i]}")
            i += 1
        return output
