package lib

type IList[T any] interface {
	IsEmpty() bool
	GetValue() []T
	Size() int
	InsertBack(value T)
	RemoveFront() T
	RemoveLast() T
	Get(index int) T
}

type list[T any] struct {
	value []T
}

func NewList[T any]() IList[T] {
	return &list[T]{
		value: make([]T, 0, 10),
	}
}

func (l *list[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *list[T]) GetValue() []T {
	return l.value
}

func (l *list[T]) Size() int {
	return len(l.value)
}

func (l *list[T]) InsertBack(value T) {
	l.value = append(l.value, value)
}

func (l *list[T]) RemoveFront() T {
	first := l.value[0]   //get the first element
	l.value = l.value[1:] //remove the first element

	return first
}

func (l *list[T]) RemoveLast() T {
	last := l.value[len(l.value)-1]
	l.value = l.value[:len(l.value)-1]

	return last
}

func (l *list[T]) Get(index int) T {
	return l.value[index]
}
