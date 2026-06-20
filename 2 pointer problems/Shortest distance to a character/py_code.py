def shortestToChar(s: str, c: str) -> List[int]:
        n = len(s)
        output = [0] * n

        prev = -n
        for i in range(n):
            if s[i] == c:
                prev = i
            output[i] = i - prev

        prev = 2 * n
        for i in range(n-1, -1, -1):
            if s[i] == c:
                prev = i
            output[i] = min(output[i], prev - i)
        return output
