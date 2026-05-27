def duplicateZeros(arr: List[int]) -> None:
        """
        Do not return anything, modify arr in-place instead.
        """
        n = len(arr)
        zeros = arr.count(0)
        i, j = n - 1, n + zeros - 1

        while i >= 0:
            if arr[i] == 0:
                if j < n:
                    arr[j] = 0
                j -= 1
                if j < n:
                    arr[j] = 0
            else:
                if j < n:
                    arr[j] = arr[i]
            i -= 1
            j -= 1
