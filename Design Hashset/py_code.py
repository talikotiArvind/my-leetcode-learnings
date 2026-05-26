class MyHashSet:

    def __init__(self):
        self.size = 500
        self.buckets = [[] for _ in range(self.size)]

    def _hash(self, key: int) -> int:
        return key % self.size

    def add(self, key: int) -> None:
        h = self._hash(key)
        if key not in self.buckets[h]:
            self.buckets[h].append(key)

    def remove(self, key: int) -> None:
        h = self._hash(key)
        if key in self.buckets[h]:
            self.buckets[h].remove(key)
        

    def contains(self, key: int) -> bool:
        h = self._hash(key)
        return key in self.buckets[h]

