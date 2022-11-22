package pathfinding

import (
	"github.com/jakecoffman/graph/ds"
)

// UCS or Dijkstra is like BFS but takes cost into account by examining
// lower cost routes.
func UCS(start, goal *Node) (path []*Node, found bool) {
	// keeps it sorted by priority (low-to-high: min-heap)
	pq := ds.NewHeap(func(a, b *Item) bool {
		return a.Priority < b.Priority
	})
	// push the first item into the pq
	pq.Push(&Item{Node: start, Priority: 0})
	cameFrom := map[*Node]*Node{
		start: nil,
	}
	costSoFar := map[*Node]int{
		start: 0,
	}

	for pq.Len() > 0 {
		current := pq.Pop()

		if current.Node == goal {
			found = true
			break
		}

		// push all neighbors into the pq
		for _, next := range current.Neighbors {
			// cost is cost of current node plus the next cost
			newCost := costSoFar[current.Node] + Costs[next.Kind]
			// if we haven't seen this node yet OR we have but this path was better...
			if cost, ok := costSoFar[next]; !ok || newCost < cost {
				costSoFar[next] = newCost
				pq.Push(&Item{Node: next, Priority: newCost})
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
