
package main
import "fmt"

func addDigits(num int) int {
    if num == 0 {
        return 0
        
    }else if num % 9 == 0 {
        return 9
        
    }else{
        return num % 9
    }
}

func main () {
    num := 1132
    fmt.Println(addDigits(num))
}
