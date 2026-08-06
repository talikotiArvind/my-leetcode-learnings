from typing import List


def rob(nums: List[int]) -> int:
    prev, curr = 0, 0  # prev = best up to i-2, curr = best up to i-1
    for x in nums:
        prev, curr = curr, max(curr, prev + x)
    return curr


if __name__ == "__main__":
    print(rob([1, 2, 3, 1]))     # 4  -> rob houses 1 and 3
    print(rob([2, 7, 9, 3, 1]))  # 12 -> rob houses 2, 9, 1
    print(rob([5]))              # 5