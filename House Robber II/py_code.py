from typing import List


def rob(nums: List[int]) -> int:
    if len(nums) == 1:
        return nums[0]

    def rob_linear(houses: List[int]) -> int:
        prev, curr = 0, 0
        for money in houses:
            prev, curr = curr, max(curr, prev + money)
        return curr

    return max(rob_linear(nums[:-1]), rob_linear(nums[1:]))


if __name__ == "__main__":
    assert rob([2, 3, 2]) == 3
    assert rob([1, 2, 3, 1]) == 4
    assert rob([1, 1, 1, 1]) == 2
    assert rob([5]) == 5
    assert rob([0, 0]) == 0
    print("All tests passed.")
