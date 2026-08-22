from collections import deque

class AhoCorasick:
    def __init__(self, patterns):
        self.next = [{}]                  # goto edges per node
        self.fail = [0]
        self.out = [[]]                   # pattern indices ending here

        for idx, p in enumerate(patterns):
            node = 0
            for ch in p:
                nxt = self.next[node].get(ch)
                if nxt is None:
                    nxt = len(self.next)
                    self.next.append({})
                    self.fail.append(0)
                    self.out.append([])
                    self.next[node][ch] = nxt
                node = nxt
            self.out[node].append(idx)
        self._build()

    def _build(self):
        q = deque(self.next[0].values())          # depth-1 nodes fail to root
        while q:
            v = q.popleft()
            self.out[v] += self.out[self.fail[v]]  # inherit suffix matches
            for ch, nxt in self.next[v].items():
                f = self.fail[v]
                while f and ch not in self.next[f]:
                    f = self.fail[f]
                self.fail[nxt] = self.next[f].get(ch, 0)
                q.append(nxt)

    def search(self, text):
        """Yields (end_index, pattern_index) for every occurrence."""
        node = 0
        for i, ch in enumerate(text):
            while node and ch not in self.next[node]:
                node = self.fail[node]
            node = self.next[node].get(ch, 0)
            for idx in self.out[node]:
                yield i, idx