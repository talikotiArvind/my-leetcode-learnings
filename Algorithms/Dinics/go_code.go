type edge struct {
	to, rev, cap int
}

type Dinic struct {
	graph [][]edge
	level []int
	iter  []int
}

func NewDinic(n int) *Dinic {
	return &Dinic{graph: make([][]edge, n), level: make([]int, n), iter: make([]int, n)}
}

func (d *Dinic) AddEdge(u, v, cap int) {
	d.graph[u] = append(d.graph[u], edge{to: v, cap: cap, rev: len(d.graph[v])})
	d.graph[v] = append(d.graph[v], edge{to: u, cap: 0, rev: len(d.graph[u]) - 1})
}

func (d *Dinic) bfs(s, t int) bool {
	for i := range d.level {
		d.level[i] = -1
	}
	d.level[s] = 0
	queue := []int{s}
	for i := 0; i < len(queue); i++ {
		u := queue[i]
		for _, e := range d.graph[u] {
			if e.cap > 0 && d.level[e.to] < 0 {
				d.level[e.to] = d.level[u] + 1
				queue = append(queue, e.to)
			}
		}
	}
	return d.level[t] >= 0
}

func (d *Dinic) dfs(u, t, pushed int) int {
	if u == t {
		return pushed
	}
	for ; d.iter[u] < len(d.graph[u]); d.iter[u]++ {
		e := &d.graph[u][d.iter[u]]
		if e.cap > 0 && d.level[e.to] == d.level[u]+1 {
			amt := pushed
			if e.cap < amt {
				amt = e.cap
			}
			if f := d.dfs(e.to, t, amt); f > 0 {
				e.cap -= f
				d.graph[e.to][e.rev].cap += f
				return f
			}
		}
	}
	return 0
}

func (d *Dinic) MaxFlow(s, t int) int {
	flow := 0
	for d.bfs(s, t) {
		for i := range d.iter {
			d.iter[i] = 0
		}
		for {
			f := d.dfs(s, t, math.MaxInt64)
			if f == 0 {
				break
			}
			flow += f
		}
	}
	return flow
}
