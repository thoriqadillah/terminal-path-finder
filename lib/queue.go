package lib

type IQueue[T any] interface {
	IsEmpty() bool
	Enqueue(element T)
	Dequeue() T
}

func NewQueue[T any]() IQueue[T] {
	return &list[T]{}
}

func (q *list[T]) Enqueue(element T) {
	q.InsertBack(element)
}

func (q *list[T]) Dequeue() T {
	return q.RemoveFront()
}
