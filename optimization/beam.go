package optimization

import (
	"github.com/jakecoffman/graph/ds"
	"github.com/jakecoffman/graph/maze"
	"time"
)

func greater(a, b int) bool {
	return a > b
}

// Beam is like BFS but restricts the search space to save time by only looking at the best nodes.
func Beam(start *State, beamSize int, limit time.Duration) []*maze.Node {
	beam := ds.NewPriorityQueue[*State](greater)
	beam.Push(ds.NewItem(start, 0))
	nextStates := ds.NewPriorityQueue[*State](greater)
	startTime := time.Now()
	best := start

	for time.Now().Sub(startTime) < limit {
		nextStates.Clear()
		// beam size is restricted
		for b := 0; b < beamSize; b++ {
			if beam.Empty() {
				break
			}
			current := beam.Pop()
			moves := current.State.PossibleNextMoves()
			if len(moves) == 0 {
				// terminal state
				if best.Evaluation() < current.State.Evaluation() {
					best = current.State
				}
				continue
			}

			for _, move := range moves {
				next := current.State.Apply(move)
				nextStates.Push(ds.NewItem(next, next.Evaluation()))
			}
		}
		beam, nextStates = nextStates, beam
	}

	if best == nil {
		// in case we saw no terminal state, use the best so far
		best = beam.Pop().State
	}

	var path []*maze.Node
	current := best
	for current != start {
		path = append(path, current.At)
		current = current.CameFrom
	}
	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
