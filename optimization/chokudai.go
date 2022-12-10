package optimization

import (
	"github.com/jakecoffman/graph/ds"
	"log"
	"time"
)

// Chokudai is DFS but considers the highest priority nodes first and restricts the search space.
func Chokudai(start *State, width, maxTurns int, limit time.Duration) (path []*Node) {
	pqs := make([]*ds.Heap[*ds.Item[*State]], maxTurns+1)
	for i := 0; i < maxTurns+1; i++ {
		pqs[i] = ds.NewPriorityQueue[*State](greater)
	}
	pqs[0].Push(ds.NewItem(start, 0))
	timeStart := time.Now()

	for time.Now().Sub(timeStart) < limit {
		var processed int
		for depth := 0; depth < maxTurns; depth++ {
			for w := 0; w < width; w++ {
				if pqs[depth].Empty() {
					break
				}

				item := pqs[depth].Pop()
				processed++
				state := item.State

				moves := state.PossibleNextMoves()
				if len(moves) == 0 {
					// terminal state
					pqs[maxTurns].Push(item)
					continue
				}

				for _, move := range moves {
					next := state.Apply(move)
					pqs[depth+1].Push(ds.NewItem(next, next.Evaluation()))
				}
			}
		}
		if processed == 0 {
			// all queues are empty except the final one which contains the end of all the paths
			log.Println("All nodes processed")
			break
		}
		// In my experience increasing chokudai width is a great way to time out
		//width++
	}
	if pqs[maxTurns].Len() == 0 {
		return nil
	}
	best := pqs[maxTurns].Pop().State

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
