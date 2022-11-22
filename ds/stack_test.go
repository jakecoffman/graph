package ds

import (
	"testing"
)

type node struct {
	Kind int
}

func TestStack(t *testing.T) {
	stack := Stack[*node]{}
	stack.Push(&node{Kind: 3})
	stack.Push(&node{Kind: 1})
	stack.Push(&node{Kind: 2})

	v := stack.Pop()
	if v.Kind != 2 {
		t.Error(v)
	}
	v = stack.Pop()
	if v.Kind != 1 {
		t.Error(v)
	}
	v = stack.Pop()
	if v.Kind != 3 {
		t.Error(v)
	}
}
