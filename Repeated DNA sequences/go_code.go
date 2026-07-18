package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	m := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	seen := map[int]bool{}
	repeated := map[string]bool{}
	code, mask := 0, (1<<20)-1

	for i := 0; i < len(s); i++ {
		code = ((code << 2) | m[s[i]]) & mask // shift in new char, drop oldest
		if i >= 9 {
			if seen[code] {
				repeated[s[i-9:i+1]] = true
			} else {
				seen[code] = true
			}
		}
	}

	res := make([]string, 0, len(repeated))
	for k := range repeated {
		res = append(res, k)
	}
	return res
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Printf("Input:  %s\n", s)
	fmt.Printf("Output: %v\n", findRepeatedDnaSequences(s))
	// Output: [AAAAACCCCC CCCCCAAAAA]

	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA")) // [AAAAAAAAAA]
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAA"))    // [] (len == 10, no repeat)
	fmt.Println(findRepeatedDnaSequences("ACGT"))          // [] (too short)
}
