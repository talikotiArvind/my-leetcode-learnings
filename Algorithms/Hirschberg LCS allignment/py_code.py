def _last_row(a, b):
    """LCS length of a against every prefix of b — one row only."""
    prev = [0] * (len(b) + 1)
    for ch in a:
        cur = [0] * (len(b) + 1)
        for j, cb in enumerate(b, 1):
            cur[j] = prev[j - 1] + 1 if ch == cb else max(prev[j], cur[j - 1])
        prev = cur
    return prev


def hirschberg(a, b):
    if not a or not b:
        return ""
    if len(a) == 1:
        return a if a in b else ""

    mid = len(a) // 2
    left = _last_row(a[:mid], b)                    # forward scores
    right = _last_row(a[mid:][::-1], b[::-1])       # backward scores
    k = max(range(len(b) + 1),                      # optimal crossing point
            key=lambda j: left[j] + right[len(b) - j])
    return hirschberg(a[:mid], b[:k]) + hirschberg(a[mid:], b[k:])