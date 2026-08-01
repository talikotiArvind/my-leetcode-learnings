def comb_sort(arr):
    n = len(arr)

    gap = n
    shrink = 1.3
    sorted_flag = False

    while not sorted_flag:
        gap = int(gap / shrink)

        if gap <= 1:
            gap = 1
            sorted_flag = True

        i = 0

        while i + gap < n:
            if arr[i] > arr[i + gap]:
                arr[i], arr[i + gap] = arr[i + gap], arr[i]
                sorted_flag = False

            i += 1


arr = [8, 4, 1, 6, 3, 7, 2]

comb_sort(arr)
print(arr)