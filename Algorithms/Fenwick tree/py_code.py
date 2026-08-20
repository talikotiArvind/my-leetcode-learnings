class Fenwick:
    def __init__(self, n):
        self.n, self.t = n, [0] * (n + 1)   # tree is 1-indexed

    def add(self, i, delta):                # i is 0-indexed
        i += 1
        while i <= self.n:
            self.t[i] += delta
            i += i & -i

    def prefix(self, i):                    # sum of [0, i)
        s = 0
        while i > 0:
            s += self.t[i]
            i -= i & -i
        return s

    def range_sum(self, l, r):              # [l, r)
        return self.prefix(r) - self.prefix(l)