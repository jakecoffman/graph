package ds

// NewPriorityQueue creates a heap that is sorted by priority.
func NewPriorityQueue[T any, P comparable](compare func(a, b P) bool) *PriorityQueue[T, P] {
	return &PriorityQueue[T, P]{
		compare: compare,
	}
}

type pqItem[T any, P comparable] struct {
	value    T
	priority P
}

// PriorityQueue will pop items with the highest or lowest priority depending on compare.
type PriorityQueue[T any, P comparable] struct {
	data    []pqItem[T, P]
	compare func(a, b P) bool
}

// Clear removes all items from the queue.
func (h *PriorityQueue[T, P]) Clear() {
	h.data = h.data[0:0]
}

// Empty returns true if the queue is empty.
func (h *PriorityQueue[T, P]) Empty() bool {
	return h.Len() == 0
}

// Len returns the number of items in the queue.
func (h *PriorityQueue[T, P]) Len() int {
	return len(h.data)
}

// Push adds an item to the queue with the given priority.
func (h *PriorityQueue[T, P]) Push(v T, priority P) {
	h.data = append(h.data, pqItem[T, P]{value: v, priority: priority})
	h.up(h.Len() - 1)
}

// Pop removes the item with the highest or lowest priority depending on compare.
func (h *PriorityQueue[T, P]) Pop() T {
	n := h.Len() - 1
	if n > 0 {
		h.swap(0, n)
		h.down()
	}
	v := h.data[n]
	h.data = h.data[0:n]
	return v.value
}

func (h *PriorityQueue[T, P]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *PriorityQueue[T, P]) up(index int) {
	for {
		i := parent(index)
		if i == index || !h.compare(h.data[index].priority, h.data[i].priority) {
			break
		}
		h.swap(i, index)
		index = i
	}
}

func (h *PriorityQueue[T, P]) down() {
	n := h.Len() - 1
	i1 := 0
	for {
		j1 := left(i1)
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		j2 := right(i1)
		if j2 < n && h.compare(h.data[j2].priority, h.data[j1].priority) {
			j = j2
		}
		if !h.compare(h.data[j].priority, h.data[i1].priority) {
			break
		}
		h.swap(i1, j)
		i1 = j
	}
}
