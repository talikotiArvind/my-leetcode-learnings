def reversePrefix(word: str, ch: str) -> str:
        output = word.find(ch)
        if output == -1:
            return word
        return word[:output + 1][::-1] + word[output + 1:]
