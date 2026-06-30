def checkPerfectNumber(num: int) -> bool:
        if num <= 1:
            return False
        total = 1  # 1 is always a proper divisor for num > 1
        i = 2
        while i * i <= num:
            if num % i == 0:
                total += i
                if i != num // i:  # don't double-count the square root
                    total += num // i
            i += 1
        return total == num 
