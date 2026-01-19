package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	data := len(nums)
	j := 1

	for i := 1; i < data; i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j

}

func main() {
	// Example usage
	nums := []int{1, 1, 2, 2, 3}
	newLength := removeDuplicates(nums)
	println("New length:", newLength)
}
