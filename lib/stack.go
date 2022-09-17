package lib

type iStack[T any] interface {
	IsEmpty() bool
	Push(element T)
	Pop() T
}

type stack[T any] struct {
	*list[T]
}

func NewStack[T any]() iStack[T] {
	return &stack[T]{
		newList[T](),
	}
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.Value) == 0
}

func (s *stack[T]) Push(element T) {
	s.list.insertBack(element)
}

func (s *stack[T]) Pop() T {
	return s.list.removeLast()
}
