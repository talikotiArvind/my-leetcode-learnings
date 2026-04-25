func intersect(nums1 []int, nums2 []int) [] int {
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }
    
    count := make(map[int]int)
    for _, n := range(nums1) {
        count[n]++
    }
    
    result := []int{}
    
    for _, element := range nums1 {
        if count[element] > 0 {
            result = append(result, element)
            count[element]--
        }
    }
    return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2)) // [2 2]
}
