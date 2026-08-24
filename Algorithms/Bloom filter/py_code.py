import math
from hashlib import blake2b

class BloomFilter:
    def __init__(self, n, fp_rate=0.01):
        self.m = max(8, math.ceil(-n * math.log(fp_rate) / (math.log(2) ** 2)))
        self.k = max(1, round(self.m / n * math.log(2)))
        self.bits = bytearray((self.m + 7) // 8)

    def _indices(self, item):
        if isinstance(item, str):
            item = item.encode()
        d = blake2b(item, digest_size=16).digest()
        h1 = int.from_bytes(d[:8], "big")
        h2 = int.from_bytes(d[8:], "big") | 1        # odd => full cycle
        for i in range(self.k):
            yield (h1 + i * h2) % self.m             # Kirsch-Mitzenmacher

    def add(self, item):
        for idx in self._indices(item):
            self.bits[idx >> 3] |= 1 << (idx & 7)

    def __contains__(self, item):
        return all(self.bits[i >> 3] >> (i & 7) & 1 for i in self._indices(item))