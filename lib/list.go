package lib

type list[T any] struct {
	Value []T
}

func newList[T any]() *list[T] {
	return &list[T]{
		Value: make([]T, 0, 10),
	}
}

func (l *list[T]) insertBack(value T) {
	l.Value = append(l.Value, value)
}

func (l *list[T]) removeFront() T {
	first := l.Value[0]   //get the first element
	l.Value = l.Value[1:] //remove the first element

	return first
}

func (l *list[T]) removeLast() T {
	last := l.Value[len(l.Value)-1]
	l.Value = l.Value[:len(l.Value)-1]

	return last
}
