type AhoCorasick struct {
	next []map[byte]int
	fail []int
	out  [][]int
}

func NewAhoCorasick(patterns []string) *AhoCorasick {
	ac := &AhoCorasick{next: []map[byte]int{{}}, fail: []int{0}, out: [][]int{nil}}

	for idx, p := range patterns {
		node := 0
		for i := 0; i < len(p); i++ {
			ch := p[i]
			nxt, ok := ac.next[node][ch]
			if !ok {
				nxt = len(ac.next)
				ac.next = append(ac.next, map[byte]int{})
				ac.fail = append(ac.fail, 0)
				ac.out = append(ac.out, nil)
				ac.next[node][ch] = nxt
			}
			node = nxt
		}
		ac.out[node] = append(ac.out[node], idx)
	}
	ac.build()
	return ac
}

func (ac *AhoCorasick) build() {
	queue := make([]int, 0, len(ac.next))
	for _, nxt := range ac.next[0] {
		queue = append(queue, nxt)
	}
	for i := 0; i < len(queue); i++ {   // BFS: queue grows as we walk it
		v := queue[i]
		ac.out[v] = append(ac.out[v], ac.out[ac.fail[v]]...)
		for ch, nxt := range ac.next[v] {
			f := ac.fail[v]
			for f != 0 {
				if _, ok := ac.next[f][ch]; ok {
					break
				}
				f = ac.fail[f]
			}
			if t, ok := ac.next[f][ch]; ok {
				ac.fail[nxt] = t
			}
			queue = append(queue, nxt)
		}
	}
}

// Search returns {endIndex, patternIndex} pairs for every occurrence.
func (ac *AhoCorasick) Search(text string) [][2]int {
	node := 0
	var hits [][2]int
	for i := 0; i < len(text); i++ {
		ch := text[i]
		for node != 0 {
			if _, ok := ac.next[node][ch]; ok {
				break
			}
			node = ac.fail[node]
		}
		if t, ok := ac.next[node][ch]; ok {
			node = t
		} else {
			node = 0
		}
		for _, idx := range ac.out[node] {
			hits = append(hits, [2]int{i, idx})
		}
	}
	return hits
}