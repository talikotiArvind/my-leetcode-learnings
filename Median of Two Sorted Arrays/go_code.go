package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	lo, hi := 0, m
	half := (m + n + 1) / 2

	for lo <= hi {
		i := (lo + hi) / 2
		j := half - i

		left1, right1 := math.Inf(-1), math.Inf(1)
		if i > 0 {
			left1 = float64(nums1[i-1])
		}
		if i < m {
			right1 = float64(nums1[i])
		}

		left2, right2 := math.Inf(-1), math.Inf(1)
		if j > 0 {
			left2 = float64(nums2[j-1])
		}
		if j < n {
			right2 = float64(nums2[j])
		}

		if left1 <= right2 && left2 <= right1 {
			if (m+n)%2 == 0 {
				return (math.Max(left1, left2) + math.Min(right1, right2)) / 2
			}
			return math.Max(left1, left2)
		} else if left1 > right2 {
			hi = i - 1
		} else {
			lo = i + 1
		}
	}
	panic("input arrays are not sorted correctly")
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    // 2
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.5
}
