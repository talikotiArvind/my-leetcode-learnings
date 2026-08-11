def uniquePaths(m: int, n: int) -> int:
    row = [1] * n
    for _ in range(1, m):
        for j in range(1, n):
            row[j] += row[j - 1]
    return row[-1]


if __name__ == "__main__":
    print(uniquePaths(3, 7))  # 28
    print(uniquePaths(3, 2))  # 3
    print(uniquePaths(1, 1))  # 1