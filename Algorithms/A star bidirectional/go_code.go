// Bidirectional A* on a weighted graph with 2D coordinates.
//
// Two heuristic searches run at once - one forward from start, one backward
// from goal - and meet in the middle. To stay optimal, this uses the
// *consistent* (averaged-potential) scheme rather than naively stopping when
// the frontiers first touch:
//
//	pf(n) = (h(n,goal) - h(n,start)) / 2   // forward potential
//	pb(n) = -pf(n)                         // backward potential
//
// Each queue is keyed on g + potential, keeping reduced edge costs non-negative
// so the search is still Dijkstra underneath. The offsets baked into the keys
// cancel, so termination is the same rule as bidirectional Dijkstra:
// stop once topF + topB >= best.
//
// Run standalone:  go run bidirectional_astar.go
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

func reverseGraph(g Graph) Graph {
	rev := make(Graph, len(g))
	for u := range g {
		if _, ok := rev[u]; !ok {
			rev[u] = nil
		}
	}
	for u, edges := range g {
		for _, e := range edges {
			rev[e.To] = append(rev[e.To], Edge{To: u, W: e.W})
		}
	}
	return rev
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
func (pq priorityQueue) peek() float64 { return pq[0].priority } // min key; Len()>0 only

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

// start..meet (forward tree) + meet+1..goal (backward tree)
func stitch(parentF, parentB map[string]string, meet string) []string {
	path := reconstruct(parentF, meet)
	for n := parentB[meet]; n != ""; n = parentB[n] {
		path = append(path, n)
	}
	return path
}

func BiAStar(g Graph, c Coords, start, goal string) ([]string, float64, bool) {
	if start == goal {
		return []string{start}, 0, true
	}
	rev := reverseGraph(g)
	pf := func(n string) float64 { return (euclid(c, n, goal) - euclid(c, n, start)) / 2 }
	pb := func(n string) float64 { return -pf(n) }

	distF := map[string]float64{start: 0}
	distB := map[string]float64{goal: 0}
	parentF := map[string]string{start: ""}
	parentB := map[string]string{goal: ""}
	settledF := map[string]bool{}
	settledB := map[string]bool{}

	pqF := &priorityQueue{{node: start, priority: pf(start)}}
	pqB := &priorityQueue{{node: goal, priority: pb(goal)}}
	heap.Init(pqF)
	heap.Init(pqB)

	best := math.Inf(1)
	meet := ""

	for pqF.Len() > 0 && pqB.Len() > 0 {
		if pqF.peek()+pqB.peek() >= best {
			break
		}
		if pqF.peek() <= pqB.peek() { // expand smaller frontier
			u := heap.Pop(pqF).(pqItem).node
			if settledF[u] {
				continue
			}
			settledF[u] = true
			if db, ok := distB[u]; ok && distF[u]+db < best {
				best, meet = distF[u]+db, u
			}
			for _, e := range g[u] {
				nd := distF[u] + e.W
				if old, ok := distF[e.To]; !ok || nd < old {
					distF[e.To] = nd
					parentF[e.To] = u
					heap.Push(pqF, pqItem{node: e.To, priority: nd + pf(e.To)})
				}
			}
		} else {
			u := heap.Pop(pqB).(pqItem).node
			if settledB[u] {
				continue
			}
			settledB[u] = true
			if df, ok := distF[u]; ok && df+distB[u] < best {
				best, meet = df+distB[u], u
			}
			for _, e := range rev[u] {
				nd := distB[u] + e.W
				if old, ok := distB[e.To]; !ok || nd < old {
					distB[e.To] = nd
					parentB[e.To] = u
					heap.Push(pqB, pqItem{node: e.To, priority: nd + pb(e.To)})
				}
			}
		}
	}

	if meet == "" {
		return nil, math.Inf(1), false
	}
	return stitch(parentF, parentB, meet), best, true
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
	path, cost, _ := BiAStar(g, coords, "S", "T")
	fmt.Printf("Bidirectional A*  cost=%g  path=%s\n", cost, strings.Join(path, " -> "))
}
