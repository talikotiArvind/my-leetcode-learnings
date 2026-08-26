func lastRow(a, b string) []int {
	prev := make([]int, len(b)+1)
	cur := make([]int, len(b)+1)
	for i := 0; i < len(a); i++ {
		cur[0] = 0
		for j := 1; j <= len(b); j++ {
			switch {
			case a[i] == b[j-1]:
				cur[j] = prev[j-1] + 1
			case prev[j] >= cur[j-1]:
				cur[j] = prev[j]
			default:
				cur[j] = cur[j-1]
			}
		}
		prev, cur = cur, prev
	}
	return prev
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func Hirschberg(a, b string) string {
	switch {
	case len(a) == 0 || len(b) == 0:
		return ""
	case len(a) == 1:
		if strings.IndexByte(b, a[0]) >= 0 {
			return a
		}
		return ""
	}

	mid := len(a) / 2
	left := lastRow(a[:mid], b)
	right := lastRow(reverse(a[mid:]), reverse(b))

	best, k := -1, 0
	for j := 0; j <= len(b); j++ {
		if s := left[j] + right[len(b)-j]; s > best {
			best, k = s, j
		}
	}
	return Hirschberg(a[:mid], b[:k]) + Hirschberg(a[mid:], b[k:])
}