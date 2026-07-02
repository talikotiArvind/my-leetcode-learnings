def isPrefixString(s: str, words: List[str]) -> bool:
        result = ""
        for word in words:
            result += word
            if result == s:
                return True
            if len(result) > len(s):
                return False
        return False
