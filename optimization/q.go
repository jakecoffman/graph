package optimization

type Queue struct {
	queue []*State
}

// Put adds node to the end
func (q *Queue) Put(n *State) *Queue {
	q.queue = append(q.queue, n)
	return q
}

// Pop removes the first element and returns it
func (q *Queue) Pop() *State {
	val := q.queue[0]
	q.queue = q.queue[1:]
	return val
}

// Empty returns true if the queue is empty
func (q *Queue) Empty() bool {
	return len(q.queue) == 0
}

// Clear empties the queue but retains the memory
func (q *Queue) Clear() {
	q.queue = q.queue[0:0]
}
