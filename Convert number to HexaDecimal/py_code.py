def toHex(num: int) -> str:
        if num == 0:
            return "0"

        if num < 0:
            num = num + (1 << 32)

        chars = "0123456789abcdef"
        result = []
        while num > 0:
            result.append(chars[num & 0xf])
            num >>= 4
        return ''.join(reversed(result))
