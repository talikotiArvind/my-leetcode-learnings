def pigeonhole_sort(arr):
    if not arr:
        return arr

    min_val = min(arr)
    max_val = max(arr)

    # Number of pigeonholes
    size = max_val - min_val + 1
    holes = [0] * size

    # Count each number
    for num in arr:
        holes[num - min_val] += 1

    # Put numbers back in sorted order
    index = 0

    for i in range(size):
        while holes[i] > 0:
            arr[index] = i + min_val
            index += 1
            holes[i] -= 1

    return arr


# Example
arr = [8, 3, 2, 7, 4, 6, 8]

print("Before:", arr)

pigeonhole_sort(arr)

print("After: ", arr)