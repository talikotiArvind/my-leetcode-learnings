from typing import List


def coinChange(coins: List[int], amount: int) -> int:
    dp = [0] + [float("inf")] * amount
    for coin in coins:
        for a in range(coin, amount + 1):
            dp[a] = min(dp[a], dp[a - coin] + 1)
    return dp[amount] if dp[amount] != float("inf") else -1


if __name__ == "__main__":
    print(coinChange([1, 2, 5], 11))  # 3
    print(coinChange([2], 3))         # -1
    print(coinChange([1], 0))         # 0
