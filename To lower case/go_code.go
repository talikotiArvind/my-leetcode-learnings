package main

import (
	"fmt"
	"strings"
)

func main() {
	// Sample mixed-case string
	originalText := "GoLang iS AwEsOmE! 123"

	// Convert the string to lower case
	lowercaseText := strings.ToLower(originalText)

	// Display the results
	fmt.Printf("Original: %s\n", originalText)
	fmt.Printf("Lowercase: %s\n", lowercaseText)
}
