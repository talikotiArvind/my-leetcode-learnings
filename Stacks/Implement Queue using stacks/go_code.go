package main

import "fmt"

type MyQueue struct {
	inputStack  []int
	outputStack []int
}

// Constructor initializes a new queue object
func Constructor() MyQueue {
	return MyQueue{
		inputStack:  make([]int, 0),
		outputStack: make([]int, 0),
	}
}

// Push adds an element to the back of the queue
func (this *MyQueue) Push(x int) {
	this.inputStack = append(this.inputStack, x)
}

// moveElements shifts items from input to output stack if output is empty
func (this *MyQueue) moveElements() {
	if len(this.outputStack) == 0 {
		for len(this.inputStack) > 0 {
			// Pop from inputStack
			lastIdx := len(this.inputStack) - 1
			topItem := this.inputStack[lastIdx]
			this.inputStack = this.inputStack[:lastIdx]

			// Push to outputStack
			this.outputStack = append(this.outputStack, topItem)
		}
	}
}

// Pop removes the element from the front of the queue and returns it
func (this *MyQueue) Pop() int {
	this.moveElements()
	
	// Pop from outputStack
	lastIdx := len(this.outputStack) - 1
	frontItem := this.outputStack[lastIdx]
	this.outputStack = this.outputStack[:lastIdx]
	
	return frontItem
}

// Peek returns the element at the front of the queue without removing it
func (this *MyQueue) Peek() int {
	this.moveElements()
	return this.outputStack[len(this.outputStack)-1]
}

// Empty returns whether the queue is empty
func (this *MyQueue) Empty() bool {
	return len(this.inputStack) == 0 && len(this.outputStack) == 0
}

func main() {
	// Initialize the queue
	queue := Constructor()
	fmt.Println("Initial queue empty?", queue.Empty()) // Expects: true

	// Push elements
	fmt.Println("\nPushing 1, then 2, then 3...")
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	// Peek the front element
	fmt.Println("Front element (peek):", queue.Peek()) // Expects: 1

	// Pop elements
	fmt.Println("Popped:", queue.Pop()) // Expects: 1
	fmt.Println("Popped:", queue.Pop()) // Expects: 2

	// Push another element to show mixed operations work seamlessly
	fmt.Println("Pushing 4...")
	queue.Push(4)

	// Check remaining items
	fmt.Println("Front element (peek):", queue.Peek()) // Expects: 3
	fmt.Println("Popped:", queue.Pop())                // Expects: 3
	fmt.Println("Popped:", queue.Pop())                // Expects: 4
	fmt.Println("Is queue empty now?", queue.Empty())  // Expects: true
}
