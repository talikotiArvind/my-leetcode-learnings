package main

import (
    "fmt"
    "sort"
)

func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    child, cookie := 0, 0

    for child < len(g) && cookie < len(s) {
        if s[cookie] >= g[child] {
            child++
        }
        cookie++
    }

    return child
}

func main() {
    fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))     // 1
    fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))     // 2
    fmt.Println(findContentChildren([]int{10, 9, 8}, []int{5, 6, 7})) // 0
}
