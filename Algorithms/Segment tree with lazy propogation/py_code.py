class LazySegTree:
    """Range add, range sum. Half-open intervals [l, r)."""

    def __init__(self, data):
        self.n = len(data)
        self.tree = [0] * (4 * self.n)
        self.lazy = [0] * (4 * self.n)
        self._build(1, 0, self.n, data)

    def _build(self, node, lo, hi, data):
        if hi - lo == 1:
            self.tree[node] = data[lo]
            return
        mid = (lo + hi) // 2
        self._build(2 * node, lo, mid, data)
        self._build(2 * node + 1, mid, hi, data)
        self.tree[node] = self.tree[2 * node] + self.tree[2 * node + 1]

    def _apply(self, node, lo, hi, val):
        self.tree[node] += val * (hi - lo)      # sum grows by val per element
        self.lazy[node] += val

    def _push(self, node, lo, hi):
        if self.lazy[node]:
            mid = (lo + hi) // 2
            self._apply(2 * node, lo, mid, self.lazy[node])
            self._apply(2 * node + 1, mid, hi, self.lazy[node])
            self.lazy[node] = 0

    def add(self, l, r, val, node=1, lo=0, hi=None):
        if hi is None:
            hi = self.n
        if r <= lo or hi <= l:
            return
        if l <= lo and hi <= r:
            self._apply(node, lo, hi, val)
            return
        self._push(node, lo, hi)
        mid = (lo + hi) // 2
        self.add(l, r, val, 2 * node, lo, mid)
        self.add(l, r, val, 2 * node + 1, mid, hi)
        self.tree[node] = self.tree[2 * node] + self.tree[2 * node + 1]

    def query(self, l, r, node=1, lo=0, hi=None):
        if hi is None:
            hi = self.n
        if r <= lo or hi <= l:
            return 0
        if l <= lo and hi <= r:
            return self.tree[node]
        self._push(node, lo, hi)
        mid = (lo + hi) // 2
        return (self.query(l, r, 2 * node, lo, mid)
                + self.query(l, r, 2 * node + 1, mid, hi))
