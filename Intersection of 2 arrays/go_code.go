package main
import "fmt"


func toSet(array [] int) map[int]struct{} {
    output := make(map[int]struct{}, len(array))
    for _, n := range array {
        output[n] = struct{}{}
    }
    return output
}

func Intersect2Arrays(nums1 [] int, nums2 [] int) map[int] struct{} {
    var p, q [] int
    output := make(map[int]struct {})
    if len(nums1) > len(nums2) {
        p, q = nums1, nums2
    } else {
        p, q = nums2, nums1
    }
    setP := toSet(p)
    setQ := toSet(q)
    for i := range setP {
        if _, ok := setQ[i]; ok {
            output[i] = struct{}{}
        }
    }
    return output
}

func main() {
  nums1 := []int{1, 2, 2, 1, 3}
  nums2 := []int{2, 2, 3}
  fmt.Println(Intersect2Arrays(nums1, nums2))
}
