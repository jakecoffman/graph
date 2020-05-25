package pathfinding

type Queue struct {
	queue []*Node
}

func (q *Queue) Put(n *Node) *Queue {
	q.queue = append(q.queue, n)
	return q
}

func (q *Queue) Get() *Node {
	val := q.queue[0]
	q.queue = q.queue[1:]
	return val
}

func (q *Queue) Empty() bool {
	return len(q.queue) == 0
}

func BFS(world *World, start, goal *Node) []*Node {
	frontier := Queue{}
	frontier.Put(start)
	cameFrom := map[*Node]*Node{
		start: nil,
	}

	for !frontier.Empty() {
		current := frontier.Get()
		for _, next := range current.Neighbors {
			if _, ok := cameFrom[next]; !ok {
				frontier.Put(next)
				cameFrom[next] = current
			}
		}
	}

	current := goal
	var path []*Node
	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
