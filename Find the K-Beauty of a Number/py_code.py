def divisorSubstrings(num: int, k: int) -> int:
        data = str(num)
        return sum(
            1 for i in range(len(data) - k + 1)
            if (w := int(data[i:i+k])) and num % w == 0
        )
