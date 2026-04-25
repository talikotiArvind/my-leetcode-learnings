def commonChars(words: List[str]) -> List[str]:
        min_count = set(words[0])
        output = []
        for element in min_count:
            count = min([word.count(element) for word in words])
            output.extend(count * element)
        return output
