def rabin_karp(text, pat, base=256, mod=(1 << 61) - 1):
    n, m = len(text), len(pat)
    if m == 0 or m > n:
        return []
    high = pow(base, m - 1, mod)          # weight of the outgoing char
    ph = th = 0
    for i in range(m):
        ph = (ph * base + ord(pat[i])) % mod
        th = (th * base + ord(text[i])) % mod

    res = []
    for s in range(n - m + 1):
        if ph == th and text[s:s + m] == pat:      # verify — hashes collide
            res.append(s)
        if s < n - m:
            th = ((th - ord(text[s]) * high) * base + ord(text[s + m])) % mod
    return res
