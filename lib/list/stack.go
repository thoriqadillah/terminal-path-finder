package list

type IStack[T any] interface {
	IsEmpty() bool
	Push(element T)
	Pop() T
}

func NewStack[T any]() IStack[T] {
	return &list[T]{}
}

func (s *list[T]) Push(element T) {
	s.InsertBack(element)
}

func (s *list[T]) Pop() T {
	return s.RemoveLast()
}
