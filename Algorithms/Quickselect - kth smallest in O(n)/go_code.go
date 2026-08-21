func QuickSelect(a []int, k int) int {
	lo, hi := 0, len(a)-1
	for lo < hi {
		pivot := a[lo+rand.Intn(hi-lo+1)]
		lt, i, gt := lo, lo, hi
		for i <= gt {
			switch {
			case a[i] < pivot:
				a[lt], a[i] = a[i], a[lt]
				lt++
				i++
			case a[i] > pivot:
				a[gt], a[i] = a[i], a[gt]
				gt--
			default:
				i++
			}
		}
		switch {
		case k < lt:
			hi = lt - 1
		case k > gt:
			lo = gt + 1
		default:
			return a[k]
		}
	}
	return a[lo]
}