package pathfinding

type Stack []*Node

func (s *Stack) Push(node *Node) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *Node {
	node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return node
}

func (s *Stack) Len() int {
	return len(*s)
}
