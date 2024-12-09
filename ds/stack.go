package ds

type Stack[T any] []T

func NewStack[T any]() Stack[T] {
	return []T{}
}

func (stack *Stack[T]) Push(value T) {
	*stack = append(*stack, value)
}

func (stack *Stack[T]) Pop() *T {
	if stack.IsEmpty() {
		return nil
	}

	last := (*stack)[len(*stack)-1]
	*stack = (*stack)[0 : len(*stack)-1]

	return &last
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(*stack) == 0
}
