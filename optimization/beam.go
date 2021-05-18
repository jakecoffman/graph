package optimization

import (
	"container/heap"
	"time"
)

// Beam is like BFS but restricts the search space to save time by only looking at the best nodes.
func Beam(beamSize int, start *State) []*State {
	beam := &PriorityQueue{}
	heap.Push(beam, start)
	nextStates := &PriorityQueue{}
	cameFrom := map[*State]*State{
		start: nil,
	}
	startTime := time.Now()

	for time.Now().Sub(startTime) < limit {
		nextStates.Clear()
		for b := 0; b < beamSize; b++ {
			if beam.Empty() {
				break
			}
			state := heap.Pop(beam).(*Item).State

			for _, move := range state.PossibleNextMoves() {
				next := state.Apply(move)
				if _, seen := cameFrom[next]; !seen {
					heap.Push(nextStates, &Item{
						State:    next,
						Priority: state.Evaluation(),
					})
					cameFrom[next] = state
				}
			}
		}
		beam = nextStates
	}

	best := heap.Pop(nextStates).(*Item).State

	var path []*State
	current := best
	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
