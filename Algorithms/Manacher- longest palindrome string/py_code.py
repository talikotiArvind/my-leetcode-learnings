def longest_palindrome(s):
    if not s:
        return ""
    t = "#" + "#".join(s) + "#"
    n = len(t)
    p = [0] * n                       # p[i] = radius, == length in original s
    center = right = best = 0
    for i in range(n):
        if i < right:
            p[i] = min(right - i, p[2 * center - i])    # mirror
        while (i - p[i] - 1 >= 0 and i + p[i] + 1 < n
               and t[i - p[i] - 1] == t[i + p[i] + 1]):
            p[i] += 1
        if i + p[i] > right:
            center, right = i, i + p[i]
        if p[i] > p[best]:
            best = i
    start = (best - p[best]) // 2
    return s[start:start + p[best]]