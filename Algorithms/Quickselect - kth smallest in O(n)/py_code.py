import random

def quickselect(a, k):
    """k-th smallest, 0-indexed. Mutates a."""
    lo, hi = 0, len(a) - 1
    while lo < hi:
        pivot = a[random.randint(lo, hi)]
        lt, i, gt = lo, lo, hi                  # Dutch-flag partition
        while i <= gt:
            if a[i] < pivot:
                a[lt], a[i] = a[i], a[lt]
                lt += 1
                i += 1
            elif a[i] > pivot:
                a[gt], a[i] = a[i], a[gt]
                gt -= 1
            else:
                i += 1
        if k < lt:
            hi = lt - 1
        elif k > gt:
            lo = gt + 1
        else:
            return a[k]                          # k fell inside the equal block
    return a[lo]