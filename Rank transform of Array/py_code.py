def arrayRankTransform(arr: List[int]) -> List[int]:
        output = {value: index for index, value in enumerate(sorted(set(arr)), start=1)}
        return [output[value] for value in arr]
