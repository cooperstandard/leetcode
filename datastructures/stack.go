package datastructures

type Stack[T any] struct {
	Items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil} //stack holding a nil slice
}

func (s *Stack[T]) Push(e T) {
	s.Items = append(s.Items, e)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		return *new(T)
	}

	return s.Items[len(s.Items)-1]

}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		return *new(T)
	}

	top := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return top

}
