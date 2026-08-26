type LCA struct {
	log   int
	up    [][]int
	depth []int
}

func NewLCA(n int, adj [][]int, root int) *LCA {
	log := 1
	for 1<<log < n {
		log++
	}
	l := &LCA{log: log, up: make([][]int, log), depth: make([]int, n)}
	for k := range l.up {
		l.up[k] = make([]int, n)
		for i := range l.up[k] {
			l.up[k][i] = -1
		}
	}

	visited := make([]bool, n)
	visited[root] = true
	type frame struct{ v, p int }
	stack := []frame{{root, -1}}
	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		l.up[0][f.v] = f.p
		for _, w := range adj[f.v] {
			if !visited[w] {
				visited[w] = true
				l.depth[w] = l.depth[f.v] + 1
				stack = append(stack, frame{w, f.v})
			}
		}
	}

	for k := 1; k < log; k++ {
		for v := 0; v < n; v++ {
			if mid := l.up[k-1][v]; mid >= 0 {
				l.up[k][v] = l.up[k-1][mid]
			}
		}
	}
	return l
}

func (l *LCA) KthAncestor(v, k int) int {
	for i := 0; i < l.log && v >= 0; i++ {
		if k>>i&1 == 1 {
			v = l.up[i][v]
		}
	}
	return v
}

func (l *LCA) Query(u, v int) int {
	if l.depth[u] < l.depth[v] {
		u, v = v, u
	}
	u = l.KthAncestor(u, l.depth[u]-l.depth[v])
	if u == v {
		return u
	}
	for k := l.log - 1; k >= 0; k-- {
		if l.up[k][u] != l.up[k][v] {
			u, v = l.up[k][u], l.up[k][v]
		}
	}
	return l.up[0][u]
}