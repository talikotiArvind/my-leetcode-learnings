def thirdMax(nums: List[int]) -> int:
        result = set(nums)
        unique_val = sorted(result)

        if len(unique_val) < 3:
            return unique_val[-1]
        else:
            return unique_val[-3]
