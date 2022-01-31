package pathfinding

import (
	"container/heap"
	"github.com/jakecoffman/graph"
)

// Astar (or A*) is UCS but applies a heuristic to tell which states are better.
func Astar(start, goal *Node) (path []*Node, found bool) {
	frontier := &PriorityQueue{}
	heap.Push(frontier, &Item{
		Node:     start,
		Priority: 0,
	})
	cameFrom := map[*Node]*Node{
		start: nil,
	}
	costSoFar := map[*Node]int{
		start: 0,
	}

	for !frontier.Empty() {
		current := heap.Pop(frontier).(*Item)

		if current.Node == goal {
			found = true
			break
		}

		for _, next := range current.Neighbors {
			newCost := costSoFar[current.Node] + Costs[next.Kind]
			if cost, ok := costSoFar[next]; !ok || newCost < cost {
				costSoFar[next] = newCost
				priority := newCost
				// this next line is the only difference between UCS and astar
				priority += graph.ManhattanDistance(goal.Pos, next.Pos)
				heap.Push(frontier, &Item{
					Node:     next,
					Priority: priority,
				})
				cameFrom[next] = current.Node
			}
		}
	}

	if !found {
		return
	}

	current := goal
	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return
}
