package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{}
	for _, part := range strings.Split(path, "/") {
		switch part {
		case "", ".":
			// skip empty segments and current-dir markers
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // pop: go up one level
			}
		default:
			stack = append(stack, part) // push a real directory name
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	fmt.Println(simplifyPath("/home/"))                // /home
	fmt.Println(simplifyPath("/../"))                  // /
	fmt.Println(simplifyPath("/home//foo/"))           // /home/foo
	fmt.Println(simplifyPath("/a/./b/../../c/"))       // /c
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))  // /c
}
