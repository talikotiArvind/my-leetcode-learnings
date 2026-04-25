package main
import "fmt"

func sortArrayByParity(nums [] int) [] int {
    i,j := 0,0
    for i < len(nums) {
        if nums[i] % 2 == 0 {
            nums[j], nums[i] = nums[i], nums[j]
            j++
        }
        i++
    }
    return nums
}

func main() {
	nums1 := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(nums1))
}
