def countGoodSubstrings(s: str) -> int:
        n = len(s)
        count = 0

        for i in range(n-2):
            if s[i] != s[i+1] and s[i+1]!= s[i+2] and s[i+2] != s[i]:
                count += 1
        return count #return count
