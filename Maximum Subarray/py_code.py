from typing import List


def maxSubArray(nums: List[int]) -> int:
    best = current = nums[0]
    for x in nums[1:]:
        current = max(x, current + x)
        best = max(best, current)
    return best


if __name__ == "__main__":
    print(maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4]))  # 6  -> [4,-1,2,1]
    print(maxSubArray([1]))        # 1
    print(maxSubArray([5, 4, -1, 7, 8]))  # 23