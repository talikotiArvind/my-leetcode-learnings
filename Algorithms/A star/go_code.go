// A* shortest path on a weighted graph with 2D coordinates.
//
// Heuristic: straight-line (Euclidean) distance to the goal. It is admissible
// and consistent here because every edge weight is >= the straight-line
// distance between its endpoints, which is what makes the closed-set
// optimisation safe.
//
//	f(n) = g(n) + h(n)
//	g(n) = best known cost from start to n
//	h(n) = heuristic estimate from n to goal
//
// Run standalone:  go run astar.go
package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type Edge struct {
	To string
	W  float64
}

type Graph map[string][]Edge
type Coords map[string][2]float64

func euclid(c Coords, a, b string) float64 {
	pa, pb := c[a], c[b]
	dx, dy := pa[0]-pb[0], pa[1]-pb[1]
	return math.Sqrt(dx*dx + dy*dy)
}

// ---- Priority queue (container/heap) ----

type pqItem struct {
	node     string
	priority float64
}

type priorityQueue []pqItem

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x any)        { *pq = append(*pq, x.(pqItem)) }
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func reconstruct(parent map[string]string, node string) []string {
	var path []string
	for node != "" {
		path = append(path, node)
		node = parent[node]
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func AStar(g Graph, c Coords, start, goal string) ([]string, float64, bool) {
	h := func(n string) float64 { return euclid(c, n, goal) }

	gScore := map[string]float64{start: 0}
	parent := map[string]string{start: ""}
	closed := map[string]bool{}

	pq := &priorityQueue{{node: start, priority: h(start)}}
	heap.Init(pq)

	for pq.Len() > 0 {
		u := heap.Pop(pq).(pqItem).node
		if closed[u] { // stale duplicate
			continue
		}
		if u == goal {
			return reconstruct(parent, goal), gScore[goal], true
		}
		closed[u] = true // consistent heuristic => gScore[u] is final

		for _, e := range g[u] {
			if closed[e.To] {
				continue
			}
			ng := gScore[u] + e.W
			if old, ok := gScore[e.To]; !ok || ng < old {
				gScore[e.To] = ng
				parent[e.To] = u
				heap.Push(pq, pqItem{node: e.To, priority: ng + h(e.To)})
			}
		}
	}
	return nil, math.Inf(1), false
}

func buildDemo() (Graph, Coords) {
	coords := Coords{
		"S": {0, 0}, "A": {2, 1}, "B": {2, -1}, "C": {4, 2}, "D": {4, -2},
		"E": {6, 1}, "F": {6, -1}, "G": {8, 2}, "H": {8, -2}, "T": {10, 0},
	}
	type ue struct {
		u, v string
		w    float64
	}
	undirected := []ue{
		{"S", "A", 3}, {"S", "B", 3},
		{"A", "C", 3}, {"A", "B", 2}, {"B", "D", 3},
		{"C", "E", 5}, {"C", "D", 4}, {"D", "F", 3},
		{"E", "G", 3}, {"E", "F", 2}, {"F", "H", 3},
		{"G", "T", 3}, {"G", "H", 4}, {"H", "T", 3},
	}
	g := make(Graph, len(coords))
	for n := range coords {
		g[n] = nil
	}
	for _, e := range undirected {
		g[e.u] = append(g[e.u], Edge{To: e.v, W: e.w})
		g[e.v] = append(g[e.v], Edge{To: e.u, W: e.w})
	}
	return g, coords
}

func main() {
	g, coords := buildDemo()
	path, cost, _ := AStar(g, coords, "S", "T")
	fmt.Printf("A*  cost=%g  path=%s\n", cost, strings.Join(path, " -> "))
}
