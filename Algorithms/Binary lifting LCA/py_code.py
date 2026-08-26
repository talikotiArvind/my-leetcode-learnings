class LCA:
    def __init__(self, n, adj, root=0):
        self.LOG = max(1, (n - 1).bit_length())
        self.up = [[-1] * n for _ in range(self.LOG)]
        self.depth = [0] * n

        visited = [False] * n
        visited[root] = True
        stack = [(root, -1)]
        while stack:                              # iterative: deep trees are fine
            v, p = stack.pop()
            self.up[0][v] = p
            for w in adj[v]:
                if not visited[w]:
                    visited[w] = True
                    self.depth[w] = self.depth[v] + 1
                    stack.append((w, v))

        for k in range(1, self.LOG):              # doubling
            for v in range(n):
                mid = self.up[k - 1][v]
                self.up[k][v] = -1 if mid < 0 else self.up[k - 1][mid]

    def kth_ancestor(self, v, k):
        for i in range(self.LOG):
            if v < 0:
                break
            if k >> i & 1:
                v = self.up[i][v]
        return v

    def query(self, u, v):
        if self.depth[u] < self.depth[v]:
            u, v = v, u
        u = self.kth_ancestor(u, self.depth[u] - self.depth[v])   # level them
        if u == v:
            return u
        for k in range(self.LOG - 1, -1, -1):     # climb as high as stays distinct
            if self.up[k][u] != self.up[k][v]:
                u, v = self.up[k][u], self.up[k][v]
        return self.up[0][u]