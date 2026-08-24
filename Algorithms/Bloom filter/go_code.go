type BloomFilter struct {
	bits []byte
	m    uint64
	k    int
}

func NewBloomFilter(n int, fpRate float64) *BloomFilter {
	m := uint64(math.Ceil(-float64(n) * math.Log(fpRate) / (math.Ln2 * math.Ln2)))
	if m < 8 {
		m = 8
	}
	k := int(math.Round(float64(m) / float64(n) * math.Ln2))
	if k < 1 {
		k = 1
	}
	return &BloomFilter{bits: make([]byte, (m+7)/8), m: m, k: k}
}

func (b *BloomFilter) pair(item []byte) (uint64, uint64) {
	sum := sha256.Sum256(item)
	return binary.BigEndian.Uint64(sum[0:8]), binary.BigEndian.Uint64(sum[8:16]) | 1
}

func (b *BloomFilter) Add(item []byte) {
	h1, h2 := b.pair(item)
	for i := 0; i < b.k; i++ {
		idx := (h1 + uint64(i)*h2) % b.m
		b.bits[idx>>3] |= 1 << (idx & 7)
	}
}

func (b *BloomFilter) Contains(item []byte) bool {
	h1, h2 := b.pair(item)
	for i := 0; i < b.k; i++ {
		idx := (h1 + uint64(i)*h2) % b.m
		if b.bits[idx>>3]&(1<<(idx&7)) == 0 {
			return false
		}
	}
	return true
}