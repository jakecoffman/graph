package pathfinding

import (
	"github.com/jakecoffman/graph/ds"
)

// BFS explores the breadth of the tree before the depth.
// The implementation is identical to DFS except BFS uses a Queue (FIFO).
func BFS[T Pathfinder[T]](start, goal T) (path []T, found bool) {
	frontier := ds.Queue[T]{}
	frontier.Put(start)
	cameFrom := map[T]T{}

	for !frontier.Empty() {
		current := frontier.Pop()

		// early exit
		if current == goal {
			found = true
			break
		}

		current.EachNeighbor(func(next T) {
			if _, ok := cameFrom[next]; !ok {
				frontier.Put(next)
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
