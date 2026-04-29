package main
import "fmt"

func digitSum(n int) int {
	total := 0
	for n > 0 {
		d := n % 10
		total += d * d
		n /= 10
	}
	return total
}

func isHappyNum(n int) bool {
    found := map[int]bool {}
    for n != 1 {
        n = digitSum(n)
        if found[n] {
            return false
        }
        found[n] = true
    }
    return true
}

func main () {
    n := 2
    fmt.Println(isHappyNum(n))
}
