package main

import (
	"container/heap"
	"fmt"
	"math"
)

// An Item is one (town, guess) pair sitting in the pile.
type Item struct {
	town string
	dist int
}

// PriorityQueue implements heap.Interface so the closest town is always at the front.
type PriorityQueue []Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist } // smallest dist first
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any)         { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// One road from a town: where it goes and how long it is.
type Edge struct {
	to     string
	length int
}

func dijkstra(graph map[string][]Edge, start string) map[string]int {
	// "Best guess so far" for every town. Start = 0, the rest = infinity.
	dist := make(map[string]int)
	for town := range graph {
		dist[town] = math.MaxInt
	}
	dist[start] = 0

	// The pile that always hands us the closest unfinished town first.
	pq := &PriorityQueue{{town: start, dist: 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		// Grab the closest unfinished town and "lock it in".
		cur := heap.Pop(pq).(Item)

		// Skip stale entries (an old, worse guess that got beaten later).
		if cur.dist > dist[cur.town] {
			continue
		}

		// Check each road: does it give a shorter route to a neighbor?
		for _, e := range graph[cur.town] {
			newGuess := cur.dist + e.length
			if newGuess < dist[e.to] {
				dist[e.to] = newGuess // update the neighbor's guess
				heap.Push(pq, Item{town: e.to, dist: newGuess})
			}
		}
	}
	return dist
}

func main() {
	graph := map[string][]Edge{
		"A": {{"B", 4}, {"C", 1}},
		"B": {{"A", 4}, {"C", 2}, {"D", 1}},
		"C": {{"A", 1}, {"B", 2}, {"D", 5}},
		"D": {{"B", 1}, {"C", 5}, {"E", 3}},
		"E": {{"D", 3}},
	}

	dist := dijkstra(graph, "A")

	// Print in a fixed order so the output is stable.
	for _, town := range []string{"A", "B", "C", "D", "E"} {
		fmt.Printf("%s = %d\n", town, dist[town])
	}
}
