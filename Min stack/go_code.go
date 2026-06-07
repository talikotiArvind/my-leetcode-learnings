package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	minVal := val
	if len(ms.minStack) > 0 && ms.minStack[len(ms.minStack)-1] < val {
		minVal = ms.minStack[len(ms.minStack)-1]
	}
	ms.minStack = append(ms.minStack, minVal)
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.minStack = ms.minStack[:len(ms.minStack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}

func main() {
	ms := &MinStack{}
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Println(ms.GetMin()) // -3
	ms.Pop()
	fmt.Println(ms.Top())    // 0
	fmt.Println(ms.GetMin()) // -2
}
