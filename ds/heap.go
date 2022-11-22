package ds

type Heap[T any] struct {
	data    []T
	compare func(a, b T) bool
}

func NewHeap[T any](compare func(a, b T) bool) *Heap[T] {
	return &Heap[T]{compare: compare}
}

func (h *Heap[T]) Clear() {
	h.data = h.data[0:0]
}

func (h *Heap[T]) Empty() bool {
	return h.Len() == 0
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.up(h.Len() - 1)
}

func (h *Heap[T]) Pop() T {
	n := h.Len() - 1
	if n > 0 {
		h.swap(0, n)
		h.down()
	}
	v := h.data[n]
	h.data = h.data[0:n]
	return v
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) up(index int) {
	for {
		i := parent(index)
		if i == index || !h.compare(h.data[index], h.data[i]) {
			break
		}
		h.swap(i, index)
		index = i
	}
}

func (h *Heap[T]) down() {
	n := h.Len() - 1
	i1 := 0
	for {
		j1 := left(i1)
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		j2 := right(i1)
		if j2 < n && h.compare(h.data[j2], h.data[j1]) {
			j = j2
		}
		if !h.compare(h.data[j], h.data[i1]) {
			break
		}
		h.swap(i1, j)
		i1 = j
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return (i * 2) + 1
}

func right(i int) int {
	return left(i) + 1
}
