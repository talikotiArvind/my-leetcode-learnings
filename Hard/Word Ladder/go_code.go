func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return 0
	}

	queue := []string{beginWord}
	visited := map[string]bool{beginWord: true}
	level := 1

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == endWord {
				return level
			}

			wordBytes := []byte(curr)
			for j := 0; j < len(wordBytes); j++ {
				originalByte := wordBytes[j]
				for c := 'a'; c <= 'z'; c++ {
					wordBytes[j] = byte(c)
					nextWord := string(wordBytes)

					if wordSet[nextWord] && !visited[nextWord] {
						visited[nextWord] = true
						queue = append(queue, nextWord)
					}
				}
				wordBytes[j] = originalByte
			}
		}
		level++
	}
	return 0
}
