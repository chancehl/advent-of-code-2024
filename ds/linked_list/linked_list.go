package linked_list

type ListNode struct {
	Value int
	Next  *ListNode
}

func FromSlice(values []int) *ListNode {
	if len(values) == 0 {
		return nil // Return nil if no values are provided
	}

	head := &ListNode{Value: values[0]}
	current := head

	for _, value := range values[1:] {
		current.Next = &ListNode{Value: value}
		current = current.Next
	}

	return head
}

func (ln *ListNode) ToSlice() []int {
	slice := make([]int, 0)
	current := ln

	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}

	return slice
}
