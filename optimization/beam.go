package optimization

import (
	"time"
)

// Beam is like BFS but restricts the search space to save time by only looking at the best nodes.
func Beam(start *State, beamSize int, limit time.Duration) []*Node {
	beam := NewPriorityQueue()
	beam.Push(&Item{State: start, Priority: 0})
	nextStates := NewPriorityQueue()
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
			moves := current.PossibleNextMoves()
			if len(moves) == 0 {
				// terminal state
				if best.Evaluation() < current.Evaluation() {
					best = current.State
				}
				continue
			}

			for _, move := range moves {
				next := current.Apply(move)
				nextStates.Push(&Item{
					State:    next,
					Priority: next.Evaluation(),
				})
			}
		}
		beam, nextStates = nextStates, beam
	}

	if best == nil {
		// in case we saw no terminal state, use the best so far
		best = beam.Pop().State
	}

	var path []*Node
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
