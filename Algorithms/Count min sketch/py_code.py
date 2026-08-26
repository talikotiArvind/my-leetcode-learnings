import math
from hashlib import blake2b

class CountMinSketch:
    def __init__(self, epsilon=0.001, delta=0.01):
        self.w = math.ceil(math.e / epsilon)        # width: error bound
        self.d = math.ceil(math.log(1 / delta))     # depth: confidence
        self.table = [[0] * self.w for _ in range(self.d)]
        self.total = 0

    def _cols(self, item):
        if isinstance(item, str):
            item = item.encode()
        h = blake2b(item, digest_size=16).digest()
        h1 = int.from_bytes(h[:8], "big")
        h2 = int.from_bytes(h[8:], "big") | 1
        for i in range(self.d):
            yield i, (h1 + i * h2) % self.w

    def add(self, item, count=1):
        self.total += count
        for row, col in self._cols(item):
            self.table[row][col] += count

    def estimate(self, item):
        return min(self.table[row][col] for row, col in self._cols(item))