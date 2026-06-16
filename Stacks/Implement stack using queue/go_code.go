package main

import "fmt"

type MyStack struct {
	// Using a Go slice to act as a standard FIFO queue
	queue []int
}

// Constructor initializes a new stack object
func Constructor() MyStack {
	return MyStack{queue: make([]int, 0)}
}

// Push adds an element to the stack and rearranges the queue
func (this *MyStack) Push(x int) {
	// 1. Enqueue the new element to the back
	this.queue = append(this.queue, x)
	
	// 2. Rotate the queue so the newly added element sits at the front
	size := len(this.queue)
	for i := 0; i < size-1; i++ {
		// Dequeue from front
		front := this.queue[0]
		this.queue = this.queue[1:]
		// Enqueue to back
		this.queue = append(this.queue, front)
	}
}

// Pop removes the element on the top of the stack and returns it
func (this *MyStack) Pop() int {
	// The front of the queue is the top of our stack
	topElement := this.queue[0]
	this.queue = this.queue[1:]
	return topElement
}

// Top returns the element on the top of the stack without removing it
func (this *MyStack) Top() int {
	return this.queue[0]
}

// Empty returns whether the stack is empty
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func main() {
	// Initialize the stack
	stack := Constructor()
	fmt.Println("Initial stack empty?", stack.Empty()) // Expects: true

	// Push elements
	fmt.Println("\nPushing 10, then 20, then 30...")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Peek the top element
	fmt.Println("Top element (peek):", stack.Top()) // Expects: 30

	// Pop elements
	fmt.Println("Popped:", stack.Pop()) // Expects: 30
	fmt.Println("Popped:", stack.Pop()) // Expects: 20

	// Check final state
	fmt.Println("New top element:", stack.Top())      // Expects: 10
	fmt.Println("Is stack empty now?", stack.Empty()) // Expects: false
}
