package optimization

import (
	"container/heap"
)

// Beam is like BFS but restricts the search space to save time by only looking at the best nodes.
func Beam(beamSize int, start, goal *State) (path []*State, found bool) {
	var beam Queue
	beam.Put(start)
	set := &PriorityQueue{}
	cameFrom := map[*State]*State{
		start: nil,
	}

search:
	for !beam.Empty() {
		set.Clear()
		for b := 0; !beam.Empty(); b++ {
			state := beam.Pop()

			if state == goal {
				found = true
				break search
			}

			for _, move := range state.PossibleNextMoves() {
				next := state.Apply(move)
				if _, seen := cameFrom[next]; !seen {
					heap.Push(set, &Item{
						State:    next,
						Priority: state.Evaluation(),
					})
					cameFrom[next] = state
				}
			}
		}
		beam.Clear()
		// This is where Beam search restricts the search space by only
		// considering nodes that are closer to the goal.
		for i := 0; i < beamSize && !set.Empty(); i++ {
			beam.Put(heap.Pop(set).(*Item).State)
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
