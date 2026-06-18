def myAtoi(s: str) -> int:
        INT_MAX, INT_MIN = 2 ** 31 - 1, -2 ** 31
        s = s.lstrip()
        if not s:
            return 0
        i, n = 0, len(s)
        print(f"s : {s}")
        sign = 1
        if s[i] == "+":
            sign = 1
            i += 1
        elif s[i] == "-":
            sign = -1
            i += 1

        result = 0
        while i < n and s[i].isdigit():
            print(i, n, s[i])
            digit = int(s[i])
            result = result * 10 + digit

            if sign * result <= INT_MIN:
                return INT_MIN
            if sign * result >= INT_MAX:
                return INT_MAX
            i += 1
        return sign * result
