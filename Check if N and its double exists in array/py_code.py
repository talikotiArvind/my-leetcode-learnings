def checkIfExist(arr: List[int]) -> bool:
        seen = set()
        for index, value in enumerate(arr):
            if (2 * value) in arr and index != arr.index(2 * value):
                return True
        return False
