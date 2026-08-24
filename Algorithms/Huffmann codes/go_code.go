type hnode struct {
	ch          byte
	freq        int
	left, right *hnode
}

type hpq []*hnode

func (p hpq) Len() int            { return len(p) }
func (p hpq) Less(i, j int) bool  { return p[i].freq < p[j].freq }
func (p hpq) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *hpq) Push(x any)         { *p = append(*p, x.(*hnode)) }
func (p *hpq) Pop() any {
	old := *p
	item := old[len(old)-1]
	*p = old[:len(old)-1]
	return item
}

func HuffmanCodes(text string) map[byte]string {
	freq := map[byte]int{}
	for i := 0; i < len(text); i++ {
		freq[text[i]]++
	}
	if len(freq) == 0 {
		return nil
	}

	h := &hpq{}
	for ch, f := range freq {
		*h = append(*h, &hnode{ch: ch, freq: f})
	}
	heap.Init(h)
	if h.Len() == 1 {
		return map[byte]string{(*h)[0].ch: "0"}
	}
	for h.Len() > 1 {
		lo := heap.Pop(h).(*hnode)
		hi := heap.Pop(h).(*hnode)
		heap.Push(h, &hnode{freq: lo.freq + hi.freq, left: lo, right: hi})
	}

	codes := map[byte]string{}
	var walk func(n *hnode, prefix string)
	walk = func(n *hnode, prefix string) {
		if n.left == nil && n.right == nil {
			codes[n.ch] = prefix
			return
		}
		walk(n.left, prefix+"0")
		walk(n.right, prefix+"1")
	}
	walk(heap.Pop(h).(*hnode), "")
	return codes
}