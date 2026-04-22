package main
import "fmt"
import "strconv"

func fizzBuzz(n int) [] string {
    result := []string {}
    var value string
    for i := 1; i <= n; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            value = "FizzBuzz"
        } else if i % 3 == 0 {
            value = "Fizz"
        } else if i % 5 == 0 {
            value = "Buzz"
        } else {
            value = strconv.Itoa(i)
        }
        result = append(result, value)
    }
    return result
}

func main() {
    n := 10
    fmt.Println(fizzBuzz(n))
}
