package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int
			switch token {
			case "+": res = a + b
			case "-": res = a - b
			case "*": res = a * b
			case "/": res = a / b
			}
			stack = append(stack, res)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func main() {
	tests := [][]string{
		{"2", "1", "+", "3", "*"},
		{"4", "13", "5", "/", "+"},
		{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
	}
	expected := []int{9, 6, 22}
	for i, t := range tests {
		fmt.Printf("%v\n→ %d  (expected %d)\n\n", t, evalRPN(t), expected[i])
	}
}
