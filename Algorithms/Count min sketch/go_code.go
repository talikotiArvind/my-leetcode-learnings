type CountMinSketch struct {
	table [][]uint64
	w, d  int
	Total uint64
}

func NewCountMinSketch(epsilon, delta float64) *CountMinSketch {
	w := int(math.Ceil(math.E / epsilon))
	d := int(math.Ceil(math.Log(1 / delta)))
	t := make([][]uint64, d)
	for i := range t {
		t[i] = make([]uint64, w)
	}
	return &CountMinSketch{table: t, w: w, d: d}
}

func (c *CountMinSketch) pair(item []byte) (uint64, uint64) {
	sum := sha256.Sum256(item)
	return binary.BigEndian.Uint64(sum[0:8]), binary.BigEndian.Uint64(sum[8:16]) | 1
}

func (c *CountMinSketch) Add(item []byte, count uint64) {
	h1, h2 := c.pair(item)
	c.Total += count
	for i := 0; i < c.d; i++ {
		c.table[i][(h1+uint64(i)*h2)%uint64(c.w)] += count
	}
}

func (c *CountMinSketch) Estimate(item []byte) uint64 {
	h1, h2 := c.pair(item)
	min := ^uint64(0)
	for i := 0; i < c.d; i++ {
		if v := c.table[i][(h1+uint64(i)*h2)%uint64(c.w)]; v < min {
			min = v
		}
	}
	return min
}