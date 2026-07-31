package main

import "fmt"

type State struct {
	node int
	mask int
	dist int
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	finalMask := (1 << n) - 1

	queue := []State{}
	visited := make(map[[2]int]bool)

	for i := 0; i < n; i++ {
		mask := 1 << i
		queue = append(queue, State{i, mask, 0})
		visited[[2]int{i, mask}] = true
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.mask == finalMask {
			return cur.dist
		}

		for _, nei := range graph[cur.node] {
			newMask := cur.mask | (1 << nei)
			key := [2]int{nei, newMask}

			if !visited[key] {
				visited[key] = true
				queue = append(queue, State{nei, newMask, cur.dist + 1})
			}
		}
	}

	return -1
}

func main() {
	graph := [][]int{{1,2,3},{0},{0},{0}}
	fmt.Println(shortestPathLength(graph))
}