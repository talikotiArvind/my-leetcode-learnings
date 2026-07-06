package main

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make([]bool, n)
	visited[source] = true
	queue := []int{source}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, nei := range adj[node] {
			if nei == destination {
				return true
			}
			if !visited[nei] {
				visited[nei] = true
				queue = append(queue, nei)
			}
		}
	}

	return false
}

func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))                     // true
	fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))     // false
	fmt.Println(validPath(1, [][]int{}, 0, 0))                                            // true
}
