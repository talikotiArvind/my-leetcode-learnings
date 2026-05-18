package main
import "fmt"

func fibonacci(n int) int {
    if n <= 1 {return n}
    x, y := 0, 1
    
    for i := 2; i<= n; i++ {
        x, y = y, x+y
    }
    return y
}

func main() {
    n := 6
    fmt.Println(fibonacci(n))
}
