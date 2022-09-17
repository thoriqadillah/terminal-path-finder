package lib

type iQueue[T any] interface {
	IsEmpty() bool
	Enqueue(element T)
	Dequeue() T
}

type queue[T any] struct {
	*list[T]
}

func NewQueue[T any]() iQueue[T] {
	return &queue[T]{
		newList[T](),
	}
}

func (q *queue[T]) IsEmpty() bool {
	return len(q.Value) == 0
}

func (q *queue[T]) Enqueue(element T) {
	q.list.insertBack(element)
}

func (q queue[T]) Dequeue() T {
	return q.list.removeFront()
}
