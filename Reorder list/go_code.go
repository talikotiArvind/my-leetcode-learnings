func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var nodes []*ListNode
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j] // append first
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i] // append last
		j--
	}
	nodes[i].Next = nil
}

func main() {
	l1 := build([]int{1, 2, 3, 4})
	reorderList(l1)
	fmt.Println(show(l1)) // 1->4->2->3

	l2 := build([]int{1, 2, 3, 4, 5})
	reorderList(l2)
	fmt.Println(show(l2)) // 1->5->2->4->3
}
