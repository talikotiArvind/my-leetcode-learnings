from collections import deque

class Solution:
    def shortestPathLength(self, graph):
        n = len(graph)

        queue = deque()
        visited = set()

        for i in range(n):
            mask = 1 << i
            queue.append((i, mask, 0))
            visited.add((i, mask))

        final_mask = (1 << n) - 1

        while queue:
            node, mask, dist = queue.popleft()

            if mask == final_mask:
                return dist

            for nei in graph[node]:
                new_mask = mask | (1 << nei)

                if (nei, new_mask) not in visited:
                    visited.add((nei, new_mask))
                    queue.append((nei, new_mask, dist + 1))


if __name__ == "__main__":
    graph = [[1,2,3],[0],[0],[0]]
    print(Solution().shortestPathLength(graph))