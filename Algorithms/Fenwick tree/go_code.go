type Fenwick struct {
	n int
	t []int
}

func NewFenwick(n int) *Fenwick { return &Fenwick{n: n, t: make([]int, n+1)} }

func (f *Fenwick) Add(i, delta int) {
	for i++; i <= f.n; i += i & -i {
		f.t[i] += delta
	}
}

func (f *Fenwick) Prefix(i int) int {
	s := 0
	for ; i > 0; i -= i & -i {
		s += f.t[i]
	}
	return s
}

func (f *Fenwick) RangeSum(l, r int) int { return f.Prefix(r) - f.Prefix(l) }