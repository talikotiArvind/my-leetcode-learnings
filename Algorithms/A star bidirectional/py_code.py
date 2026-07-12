"""Bidirectional A* on a weighted graph with 2D coordinates.

Two heuristic searches run at once - one forward from start, one backward from
goal - and meet in the middle. To stay optimal, this uses the *consistent*
(averaged-potential) scheme rather than naively stopping when the frontiers
first touch:

    p_f(n) = (h(n, goal) - h(n, start)) / 2      # forward potential
    p_b(n) = -p_f(n)                             # backward potential

Each queue is keyed on g + potential, keeping reduced edge costs non-negative so
the search is still Dijkstra underneath. The offsets baked into the keys cancel,
so termination is the same rule as bidirectional Dijkstra:
stop once  top_f + top_b >= best_path_found.
"""

import heapq
import math
from typing import Dict, List, Tuple, Optional

Graph = Dict[str, List[Tuple[str, float]]]   # node -> [(neighbour, weight), ...]
Coords = Dict[str, Tuple[float, float]]       # node -> (x, y)


def euclidean(coords: Coords, a: str, b: str) -> float:
    (x1, y1), (x2, y2) = coords[a], coords[b]
    return math.hypot(x1 - x2, y1 - y2)


def reverse_graph(graph: Graph) -> Graph:
    """Edges reversed - the backward search runs Dijkstra on this."""
    rev: Graph = {node: [] for node in graph}
    for u, edges in graph.items():
        for v, w in edges:
            rev.setdefault(v, []).append((u, w))
    return rev


def _reconstruct(parent: Dict[str, Optional[str]], node: str) -> List[str]:
    path = []
    while node is not None:
        path.append(node)
        node = parent[node]
    path.reverse()
    return path


def _stitch(parent_f: Dict[str, Optional[str]],
            parent_b: Dict[str, Optional[str]],
            meet: str) -> List[str]:
    """start ... meet  (forward tree)  +  meet+1 ... goal (backward tree)."""
    forward = _reconstruct(parent_f, meet)
    node = parent_b[meet]
    while node is not None:
        forward.append(node)
        node = parent_b[node]
    return forward


def bidirectional_astar(graph: Graph, coords: Coords, start: str, goal: str
                        ) -> Tuple[Optional[List[str]], float]:
    if start == goal:
        return [start], 0.0
    rev = reverse_graph(graph)

    def p_f(n: str) -> float:
        return (euclidean(coords, n, goal) - euclidean(coords, n, start)) / 2.0

    def p_b(n: str) -> float:
        return -p_f(n)

    dist_f: Dict[str, float] = {start: 0.0}
    dist_b: Dict[str, float] = {goal: 0.0}
    parent_f: Dict[str, Optional[str]] = {start: None}
    parent_b: Dict[str, Optional[str]] = {goal: None}
    heap_f: List[Tuple[float, str]] = [(p_f(start), start)]   # (g + potential, node)
    heap_b: List[Tuple[float, str]] = [(p_b(goal), goal)]
    settled_f, settled_b = set(), set()

    best = math.inf
    meet: Optional[str] = None

    while heap_f and heap_b:
        if heap_f[0][0] + heap_b[0][0] >= best:      # frontiers have passed each other
            break

        if heap_f[0][0] <= heap_b[0][0]:             # expand the smaller frontier
            _, u = heapq.heappop(heap_f)
            if u in settled_f:
                continue
            settled_f.add(u)
            if u in dist_b and dist_f[u] + dist_b[u] < best:
                best, meet = dist_f[u] + dist_b[u], u
            for v, w in graph[u]:
                nd = dist_f[u] + w
                if nd < dist_f.get(v, math.inf):
                    dist_f[v] = nd
                    parent_f[v] = u
                    heapq.heappush(heap_f, (nd + p_f(v), v))
        else:
            _, u = heapq.heappop(heap_b)
            if u in settled_b:
                continue
            settled_b.add(u)
            if u in dist_f and dist_f[u] + dist_b[u] < best:
                best, meet = dist_f[u] + dist_b[u], u
            for v, w in rev[u]:
                nd = dist_b[u] + w
                if nd < dist_b.get(v, math.inf):
                    dist_b[v] = nd
                    parent_b[v] = u
                    heapq.heappush(heap_b, (nd + p_b(v), v))

    if meet is None:
        return None, math.inf
    return _stitch(parent_f, parent_b, meet), best


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
    path, cost = bidirectional_astar(graph, coords, "S", "T")
    print(f"Bidirectional A*  cost={cost:g}  path={' -> '.join(path)}")
  
