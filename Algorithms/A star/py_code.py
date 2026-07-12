"""A* shortest path on a weighted graph with 2D coordinates.

Heuristic: straight-line (Euclidean) distance to the goal. It is admissible and
consistent here because every edge weight is >= the straight-line distance
between its endpoints, which is what makes the closed-set optimisation safe.

    f(n) = g(n) + h(n)
    g(n) = best known cost from start to n
    h(n) = heuristic estimate from n to goal
"""

import heapq
import math
from typing import Dict, List, Tuple, Optional

Graph = Dict[str, List[Tuple[str, float]]]   # node -> [(neighbour, weight), ...]
Coords = Dict[str, Tuple[float, float]]       # node -> (x, y)


def euclidean(coords: Coords, a: str, b: str) -> float:
    (x1, y1), (x2, y2) = coords[a], coords[b]
    return math.hypot(x1 - x2, y1 - y2)


def reconstruct(parent: Dict[str, Optional[str]], node: str) -> List[str]:
    path = []
    while node is not None:
        path.append(node)
        node = parent[node]
    path.reverse()
    return path


def astar(graph: Graph, coords: Coords, start: str, goal: str
          ) -> Tuple[Optional[List[str]], float]:
    def h(n: str) -> float:
        return euclidean(coords, n, goal)

    open_heap: List[Tuple[float, float, str]] = [(h(start), 0.0, start)]  # (f, g, node)
    g_score: Dict[str, float] = {start: 0.0}
    parent: Dict[str, Optional[str]] = {start: None}
    closed = set()

    while open_heap:
        _, g, u = heapq.heappop(open_heap)
        if u in closed:                     # stale duplicate, already finalised
            continue
        if u == goal:
            return reconstruct(parent, goal), g
        closed.add(u)                       # consistent heuristic => g[u] is final

        for v, w in graph[u]:
            if v in closed:
                continue
            ng = g + w
            if ng < g_score.get(v, math.inf):
                g_score[v] = ng
                parent[v] = u
                heapq.heappush(open_heap, (ng + h(v), ng, v))

    return None, math.inf


def build_demo() -> Tuple[Graph, Coords]:
    coords: Coords = {
        "S": (0, 0), "A": (2, 1), "B": (2, -1), "C": (4, 2), "D": (4, -2),
        "E": (6, 1), "F": (6, -1), "G": (8, 2), "H": (8, -2), "T": (10, 0),
    }
    undirected = [
        ("S", "A", 3), ("S", "B", 3),
        ("A", "C", 3), ("A", "B", 2), ("B", "D", 3),
        ("C", "E", 5), ("C", "D", 4), ("D", "F", 3),
        ("E", "G", 3), ("E", "F", 2), ("F", "H", 3),
        ("G", "T", 3), ("G", "H", 4), ("H", "T", 3),
    ]
    graph: Graph = {n: [] for n in coords}
    for u, v, w in undirected:
        graph[u].append((v, float(w)))
        graph[v].append((u, float(w)))
    return graph, coords


if __name__ == "__main__":
    graph, coords = build_demo()
    path, cost = astar(graph, coords, "S", "T")
    print(f"A*  cost={cost:g}  path={' -> '.join(path)}")
  
