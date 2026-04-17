package main
import "fmt"

func waterContainer(nums [] int) int {
    left, right := 0, len(nums) - 1
    water, maxWater := 0, 0
    for left < right {
        width:= right - left
        water = width * min(nums[left], nums[right])
        maxWater = max(maxWater, water)
        
        if nums[left] < nums[right] {
            left += 1
        } else {
            right -= 1
        }
    }
    return maxWater
}

func main() {
    nums := [] int {1,8,6,2,5,4,8,3,7}
    fmt.Println(waterContainer(nums))
}
