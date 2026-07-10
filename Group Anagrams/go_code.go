package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		chars := []byte(s)
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
		key := string(chars)
		groups[key] = append(groups[key], s)
	}

	result := [][]string{}
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// [[eat tea ate] [tan nat] [bat]]
	fmt.Println(groupAnagrams([]string{""}))  // [[]]
	fmt.Println(groupAnagrams([]string{"a"})) // [[a]]
}
