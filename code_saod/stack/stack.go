package stack

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {

	return &Stack[T]{}
}

func (s *Stack[T]) GetSize() int {

	return len(s.data)
}

func (s *Stack[T]) Push(val T) *Stack[T] {

	s.data = append(s.data, val)

	return s
}

func (s *Stack[T]) Pop() *T {

	if s.GetSize() == 0 {
		return nil
	}

	lastElementIdx := s.GetSize() - 1

	result := s.data[lastElementIdx]
	s.data = s.data[:lastElementIdx]

	return &result
}
