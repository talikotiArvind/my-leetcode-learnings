func commonChars(words []string) []string {
    minFreq := [26]int{}

    // Init with first word's frequency
    for _, ch := range words[0] {
        minFreq[ch-'a']++
    }

    for _, word := range words[1:] {
        curr := [26]int{}
        for _, ch := range word {
            curr[ch-'a']++
        }
        for i := 0; i < 26; i++ {
            if curr[i] < minFreq[i] {
                minFreq[i] = curr[i]
            }
        }
    }

    result := []string{}
    for i, count := range minFreq {
        for j := 0; j < count; j++ {
            result = append(result, string(rune('a'+i)))
        }
    }
    return result
}
