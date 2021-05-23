package pathfinding

type Queue []*Node

// Put adds node to the end
func (q *Queue) Put(n *Node) *Queue {
	*q = append(*q, n)
	return q
}

// Pop removes the first element and returns it
func (q *Queue) Pop() *Node {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

// Empty returns true if the queue is empty
func (q *Queue) Empty() bool {
	return len(*q) == 0
}

// Clear empties the queue but retains the memory
func (q *Queue) Clear() {
	*q = (*q)[0:0]
}
