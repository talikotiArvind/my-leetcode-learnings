"""Bidirectional Dijkstra on a weighted graph.

Forward Dijkstra from start + backward Dijkstra from goal (on the reversed
graph). No heuristic and no coordinates are needed. `best` is the length of the
cheapest complete path discovered so far; the search is provably done once the
two frontier minima satisfy  top_f + top_b >= best.
"""

import heapq
import math
from typing import Dict, List, Tuple, Optional

Graph = Dict[str, List[Tuple[str, float]]]   # node -> [(neighbour, weight), ...]


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


def bidirectional_dijkstra(graph: Graph, start: str, goal: str
                           ) -> Tuple[Optional[List[str]], float]:
    if start == goal:
        return [start], 0.0
    rev = reverse_graph(graph)

    dist_f: Dict[str, float] = {start: 0.0}
    dist_b: Dict[str, float] = {goal: 0.0}
    parent_f: Dict[str, Optional[str]] = {start: None}
    parent_b: Dict[str, Optional[str]] = {goal: None}
    heap_f: List[Tuple[float, str]] = [(0.0, start)]
    heap_b: List[Tuple[float, str]] = [(0.0, goal)]
    settled_f, settled_b = set(), set()

    best = math.inf
    meet: Optional[str] = None

    while heap_f and heap_b:
        if heap_f[0][0] + heap_b[0][0] >= best:      # frontiers have passed each other
            break

        if heap_f[0][0] <= heap_b[0][0]:             # expand the smaller frontier
            d, u = heapq.heappop(heap_f)
            if u in settled_f:
                continue
            settled_f.add(u)
            if u in dist_b and dist_f[u] + dist_b[u] < best:
                best, meet = dist_f[u] + dist_b[u], u
            for v, w in graph[u]:
                nd = d + w
                if nd < dist_f.get(v, math.inf):
                    dist_f[v] = nd
                    parent_f[v] = u
                    heapq.heappush(heap_f, (nd, v))
        else:
            d, u = heapq.heappop(heap_b)
            if u in settled_b:
                continue
            settled_b.add(u)
            if u in dist_f and dist_f[u] + dist_b[u] < best:
                best, meet = dist_f[u] + dist_b[u], u
            for v, w in rev[u]:
                nd = d + w
                if nd < dist_b.get(v, math.inf):
                    dist_b[v] = nd
                    parent_b[v] = u
                    heapq.heappush(heap_b, (nd, v))

    if meet is None:
        return None, math.inf
    return _stitch(parent_f, parent_b, meet), best


def build_demo() -> Graph:
    undirected = [
        ("S", "A", 3), ("S", "B", 3),
        ("A", "C", 3), ("A", "B", 2), ("B", "D", 3),
        ("C", "E", 5), ("C", "D", 4), ("D", "F", 3),
        ("E", "G", 3), ("E", "F", 2), ("F", "H", 3),
        ("G", "T", 3), ("G", "H", 4), ("H", "T", 3),
    ]
    nodes = {n for u, v, _ in undirected for n in (u, v)}
    graph: Graph = {n: [] for n in nodes}
    for u, v, w in undirected:
        graph[u].append((v, float(w)))
        graph[v].append((u, float(w)))
    return graph


if __name__ == "__main__":
    graph = build_demo()
    path, cost = bidirectional_dijkstra(graph, "S", "T")
    print(f"Bidirectional Dijkstra  cost={cost:g}  path={' -> '.join(path)}")
      
