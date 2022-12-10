package pathfinding

import (
	"github.com/jakecoffman/graph"
	"github.com/jakecoffman/graph/ds"
	"github.com/jakecoffman/graph/maze"
)

func less(a, b int) bool {
	return a < b
}

// Astar (or A*) is UCS but applies a heuristic to tell which states are better.
func Astar(start, goal *maze.Node) (path []*maze.Node, found bool) {
	frontier := ds.NewPriorityQueue[*maze.Node](less)
	frontier.Push(ds.NewItem(start, 0))
	cameFrom := map[*maze.Node]*maze.Node{
		start: nil,
	}
	costSoFar := map[*maze.Node]int{
		start: 0,
	}

	for frontier.Len() > 0 {
		current := frontier.Pop()

		if current.State == goal {
			found = true
			break
		}

		for _, next := range current.State.Neighbors {
			newCost := costSoFar[current.State] + maze.Costs[next.Kind]
			if cost, ok := costSoFar[next]; !ok || newCost < cost {
				costSoFar[next] = newCost
				priority := newCost
				// this next line is the only difference between UCS and astar
				priority += graph.ManhattanDistance(goal.Pos, next.Pos)
				frontier.Push(ds.NewItem(next, priority))
				cameFrom[next] = current.State
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
