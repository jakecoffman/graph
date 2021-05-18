package optimization

import (
	"container/heap"
	"time"
)

const (
	limit    = 10 * time.Millisecond
	maxTurns = 2
)

// Chokudai is like DFS but considers the highest priority nodes first,
// and restricts the search space.
func Chokudai(start *State) (path []*State) {
	pqs := make([]*PriorityQueue, maxTurns+1)
	for i := 0; i < len(pqs); i++ {
		pqs[i] = &PriorityQueue{}
	}
	heap.Push(pqs[0], &Item{
		State:    start,
		Priority: 0,
	})
	cameFrom := map[*State]*State{
		start: nil,
	}
	chokudaiWidth := 1
	timeStart := time.Now()

	for time.Now().Sub(timeStart) < limit {
		var processed int
		for depth := 0; depth < maxTurns; depth++ {
			for w := 0; w < chokudaiWidth; w++ {
				if pqs[depth].Empty() {
					break
				}

				item := pqs[depth].Pop().(*Item)
				state := item.State

				for _, move := range state.PossibleNextMoves() {
					next := state.Apply(move)
					if _, seen := cameFrom[next]; !seen {
						pqs[depth+1].Push(&Item{
							State:    next,
							Priority: state.Evaluation(),
						})
						cameFrom[next] = state
					}
				}
			}
		}
		if processed == 0 {
			// all queues are empty except the final one which contains the end of all the paths
			break
		}
		chokudaiWidth++
	}

	if pqs[maxTurns].Len() == 0 {
		return nil
	}
	best := pqs[maxTurns].Pop().(*Item).State

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
