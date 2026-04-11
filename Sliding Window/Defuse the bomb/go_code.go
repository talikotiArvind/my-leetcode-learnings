// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func decryptData(code [] int, k int) [] int {
    x := len(code)
    result := make([] int, x)
    
    if k == 0 {
        return result
    }
    start := 1
    absK := k
    if k < 0 {
        start = x + k
        absK = -k
    }
    
    end := start + absK
    
    windowSum := 0
    for i:= start; i < end; i++ {
        windowSum += code[i%x]
    }
    result[0] = windowSum
    
    for i:= 1; i < x; i++ {
        removed := code[(start + i - 1) % x]
        added:= code[(end + i - 1) % x]
        windowSum = windowSum + added - removed
        result[i] = windowSum
    }
    return result
    
}

func main() {
    data := [] int {5,7,1,4}
    k := 3
    output := decryptData(data, k)
    fmt.Println(output)
}
