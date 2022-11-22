package ds

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue[*node]{}
	q.Put(&node{Kind: 3})
	q.Put(&node{Kind: 1})
	q.Put(&node{Kind: 2})

	v := q.Pop()
	if v.Kind != 3 {
		t.Error(v)
	}
	v = q.Pop()
	if v.Kind != 1 {
		t.Error(v)
	}
	v = q.Pop()
	if v.Kind != 2 {
		t.Error(v)
	}
}
