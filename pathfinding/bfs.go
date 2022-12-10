package pathfinding

import (
	"github.com/jakecoffman/graph/ds"
	"github.com/jakecoffman/graph/maze"
)

// BFS explores the breadth of the tree before the depth.
// The implementation is identical to DFS except BFS uses a Queue (FIFO).
func BFS(start, goal *maze.Node) (path []*maze.Node, found bool) {
	frontier := ds.Queue[*maze.Node]{}
	frontier.Put(start)
	cameFrom := map[*maze.Node]*maze.Node{
		start: nil,
	}

	for !frontier.Empty() {
		current := frontier.Pop()

		// early exit
		if current == goal {
			found = true
			break
		}

		for _, next := range current.Neighbors {
			if _, ok := cameFrom[next]; !ok {
				frontier.Put(next)
				cameFrom[next] = current
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
