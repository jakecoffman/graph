package pathfinding

import (
	"github.com/jakecoffman/graph/ds"
	"github.com/jakecoffman/graph/maze"
)

// DFS explores the depth of the tree/graph before the breadth.
// The implementation is identical to BFS except it uses a stack (FILO).
func DFS(start, goal *maze.Node) (path []*maze.Node, found bool) {
	frontier := ds.Stack[*maze.Node]{}
	frontier.Push(start)
	cameFrom := map[*maze.Node]*maze.Node{
		start: nil,
	}

	for frontier.Len() > 0 {
		current := frontier.Pop()

		// early exit
		if current == goal {
			found = true
			break
		}

		for _, next := range current.Neighbors {
			if _, ok := cameFrom[next]; !ok {
				frontier.Push(next)
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
