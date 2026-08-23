type LazySegTree struct {
	n    int
	tree []int
	lazy []int
}

func NewLazySegTree(data []int) *LazySegTree {
	t := &LazySegTree{n: len(data), tree: make([]int, 4*len(data)), lazy: make([]int, 4*len(data))}
	t.build(1, 0, t.n, data)
	return t
}

func (t *LazySegTree) build(node, lo, hi int, data []int) {
	if hi-lo == 1 {
		t.tree[node] = data[lo]
		return
	}
	mid := (lo + hi) / 2
	t.build(2*node, lo, mid, data)
	t.build(2*node+1, mid, hi, data)
	t.tree[node] = t.tree[2*node] + t.tree[2*node+1]
}

func (t *LazySegTree) apply(node, lo, hi, val int) {
	t.tree[node] += val * (hi - lo)
	t.lazy[node] += val
}

func (t *LazySegTree) push(node, lo, hi int) {
	if t.lazy[node] != 0 {
		mid := (lo + hi) / 2
		t.apply(2*node, lo, mid, t.lazy[node])
		t.apply(2*node+1, mid, hi, t.lazy[node])
		t.lazy[node] = 0
	}
}

func (t *LazySegTree) Add(l, r, val int) { t.add(l, r, val, 1, 0, t.n) }

func (t *LazySegTree) add(l, r, val, node, lo, hi int) {
	if r <= lo || hi <= l {
		return
	}
	if l <= lo && hi <= r {
		t.apply(node, lo, hi, val)
		return
	}
	t.push(node, lo, hi)
	mid := (lo + hi) / 2
	t.add(l, r, val, 2*node, lo, mid)
	t.add(l, r, val, 2*node+1, mid, hi)
	t.tree[node] = t.tree[2*node] + t.tree[2*node+1]
}

func (t *LazySegTree) Query(l, r int) int { return t.query(l, r, 1, 0, t.n) }

func (t *LazySegTree) query(l, r, node, lo, hi int) int {
	if r <= lo || hi <= l {
		return 0
	}
	if l <= lo && hi <= r {
		return t.tree[node]
	}
	t.push(node, lo, hi)
	mid := (lo + hi) / 2
	return t.query(l, r, 2*node, lo, mid) + t.query(l, r, 2*node+1, mid, hi)
}
