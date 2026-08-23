func RabinKarp(text, pat string) []int {
	const base, mod = 256, 1_000_000_007
	n, m := len(text), len(pat)
	if m == 0 || m > n {
		return nil
	}
	var high uint64 = 1
	for i := 0; i < m-1; i++ {
		high = high * base % mod
	}
	var ph, th uint64
	for i := 0; i < m; i++ {
		ph = (ph*base + uint64(pat[i])) % mod
		th = (th*base + uint64(text[i])) % mod
	}

	var res []int
	for s := 0; ; s++ {
		if ph == th && text[s:s+m] == pat {
			res = append(res, s)
		}
		if s == n-m {
			break
		}
		th = (th + mod - uint64(text[s])*high%mod) % mod   // drop leading char
		th = (th*base + uint64(text[s+m])) % mod            // append trailing
	}
	return res
}
