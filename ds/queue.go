package ds

type Queue[T any] []T

// Put adds node to the end
func (q *Queue[T]) Put(n T) *Queue[T] {
	*q = append(*q, n)
	return q
}

// Pop removes the first element and returns it
func (q *Queue[T]) Pop() T {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

// Empty returns true if the queue is empty
func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

// Clear empties the queue but retains the memory
func (q *Queue[T]) Clear() {
	*q = (*q)[0:0]
}
