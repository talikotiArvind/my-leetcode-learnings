package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
	i := 0 // index into name
	for j := 0; j < len(typed); j++ {
		if i < len(name) && name[i] == typed[j] {
			i++
		} else if j == 0 || typed[j-1] != typed[j] {
			return false // extra char that isn't a repeat -> invalid
		}
	}
	return i == len(name)
}

func main() {
	tests := []struct {
		name  string
		typed string
	}{
		{"alex", "aaleex"},   // true
		{"saeed", "ssaaedd"}, // false
		{"leelee", "lleeelee"}, // true
		{"laiden", "laiden"},   // true
	}

	for _, t := range tests {
		fmt.Printf("name=%-8q typed=%-10q => %v\n", t.name, t.typed, isLongPressedName(t.name, t.typed))
	}
}
