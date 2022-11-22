package ds

// Stack is First In Last Out
type Stack[T any] []T

func (s *Stack[T]) Push(node T) {
	*s = append(*s, node)
}

func (s *Stack[T]) Pop() T {
	node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return node
}

func (s *Stack[T]) Len() int {
	return len(*s)
}
