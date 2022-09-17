package lib

type IQueue[T any] interface {
	IsEmpty() bool
	Enqueue(element T)
	Dequeue() T
	Length() int
}

type queue[T any] struct {
	*list[T]
}

func NewQueue[T any]() IQueue[T] {
	return &queue[T]{
		newList[T](),
	}
}

func (q *queue[T]) Length() int {
	return len(q.Value)
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
