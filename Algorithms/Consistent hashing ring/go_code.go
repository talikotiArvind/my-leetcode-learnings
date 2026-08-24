type HashRing struct {
	vnodes int
	ring   map[uint64]string
	keys   []uint64 // kept sorted
}

func NewHashRing(vnodes int, nodes ...string) *HashRing {
	r := &HashRing{vnodes: vnodes, ring: map[uint64]string{}}
	for _, n := range nodes {
		r.Add(n)
	}
	return r
}

func hashKey(key string) uint64 {
	sum := sha256.Sum256([]byte(key))
	return binary.BigEndian.Uint64(sum[:8])
}

func (r *HashRing) search(h uint64) int {
	return sort.Search(len(r.keys), func(i int) bool { return r.keys[i] >= h })
}

func (r *HashRing) Add(node string) {
	for i := 0; i < r.vnodes; i++ {
		h := hashKey(fmt.Sprintf("%s#%d", node, i))
		r.ring[h] = node
		idx := r.search(h)
		r.keys = append(r.keys, 0)
		copy(r.keys[idx+1:], r.keys[idx:])
		r.keys[idx] = h
	}
}

func (r *HashRing) Remove(node string) {
	for i := 0; i < r.vnodes; i++ {
		h := hashKey(fmt.Sprintf("%s#%d", node, i))
		delete(r.ring, h)
		if idx := r.search(h); idx < len(r.keys) && r.keys[idx] == h {
			r.keys = append(r.keys[:idx], r.keys[idx+1:]...)
		}
	}
}

func (r *HashRing) Get(key string) (string, bool) {
	if len(r.keys) == 0 {
		return "", false
	}
	idx := r.search(hashKey(key))
	if idx == len(r.keys) {
		idx = 0 // wrap past the end of the ring
	}
	return r.ring[r.keys[idx]], true
}