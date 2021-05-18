package optimization

import "container/heap"

// Evaluation queue copy-pasted from Go docs

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

// Item is an example item for pathfinding. Replace Node with your game data.
type Item struct {
	// Node is the value of the item, it is arbitrary data.
	*State

	// Priority is set by the user and used by the heap to order.
	Priority int

	// Index is used by the heap.Interface methods to keep things sorted.
	Index int
}

// Len returns the length of the priority queue
func (pq PriorityQueue) Len() int { return len(pq) }

// Less defines if this is a min-heap or a max-heap. For optimization we want max.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

// Swap swaps items at index i and j
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an item
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop removes the lowest priority item
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the Evaluation and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value *State, priority int) {
	item.State = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

// Empty returns true if the priority queue is empty
func (pq PriorityQueue) Empty() bool {
	return len(pq) == 0
}

// Clear empties the priority queue but retains the memory
func (pq *PriorityQueue) Clear() {
	*pq = (*pq)[0:0]
}
