def isPerfectSquare(num: int) -> bool:
        root = int(num ** (0.5))
        return root**2 == num
