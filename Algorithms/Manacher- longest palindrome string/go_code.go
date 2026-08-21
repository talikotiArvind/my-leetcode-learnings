func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	t := make([]byte, 0, 2*len(s)+1)
	t = append(t, '#')
	for i := 0; i < len(s); i++ {
		t = append(t, s[i], '#')
	}

	n := len(t)
	p := make([]int, n)
	center, right, best := 0, 0, 0
	for i := 0; i < n; i++ {
		if i < right {
			if mirror := p[2*center-i]; mirror < right-i {
				p[i] = mirror
			} else {
				p[i] = right - i
			}
		}
		for i-p[i]-1 >= 0 && i+p[i]+1 < n && t[i-p[i]-1] == t[i+p[i]+1] {
			p[i]++
		}
		if i+p[i] > right {
			center, right = i, i+p[i]
		}
		if p[i] > p[best] {
			best = i
		}
	}
	start := (best - p[best]) / 2
	return s[start : start+p[best]]
}