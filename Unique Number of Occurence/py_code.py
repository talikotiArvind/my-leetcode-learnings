def uniqueOccurrences(arr: List[int]) -> bool:
        output = {i: arr.count(i) for i in set(arr)}
        result = list(output.values())
        return len(result) == len(set(result))
