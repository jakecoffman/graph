package pathfinding

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := Stack{}
	stack.Push(&Node{Kind: 1})
	stack.Push(&Node{Kind: 2})
	stack.Push(&Node{Kind: 3})

	v := stack.Pop()
	if v.Kind != 3 {
		t.Error(v)
	}
	v = stack.Pop()
	if v.Kind != 2 {
		t.Error(v)
	}
	v = stack.Pop()
	if v.Kind != 1 {
		t.Error(v)
	}
}
