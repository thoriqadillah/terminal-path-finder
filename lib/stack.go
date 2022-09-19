package lib

type IStack[T any] interface {
	IsEmpty() bool
	Push(element T)
	Pop() T
}

func NewStack[T any]() IStack[T] {
	return &List[T]{}
}

func (s *List[T]) Push(element T) {
	s.InsertBack(element)
}

func (s *List[T]) Pop() T {
	return s.RemoveLast()
}
