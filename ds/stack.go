package ds

type Stack []int

func (stack *Stack) Push(value int) {
	*stack = append(*stack, value)
}

func (stack *Stack) Pop() *int {
	if stack.IsEmpty() {
		return nil
	}

	last := (*stack)[len(*stack)-1]
	*stack = (*stack)[0 : len(*stack)-1]

	return &last
}

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}
