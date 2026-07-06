from collections import deque

class Solution:
    def validPath(self, n: int, edges: List[List[int]], source: int, destination: int) -> bool:
        if source == destination:
            return True

        output = [[] for _ in range(n)]
        for u, v in edges:
            output[u].append(v)
            output[v].append(u)

        visited = [False] * n
        visited[source] = True
        queue = deque([source])

        while queue:
            node = queue.popleft()
            for nei in output[node]:
                if nei == destination:
                    return True
                if not visited[nei]:
                    visited[nei] = True
                    queue.append(nei)
        return False
