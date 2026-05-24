def findWords(words: List[str]) -> List[str]:
        rows = [
            set("qwertyuiop"),
            set("asdfghjkl"),
            set("zxcvbnm")
        ]

        result = []
        for word in words:
            lower = set(word.lower())
            if any(lower <= row for row in rows):
                result.append(word)
        return result
