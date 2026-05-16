def threeSum(nums: list[int]) -> list[list[int]]:
    nums.sort()
    result = []

    for i, a in enumerate(nums):
        # Skip duplicate values for the outer element
        if i > 0 and nums[i] == nums[i - 1]:
            continue
        # All remaining numbers are positive — no triplet can sum to 0
        if a > 0:
            break

        left, right = i + 1, len(nums) - 1
        while left < right:
            total = a + nums[left] + nums[right]

            if total < 0:
                left += 1
            elif total > 0:
                right -= 1
            else:
                result.append([a, nums[left], nums[right]])
                left += 1
                # Skip duplicates for the inner two pointers
                while left < right and nums[left] == nums[left - 1]:
                    left += 1

    return result
