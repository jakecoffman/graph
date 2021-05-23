package pathfinding

import "testing"

func TestQueue(t *testing.T) {
	q := Queue{}
	q.Put(&Node{Kind: 1})
	q.Put(&Node{Kind: 2})
	q.Put(&Node{Kind: 3})

	v := q.Pop()
	if v.Kind != 1 {
		t.Error(v)
	}
	v = q.Pop()
	if v.Kind != 2 {
		t.Error(v)
	}
	v = q.Pop()
	if v.Kind != 3 {
		t.Error(v)
	}
}
