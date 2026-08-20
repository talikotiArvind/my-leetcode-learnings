INF = float('inf')

def bellman_ford(n, edges, src):
    """edges: list of (u, v, w). Returns (dist, parent)."""
    dist, parent = [INF] * n, [-1] * n
    dist[src] = 0
    for _ in range(n - 1):
        changed = False
        for u, v, w in edges:
            if dist[u] != INF and dist[u] + w < dist[v]:
                dist[v], parent[v] = dist[u] + w, u
                changed = True
        if not changed:
            break                      # early exit: already settled
    for u, v, w in edges:              # nth pass still relaxing => neg cycle
        if dist[u] != INF and dist[u] + w < dist[v]:
            raise ValueError("negative cycle reachable from source")
    return dist, parent