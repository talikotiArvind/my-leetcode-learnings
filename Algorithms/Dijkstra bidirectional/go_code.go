// Bidirectional Dijkstra on a weighted graph.
//
// Forward Dijkstra from start + backward Dijkstra from goal (on the reversed
// graph). No heuristic and no coordinates are needed. best = length of the
// cheapest complete path discovered so far; the search is provably done once
// the two frontier minima satisfy  topF + topB >= best.
//
// Run standalone:  go run bidirectional_dijkstra.go
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

func BiDijkstra(g Graph, start, goal string) ([]string, float64, bool) {
	if start == goal {
		return []string{start}, 0, true
	}
	rev := reverseGraph(g)

	distF := map[string]float64{start: 0}
	distB := map[string]float64{goal: 0}
	parentF := map[string]string{start: ""}
	parentB := map[string]string{goal: ""}
	settledF := map[string]bool{}
	settledB := map[string]bool{}

	pqF := &priorityQueue{{node: start, priority: 0}}
	pqB := &priorityQueue{{node: goal, priority: 0}}
	heap.Init(pqF)
	heap.Init(pqB)

	best := math.Inf(1)
	meet := ""

	for pqF.Len() > 0 && pqB.Len() > 0 {
		if pqF.peek()+pqB.peek() >= best {
			break
		}
		if pqF.peek() <= pqB.peek() { // expand smaller frontier
			cur := heap.Pop(pqF).(pqItem)
			u := cur.node
			if settledF[u] {
				continue
			}
			settledF[u] = true
			if db, ok := distB[u]; ok && distF[u]+db < best {
				best, meet = distF[u]+db, u
			}
			for _, e := range g[u] {
				nd := cur.priority + e.W
				if old, ok := distF[e.To]; !ok || nd < old {
					distF[e.To] = nd
					parentF[e.To] = u
					heap.Push(pqF, pqItem{node: e.To, priority: nd})
				}
			}
		} else {
			cur := heap.Pop(pqB).(pqItem)
			u := cur.node
			if settledB[u] {
				continue
			}
			settledB[u] = true
			if df, ok := distF[u]; ok && df+distB[u] < best {
				best, meet = df+distB[u], u
			}
			for _, e := range rev[u] {
				nd := cur.priority + e.W
				if old, ok := distB[e.To]; !ok || nd < old {
					distB[e.To] = nd
					parentB[e.To] = u
					heap.Push(pqB, pqItem{node: e.To, priority: nd})
				}
			}
		}
	}

	if meet == "" {
		return nil, math.Inf(1), false
	}
	return stitch(parentF, parentB, meet), best, true
}

func buildDemo() Graph {
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
	g := Graph{}
	for _, e := range undirected {
		if _, ok := g[e.u]; !ok {
			g[e.u] = nil
		}
		if _, ok := g[e.v]; !ok {
			g[e.v] = nil
		}
		g[e.u] = append(g[e.u], Edge{To: e.v, W: e.w})
		g[e.v] = append(g[e.v], Edge{To: e.u, W: e.w})
	}
	return g
}

func main() {
	g := buildDemo()
	path, cost, _ := BiDijkstra(g, "S", "T")
	fmt.Printf("Bidirectional Dijkstra  cost=%g  path=%s\n", cost, strings.Join(path, " -> "))
}
