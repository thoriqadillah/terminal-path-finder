package lib

type IQueue[T any] interface {
	IsEmpty() bool
	Enqueue(element T)
	Dequeue() T
}

func NewQueue[T any]() IQueue[T] {
	return &List[T]{}
}

func (q *List[T]) Enqueue(element T) {
	q.InsertBack(element)
}

func (q *List[T]) Dequeue() T {
	return q.RemoveFront()
}
