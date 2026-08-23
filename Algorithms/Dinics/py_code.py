from collections import deque

class Dinic:
    def __init__(self, n):
        self.n = n
        self.graph = [[] for _ in range(n)]      # edge = [to, cap, rev_index]

    def add_edge(self, u, v, cap):
        self.graph[u].append([v, cap, len(self.graph[v])])
        self.graph[v].append([u, 0, len(self.graph[u]) - 1])   # residual

    def _bfs(self, s, t):
        self.level = [-1] * self.n
        self.level[s] = 0
        q = deque([s])
        while q:
            u = q.popleft()
            for v, cap, _ in self.graph[u]:
                if cap > 0 and self.level[v] < 0:
                    self.level[v] = self.level[u] + 1
                    q.append(v)
        return self.level[t] >= 0

    def _dfs(self, u, t, pushed):
        if u == t:
            return pushed
        while self.it[u] < len(self.graph[u]):    # it[] = current-arc optimization
            e = self.graph[u][self.it[u]]
            v, cap = e[0], e[1]
            if cap > 0 and self.level[v] == self.level[u] + 1:
                d = self._dfs(v, t, min(pushed, cap))
                if d > 0:
                    e[1] -= d
                    self.graph[v][e[2]][1] += d
                    return d
            self.it[u] += 1                       # dead arc, never revisit
        return 0

    def max_flow(self, s, t):
        flow = 0
        while self._bfs(s, t):
            self.it = [0] * self.n
            while (pushed := self._dfs(s, t, float('inf'))) > 0:
                flow += pushed
        return flow
