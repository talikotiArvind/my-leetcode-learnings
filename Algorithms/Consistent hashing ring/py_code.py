import bisect
from hashlib import blake2b

class HashRing:
    def __init__(self, nodes=(), vnodes=150):
        self.vnodes = vnodes
        self.ring = {}            # hash -> node
        self.keys = []            # sorted hashes
        for n in nodes:
            self.add(n)

    @staticmethod
    def _hash(key):
        return int.from_bytes(blake2b(key.encode(), digest_size=8).digest(), "big")

    def add(self, node):
        for i in range(self.vnodes):
            h = self._hash(f"{node}#{i}")
            self.ring[h] = node
            bisect.insort(self.keys, h)

    def remove(self, node):
        for i in range(self.vnodes):
            h = self._hash(f"{node}#{i}")
            if self.ring.pop(h, None) is not None:
                self.keys.pop(bisect.bisect_left(self.keys, h))

    def get(self, key):
        if not self.keys:
            return None
        idx = bisect.bisect(self.keys, self._hash(key)) % len(self.keys)  # wrap
        return self.ring[self.keys[idx]]