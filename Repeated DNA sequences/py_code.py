def findRepeatedDnaSequences(s: str) -> List[str]:
        seen, repeated = set(), set()
        for i in range(len(s) - 9):
            window = s[i:i+10]
            if window in seen:
                repeated.add(window)
            else:
                seen.add(window)
        return list(repeated)
