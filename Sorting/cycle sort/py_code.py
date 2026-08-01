def cycle_sort(arr):
    n = len(arr)

    for cycle_start in range(n - 1):
        item = arr[cycle_start]

        # Find where item should go
        pos = cycle_start
        for i in range(cycle_start + 1, n):
            if arr[i] < item:
                pos += 1

        # Item is already in correct position
        if pos == cycle_start:
            continue

        # Skip duplicates
        while item == arr[pos]:
            pos += 1

        arr[pos], item = item, arr[pos]

        # Rotate the rest of the cycle
        while pos != cycle_start:
            pos = cycle_start

            for i in range(cycle_start + 1, n):
                if arr[i] < item:
                    pos += 1

            while item == arr[pos]:
                pos += 1

            arr[pos], item = item, arr[pos]


arr = [5, 3, 2, 4, 1]

cycle_sort(arr)
print(arr)