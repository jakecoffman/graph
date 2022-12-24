package pathfinding

import (
	"github.com/jakecoffman/graph/ds"
)

// UCS or Dijkstra is like BFS but takes cost into account by examining
// lower cost routes.
func UCS[T Pathfinder[T]](start, goal T) (path []T, found bool) {
	// keeps it sorted by priority (low-to-high: min-heap)
	pq := ds.NewPriorityQueue[T](less)
	// push the first item into the pq
	pq.Push(start, 0)
	cameFrom := make(map[T]T)
	costSoFar := map[T]int{
		start: 0,
	}

	for pq.Len() > 0 {
		current := pq.Pop()

		if current == goal {
			found = true
			break
		}

		// push all neighbors into the pq
		current.EachNeighbor(func(next T) {
			// cost is cost of current node plus the next cost
			newCost := costSoFar[current] + next.Cost()
			// if we haven't seen this node yet OR we have but this path was better...
			if cost, ok := costSoFar[next]; !ok || newCost < cost {
				costSoFar[next] = newCost
				pq.Push(next, newCost)
				cameFrom[next] = current
			}
		})
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
