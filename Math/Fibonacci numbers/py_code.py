def fib(n: int) -> int:
        if n <= 1:
            return n
        m, x = 0, 1

        for _ in range(2, n + 1):
            m, x = x, m + x
        return x
