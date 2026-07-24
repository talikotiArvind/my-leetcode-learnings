package main

import "fmt"

func shortestCompletingWord(licensePlate string, words []string) string {
	need := count(licensePlate)

	best := ""
	for _, w := range words {
		have := count(w)
		ok := true
		for i := 0; i < 26; i++ {
			if have[i] < need[i] {
				ok = false
				break
			}
		}
		if ok && (best == "" || len(w) < len(best)) {
			best = w
		}
	}
	return best
}

// count returns letter frequencies (a-z), ignoring digits, spaces and case.
func count(s string) [26]int {
	var freq [26]int
	for _, ch := range s {
		switch {
		case ch >= 'a' && ch <= 'z':
			freq[ch-'a']++
		case ch >= 'A' && ch <= 'Z':
			freq[ch-'A']++
		}
	}
	return freq
}

func main() {
	fmt.Println(shortestCompletingWord(
		"1s3 PSt",
		[]string{"step", "steps", "stripe", "stepple"},
	)) // steps

	fmt.Println(shortestCompletingWord(
		"1s3 456",
		[]string{"looks", "pest", "stew", "show"},
	)) // pest

	fmt.Println(shortestCompletingWord(
		"Ah71752",
		[]string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"},
	)) // husband
}
