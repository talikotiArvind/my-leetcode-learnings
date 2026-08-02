class Solution:
    def maxCoins(self, nums: list[int]) -> int:
        balloons = [1] + nums + [1]
        n = len(balloons)

        dp = [[0] * n for _ in range(n)]

        # Length of interval
        for length in range(2, n):

            for left in range(n - length):
                right = left + length

                # k is the last balloon burst
                for k in range(left + 1, right):

                    coins = (
                        dp[left][k]
                        + balloons[left] * balloons[k] * balloons[right]
                        + dp[k][right]
                    )

                    dp[left][right] = max(dp[left][right], coins)

        return dp[0][n - 1]


if __name__ == "__main__":
    solution = Solution()

    nums = [3, 1, 5, 8]

    print(solution.maxCoins(nums))