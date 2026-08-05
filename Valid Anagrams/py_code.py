from collections import Counter


def isAnagram(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False
    return Counter(s) == Counter(t)


if __name__ == "__main__":
    print(isAnagram("anagram", "nagaram"))  # True
    print(isAnagram("rat", "car"))          # False
    print(isAnagram("a", "ab"))             # False