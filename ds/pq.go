package ds

// Item is an entry in the PriorityQueue.
type Item[T any] struct {
	// State is the value of the item, it is arbitrary data.
	State T
	// Priority is set by the user and used by the heap to order.
	Priority int
}

// NewItem creates a new item with the given state and priority.
func NewItem[T any](state T, priority int) *Item[T] {
	return &Item[T]{
		State:    state,
		Priority: priority,
	}
}

// NewPriorityQueue creates a heap that is sorted by priority.
func NewPriorityQueue[T any](compare func(a, b int) bool) *Heap[*Item[T]] {
	return NewHeap(func(a, b *Item[T]) bool {
		return compare(a.Priority, b.Priority)
	})
}
