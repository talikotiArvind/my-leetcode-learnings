def isPowerOfFour(n: int) -> bool:
        return n > 0 and (n & (n - 1)) == 0 and n % 3 == 1
