package datastructures

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil} //stack holding a nil slice
}

func (s *Stack[T]) Push(e T) {
	s.items = append(s.items, e)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}

	return s.items[len(s.items)-1], true

}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true

}
