package optimization

import "github.com/jakecoffman/graph/ds"

// Item is an example item for pathfinding. Replace Node with your game data.
type Item struct {
	// Node is the value of the item, it is arbitrary data.
	*State

	// Priority is set by the user and used by the heap to order.
	Priority int

	// Index is used by the heap.Interface methods to keep things sorted.
	Index int
}

func NewPriorityQueue() *ds.Heap[*Item] {
	return ds.NewHeap(func(a, b *Item) bool {
		return a.Priority > b.Priority
	})
}
