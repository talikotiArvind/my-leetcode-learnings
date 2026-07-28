import "strings"

func maxRepeating(sequence string, word string) int {
	k := 0
	repeatedWord := word
	for strings.Contains(sequence, repeatedWord) {
		k++
		repeatedWord += word
	}
	
	return k
}
