const Inf = math.MaxInt32

type Edge struct{ U, V, W int }

func BellmanFord(n int, edges []Edge, src int) ([]int, error) {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = Inf
	}
	dist[src] = 0

	for i := 0; i < n-1; i++ {
		changed := false
		for _, e := range edges {
			if dist[e.U] != Inf && dist[e.U]+e.W < dist[e.V] {
				dist[e.V] = dist[e.U] + e.W
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	for _, e := range edges {
		if dist[e.U] != Inf && dist[e.U]+e.W < dist[e.V] {
			return nil, errors.New("negative cycle reachable from source")
		}
	}
	return dist, nil
}