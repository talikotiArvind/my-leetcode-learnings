# Pattern Signal Reference

A static lookup so pattern-naming doesn't rely purely on recall. Use this
to confirm a hunch, not as a replacement for asking the user leading
questions first (see `modes/learn-mode.md`).

| Signal in problem statement | Likely pattern |
|---|---|
| "contiguous subarray/substring", "longest/shortest ... with" | Sliding Window |
| "sorted array", "pairs/triplets that sum to" | Two Pointers |
| "all combinations/permutations", "generate all" | Backtracking |
| "maximum/minimum" + overlapping subproblems | Dynamic Programming |
| "count number of ways" | DP or combinatorics |
| "is it possible to" + constraints | DP, Greedy, or Graph |
| "shortest path", unweighted | BFS |
| "shortest path", weighted | Dijkstra |
| "connected components", "is there a cycle" | DFS / Union-Find |
| "level order" | BFS with a queue |
| "top K elements", "kth largest", "running median" | Heap |
| "find in sorted array", "first/last occurrence" | Binary Search |
| "find minimum X such that" | Binary search on the answer |
| local choice leads to global optimum (needs proof) | Greedy |

## Two Pointers vs Sliding Window

Two Pointers: pointers move independently, often toward each other, no
notion of a "window" being maintained. Sliding Window: a contiguous range
that grows/shrinks, tracking some running property of everything inside it.

## DFS vs Backtracking

DFS: traverse/visit, doesn't undo state. Backtracking: build a partial
solution, explore, then undo ("path.pop()") to try the next branch.

## DP vs Greedy

DP: must compare subproblem results because the locally-best choice isn't
provably globally best. Greedy: the locally-best choice is provably
globally best — if you can't prove it, it's probably DP.

## Complexity Cheatsheet

```
O(1) < O(log n) < O(n) < O(n log n) < O(n²) < O(2^n) < O(n!)
```

| Structure | Access | Search | Insert | Delete |
|---|---|---|---|---|
| Array | O(1) | O(n) | O(1) amortized (end) | O(n) |
| Hash Map/Set | — | O(1) avg | O(1) avg | O(1) avg |
| Balanced BST | — | O(log n) | O(log n) | O(log n) |
| Heap | O(1) min/max | — | O(log n) | O(log n) |

DFS/BFS on a graph: O(V + E) time, O(V) space.
