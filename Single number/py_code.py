from typing import List


def singleNumber(nums: List[int]) -> int:
    result = 0
    for x in nums:
        result ^= x
    return result


if __name__ == "__main__":
    print(singleNumber([2, 2, 1]))        # 1
    print(singleNumber([4, 1, 2, 1, 2]))  # 4
    print(singleNumber([1]))              # 1