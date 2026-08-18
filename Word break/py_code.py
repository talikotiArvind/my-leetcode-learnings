from typing import List


def wordBreak(s: str, wordDict: List[str]) -> bool:
    words = set(wordDict)
    dp = [False] * (len(s) + 1)
    dp[0] = True

    for i in range(1, len(s) + 1):
        for j in range(i):
            if dp[j] and s[j:i] in words:
                dp[i] = True
                break

    return dp[len(s)]


if __name__ == "__main__":
    print(wordBreak("leetcode", ["leet", "code"]))            # True
    print(wordBreak("applepenapple", ["apple", "pen"]))       # True
    print(wordBreak("catsandog", ["cats", "dog", "sand", "and", "cat"]))  # False
