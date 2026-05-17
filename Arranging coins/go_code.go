package main
import "fmt"

func arrangingCoins(n int) int {
    low, high := 0, n
    coinsNeeded := 0
    
    for low <= high {
        mid := (low + high) / 2
        coinsNeeded = (mid)*(mid + 1)/2
        if coinsNeeded == n {
            return mid
        } else if coinsNeeded < n {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return high
}

func main () {
    n := 11
    fmt.Println(arrangingCoins(n))
}
