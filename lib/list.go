package lib

type List[T any] struct {
	value []T
}

func NewList[T any]() List[T] {
	return List[T]{
		value: make([]T, 0, 10),
	}
}

func (l *List[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *List[T]) GetValue() []T {
	return l.value
}

func (l *List[T]) Size() int {
	return len(l.value)
}

func (l *List[T]) InsertBack(value T) {
	l.value = append(l.value, value)
}

func (l *List[T]) RemoveFront() T {
	first := l.value[0]   //get the first element
	l.value = l.value[1:] //remove the first element

	return first
}

func (l *List[T]) RemoveLast() T {
	last := l.value[len(l.value)-1]
	l.value = l.value[:len(l.value)-1]

	return last
}

func (l *List[T]) Get(index int) T {
	return l.value[index]
}
