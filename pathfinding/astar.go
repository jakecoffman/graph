package pathfinding

import (
	"github.com/jakecoffman/graph/ds"
)

func less(a, b int) bool {
	return a < b
}

type Pathfinder[T any] interface {
	comparable
	// EachNeighbor calls f for each neighbor of this node.
	EachNeighbor(func(T))
	// Cost returns the cost of moving to this node.
	Cost() int
	// Heuristic returns the estimated cost to the goal.
	// For grid based games this is usually the manhattan distance.
	Heuristic(T) int
}

// Astar (or A*) is UCS but applies a heuristic to tell which states are better.
func Astar[T Pathfinder[T]](start, goal T) (path []T, found bool) {
	frontier := ds.NewPriorityQueue[T](less)
	frontier.Push(start, 0)
	cameFrom := make(map[T]T)
	costSoFar := map[T]int{
		start: 0,
	}

	for frontier.Len() > 0 {
		current := frontier.Pop()

		if current == goal {
			found = true
			break
		}

		current.EachNeighbor(func(next T) {
			newCost := costSoFar[current] + next.Cost()
			if cost, ok := costSoFar[next]; !ok || newCost < cost {
				costSoFar[next] = newCost
				priority := newCost
				// this next line is the only difference between UCS and astar
				priority += goal.Heuristic(next)
				frontier.Push(next, priority)
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
