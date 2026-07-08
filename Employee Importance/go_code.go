package main

import "fmt"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	empMap := make(map[int]*Employee)
	for _, emp := range employees {
		empMap[emp.Id] = emp
	}

	total := 0
	queue := []int{id}
	for len(queue) > 0 {
		currID := queue[0]
		queue = queue[1:]
		emp := empMap[currID]
		total += emp.Importance
		queue = append(queue, emp.Subordinates...)
	}
	return total
}

func main() {
	employees := []*Employee{
		{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
		{Id: 2, Importance: 3, Subordinates: []int{}},
		{Id: 3, Importance: 3, Subordinates: []int{}},
	}
	fmt.Println(getImportance(employees, 1)) // 11
}
