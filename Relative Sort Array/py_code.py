def relativeSortArray(arr1: List[int], arr2: List[int]) -> List[int]:
        n = len(arr1)
        k = 0
        output = []
        for element in arr2:
            count_val = arr1.count(element)
            while count_val > 0:
                output.append(element)
                count_val -= 1
        bal = sorted([i for i in arr1 if i not in arr2])
        return output + bal
