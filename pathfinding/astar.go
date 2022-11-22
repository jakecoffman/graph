package pathfinding

import (
	"github.com/jakecoffman/graph"
	"github.com/jakecoffman/graph/ds"
)

// Item is an example item for pathfinding. Replace Node with your game data.
type Item struct {
	// Node is the value of the item, it is arbitrary data.
	*Node

	// Priority is set by the user and used by the heap to order.
	Priority int

	// Index is used by the heap.Interface methods to keep things sorted.
	Index int
}

// Astar (or A*) is UCS but applies a heuristic to tell which states are better.
func Astar(start, goal *Node) (path []*Node, found bool) {
	frontier := ds.NewHeap[*Item](func(a, b *Item) bool {
		return a.Priority < b.Priority
	})
	frontier.Push(&Item{Node: start, Priority: 0})
	cameFrom := map[*Node]*Node{
		start: nil,
	}
	costSoFar := map[*Node]int{
		start: 0,
	}

	for frontier.Len() > 0 {
		current := frontier.Pop()

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
				frontier.Push(&Item{Node: next, Priority: priority})
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
