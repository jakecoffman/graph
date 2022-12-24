package pathfinding

import (
	"github.com/jakecoffman/graph"
	"github.com/jakecoffman/graph/ds"
	"github.com/jakecoffman/graph/maze"
	"math"
)

// AstarArea is A* but targets an area
func AstarArea(start *maze.Node, goals []*maze.Node) (path []*maze.Node, found bool) {
	frontier := ds.NewPriorityQueue[*maze.Node](less)
	frontier.Push(start, 0)
	cameFrom := map[*maze.Node]*maze.Node{
		start: nil,
	}
	costSoFar := map[*maze.Node]int{
		start: 0,
	}
	var foundGoal *maze.Node

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

		for _, next := range current.Neighbors {
			newCost := costSoFar[current] + maze.Costs[next.Kind]
			if cost, ok := costSoFar[next]; !ok || newCost < cost {
				costSoFar[next] = newCost
				priority := newCost
				closestCost := math.MaxInt64
				for _, goal := range goals {
					dist := graph.ManhattanDistance(goal.Pos, next.Pos)
					if dist < closestCost {
						closestCost = dist
					}
				}
				priority += closestCost
				frontier.Push(next, priority)
				cameFrom[next] = current
			}
		}
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
