from bisect import bisect_left
from typing import List


def lengthOfLIS(nums: List[int]) -> int:
    tails = []
    for num in nums:
        i = bisect_left(tails, num)
        if i == len(tails):
            tails.append(num)
        else:
            tails[i] = num
    return len(tails)


if __name__ == "__main__":
    print(lengthOfLIS([10, 9, 2, 5, 3, 7, 101, 18]))  # 4
    print(lengthOfLIS([0, 1, 0, 3, 2, 3]))             # 4
    print(lengthOfLIS([7, 7, 7, 7, 7, 7, 7]))          # 1
