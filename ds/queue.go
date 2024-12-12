package ds

type Queue[T any] []T

func NewQueue[T any]() Queue[T] {
	return make([]T, 0)
}

func (queue *Queue[T]) Enqueue(element T) {
	*queue = append(*queue, element)
}

func (queue *Queue[T]) Peek() T {
	return (*queue)[0]
}

func (queue *Queue[T]) Dequeue() T {
	value := queue.Peek()
	*queue = (*queue)[1:]
	return value
}

func (queue *Queue[T]) IsEmpty() bool {
	return len(*queue) == 0
}

func (queue *Queue[T]) Size() int {
	return len(*queue)
}
