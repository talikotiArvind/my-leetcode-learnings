package main
import "fmt"

func reverseBits(n  uint32) uint32 {
    var result uint32
    for i := 0; i < 32; i++ {
        result = (result << 1) | (n & 1)
        n >>= 1
    }
    return result
}

func main() {
  n := uint32(43261596)
  fmt.Printf("Input:  %032b (%d)\n", n, n)
  r := reverseBits(n)
  fmt.Printf("Output: %032b (%d)\n", r, r)
}
