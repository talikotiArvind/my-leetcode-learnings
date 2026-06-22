package main

import "fmt"

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	for i, x := range nums {
		prefix[i+1] = prefix[i] + x
	}
	return NumArray{prefix: prefix}
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.prefix[right+1] - na.prefix[left]
}

func main() {=
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2)) // 1   (-2 + 0 + 3)
	fmt.Println(obj.SumRange(2, 5)) // -1  (3 - 5 + 2 - 1)
	fmt.Println(obj.SumRange(0, 5)) // -3  (whole array)
	fmt.Println(obj.SumRange(3, 3)) // -5  (single element)

	obj2 := Constructor([]int{1, 2, 3, 4, 5})
	fmt.Println(obj2.SumRange(1, 3)) // 9   (2 + 3 + 4)
	fmt.Println(obj2.SumRange(0, 4)) // 15  (whole array)
}
