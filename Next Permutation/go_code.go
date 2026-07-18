func nextPermutation(nums []int) {
    n := len(nums)

    // 1. Find first index from the right where nums[i] < nums[i+1]
    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }

    // 2. If such a pivot exists, find the rightmost element > nums[i] and swap
    if i >= 0 {
        j := n - 1
        for nums[j] <= nums[i] {
            j--
        }
        nums[i], nums[j] = nums[j], nums[i]
    }

    // 3. Reverse the suffix after the pivot
    reverse(nums, i+1, n-1)
}

func reverse(nums []int, l, r int) {
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}
