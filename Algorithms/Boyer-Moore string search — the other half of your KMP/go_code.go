func BoyerMoore(text, pat string) []int {
	n, m := len(text), len(pat)
	if m == 0 || m > n {
		return nil
	}
	var last [256]int
	for i := range last {
		last[i] = -1
	}
	for i := 0; i < m; i++ {
		last[pat[i]] = i
	}
	shift := goodSuffix(pat)

	var res []int
	for s := 0; s <= n-m; {
		j := m - 1
		for j >= 0 && pat[j] == text[s+j] {
			j--
		}
		if j < 0 {
			res = append(res, s)
			s += shift[0]
			continue
		}
		step := j - last[text[s+j]]
		if shift[j+1] > step {
			step = shift[j+1]
		}
		if step < 1 {
			step = 1
		}
		s += step
	}
	return res
}

func goodSuffix(pat string) []int {
	m := len(pat)
	shift := make([]int, m+1)
	border := make([]int, m+2)
	i, j := m, m+1
	border[i] = j
	for i > 0 {
		for j <= m && pat[i-1] != pat[j-1] {
			if shift[j] == 0 {
				shift[j] = j - i
			}
			j = border[j]
		}
		i--
		j--
		border[i] = j
	}
	j = border[0]
	for i := 0; i <= m; i++ {
		if shift[i] == 0 {
			shift[i] = j
		}
		if i == j {
			j = border[j]
		}
	}
	return shift
}