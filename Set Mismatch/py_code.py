def findErrorNums(nums: List[int]) -> List[int]:
        n = len(nums)
        seen = set()
        dup = 0
        actual_sum = 0

        for x in nums:
            if x in seen:
                dup = x
            seen.add(x)
            actual_sum += x

        expected_sum = n * (n + 1) // 2
        missing = expected_sum - (actual_sum - dup)
        return [dup, missing]
