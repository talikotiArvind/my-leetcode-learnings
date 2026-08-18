def hammingWeight(n: int) -> int:
    count = 0
    while n:
        n &= n - 1
        count += 1
    return count


if __name__ == "__main__":
    print(hammingWeight(11))          # 3   (1011)
    print(hammingWeight(128))         # 1   (10000000)
    print(hammingWeight(2147483645))  # 30
