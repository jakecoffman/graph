package pathfinding

import (
	"container/heap"
)

func UCS(start, goal *Node) []*Node {
	frontier := &PriorityQueue{}
	heap.Init(frontier)
	heap.Push(frontier, &Item{
		Node:     start,
		Priority: 0,
	})
	cameFrom := map[*Node]*Node{}
	costSoFar := map[*Node]int{}
	cameFrom[start] = nil
	costSoFar[start] = 0

	for !frontier.Empty() {
		current := heap.Pop(frontier).(*Item)

		if current.Node == goal {
			break
		}

		for _, next := range current.Neighbors {
			newCost := costSoFar[current.Node] + Costs[next.Kind]
			if cost, ok := costSoFar[next]; !ok || newCost < cost {
				costSoFar[next] = newCost
				priority := newCost
				heap.Push(frontier, &Item{
					Node:     next,
					Priority: priority,
				})
				cameFrom[next] = current.Node
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
