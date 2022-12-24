package pathfinding

import (
	"github.com/jakecoffman/graph/ds"
	"math"
)

// AstarArea is A* but targets an area
func AstarArea[T Pathfinder[T]](start T, goals []T) (path []T, found bool) {
	frontier := ds.NewPriorityQueue[T](less)
	frontier.Push(start, 0)
	cameFrom := make(map[T]T)
	costSoFar := map[T]int{
		start: 0,
	}
	var foundGoal T

out:
	for frontier.Len() > 0 {
		current := frontier.Pop()

		for _, goal := range goals {
			if current == goal {
				found = true
				foundGoal = goal
				break out
			}
		}

		current.EachNeighbor(func(next T) {
			newCost := costSoFar[current] + next.Cost()
			if cost, ok := costSoFar[next]; !ok || newCost < cost {
				costSoFar[next] = newCost
				priority := newCost
				closestCost := math.MaxInt64
				for _, goal := range goals {
					dist := goal.Heuristic(next)
					if dist < closestCost {
						closestCost = dist
					}
				}
				priority += closestCost
				frontier.Push(next, priority)
				cameFrom[next] = current
			}
		})
	}

	if !found {
		return
	}

	current := foundGoal
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
