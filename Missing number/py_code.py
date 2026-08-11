from typing import List


def missingNumber(nums: List[int]) -> int:
    result = len(nums)
    for i, x in enumerate(nums):
        result ^= i ^ x
    return result


if __name__ == "__main__":
    print(missingNumber([3, 0, 1]))           # 2
    print(missingNumber([0, 1]))              # 2
    print(missingNumber([9, 6, 4, 2, 3, 5, 7, 0, 1]))  # 8