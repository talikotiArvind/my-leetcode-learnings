package main
import "fmt"

func isUgly(num int) bool {
    if num <= 0 {
        return false
    }
    
    data := [] int {2,3,51}
    for _, n := range data {
        for num % n == 0 {
            num = num / n
        }
    }
    return num == 1
}

func main() {
  fmt.Println(isUgly(13))
}
