def convertToBase7(num: int) -> str:
        if num == 0:
            return "0"
        neg = num < 0
        num = abs(num)
        digits = []
        while num > 0:
            digits.append(str(num % 7))   # remainder is the next digit
            num //= 7
        if neg:
            digits.append("-")
        return "".join(reversed(digits))  # digits came out least-significant first
