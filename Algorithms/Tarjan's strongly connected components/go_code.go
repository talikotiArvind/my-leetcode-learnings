type tarjan struct {
	adj     [][]int
	index   []int
	low     []int
	onStack []bool
	stack   []int
	counter int
	sccs    [][]int
}

func TarjanSCC(n int, adj [][]int) [][]int {
	t := &tarjan{
		adj:     adj,
		index:   make([]int, n),
		low:     make([]int, n),
		onStack: make([]bool, n),
	}
	for i := range t.index {
		t.index[i] = -1
	}
	for v := 0; v < n; v++ {
		if t.index[v] == -1 {
			t.dfs(v)
		}
	}
	return t.sccs
}

func (t *tarjan) dfs(v int) {
	t.index[v], t.low[v] = t.counter, t.counter
	t.counter++
	t.stack = append(t.stack, v)
	t.onStack[v] = true

	for _, w := range t.adj[v] {
		switch {
		case t.index[w] == -1:
			t.dfs(w)
			if t.low[w] < t.low[v] {
				t.low[v] = t.low[w]
			}
		case t.onStack[w]:
			if t.index[w] < t.low[v] {
				t.low[v] = t.index[w]
			}
		}
	}

	if t.low[v] == t.index[v] {
		var comp []int
		for {
			w := t.stack[len(t.stack)-1]
			t.stack = t.stack[:len(t.stack)-1]
			t.onStack[w] = false
			comp = append(comp, w)
			if w == v {
				break
			}
		}
		t.sccs = append(t.sccs, comp)
	}
}