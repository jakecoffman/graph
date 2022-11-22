package ds

import "testing"

func TestHeap(t *testing.T) {
	heap := NewHeap(func(a, b int) bool {
		return a < b
	})
	heap.Push(3)
	heap.Push(1)
	heap.Push(2)
	if heap.Pop() != 1 {
		t.Error("heap.Pop() != 1")
	}
	if heap.Pop() != 2 {
		t.Error("heap.Pop() != 2")
	}
	if heap.Pop() != 3 {
		t.Error("heap.Pop() != 3")
	}
}
