def bm_search(text, pat):
    n, m = len(text), len(pat)
    if m == 0 or m > n:
        return []
    last = {c: i for i, c in enumerate(pat)}      # bad-character table
    shift = _good_suffix(pat)
    res, s = [], 0
    while s <= n - m:
        j = m - 1
        while j >= 0 and pat[j] == text[s + j]:
            j -= 1
        if j < 0:
            res.append(s)
            s += shift[0]
        else:
            bad_char = j - last.get(text[s + j], -1)
            s += max(1, bad_char, shift[j + 1])
    return res


def _good_suffix(pat):
    m = len(pat)
    shift, border = [0] * (m + 1), [0] * (m + 1)
    i, j = m, m + 1
    border[i] = j
    while i > 0:                                  # case 1: matching border
        while j <= m and pat[i - 1] != pat[j - 1]:
            if shift[j] == 0:
                shift[j] = j - i
            j = border[j]
        i, j = i - 1, j - 1
        border[i] = j
    j = border[0]                                 # case 2: prefix is a suffix
    for i in range(m + 1):
        if shift[i] == 0:
            shift[i] = j
        if i == j:
            j = border[j]
    return shift