func ReservoirSample[T any](stream <-chan T, k int) []T {
	res := make([]T, 0, k)
	i := 0
	for item := range stream {
		if i < k {
			res = append(res, item)
		} else if j := rand.Intn(i + 1); j < k {
			res[j] = item
		}
		i++
	}
	return res
}