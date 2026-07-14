def dutch_flag_sort(arr):
    """Sort an array of 0s, 1s and 2s in a single pass. Returns arr for convenience."""
    low, mid, high = 0, 0, len(arr) - 1
    while mid <= high:
        if arr[mid] == 0:
            arr[low], arr[mid] = arr[mid], arr[low]
            low += 1
            mid += 1
        elif arr[mid] == 1:
            mid += 1
        else:                               # arr[mid] == 2
            arr[mid], arr[high] = arr[high], arr[mid]
            high -= 1                        # note: mid is NOT advanced here
    return arr


def main():
    flag = [2, 0, 2, 1, 1, 0, 0, 2, 1]
    print(f"before : {flag}")
    print(f"after  : {dutch_flag_sort(flag)}")


if __name__ == "__main__":
    main()
