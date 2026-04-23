package main
import "fmt"
import "strconv"

func summaryRanges(nums [] int) [] string {
    output := []string{}
    i := 0
    for (i < len(nums)) {
        start := nums[i]
        for i + 1 < len(nums) && nums[i + 1] == nums[i] + 1 {
            i++
        }
        if nums[i] == start {
            output = append(output, strconv.Itoa(start))
        } else {
            output = append(output, fmt.Sprintf("%d->%d", start, nums[i]))
        }
        i ++
    }
    return output
}

func main() {
  fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
