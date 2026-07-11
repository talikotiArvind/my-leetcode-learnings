func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return [][]string{}
	}

	adj := make(map[string][]string)
	steps := make(map[string]int)
	steps[beginWord] = 1

	queue := []string{beginWord}
	found := false

	for len(queue) > 0 && !found {
		size := len(queue)
		currLevelVisited := make(map[string]bool)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			wordBytes := []byte(curr)
			for j := 0; j < len(wordBytes); j++ {
				old := wordBytes[j]
				for c := 'a'; c <= 'z'; c++ {
					wordBytes[j] = byte(c)
					nextWord := string(wordBytes)

					if wordSet[nextWord] {
						if step, exists := steps[nextWord]; !exists || step == steps[curr]+1 {
							adj[curr] = append(adj[curr], nextWord)
							if !exists {
								steps[nextWord] = steps[curr] + 1
								currLevelVisited[nextWord] = true
								queue = append(queue, nextWord)
							}
						}
						if nextWord == endWord {
							found = true
						}
					}
				}
				wordBytes[j] = old
			}
		}
	}

	var results [][]string
	if found {
		var path []string = []string{beginWord}
		var dfs func(word string)
		dfs = func(word string) {
			if word == endWord {
				temp := make([]string, len(path))
				copy(temp, path)
				results = append(results, temp)
				return
			}
			for _, nxt := range adj[word] {
				if steps[nxt] == steps[word]+1 {
					path = append(path, nxt)
					dfs(nxt)
					path = path[:len(path)-1]
				}
			}
		}
		dfs(beginWord)
	}
	return results
}
