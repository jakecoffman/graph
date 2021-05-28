package optimization

import (
	"container/heap"
	"log"
	"time"
)

const (
	limit    = 100 * time.Millisecond
	maxTurns = 100
)

// Chokudai is DFS but considers the highest priority nodes first and restricts the search space.
func Chokudai(start *State) (path []*Node) {
	pqs := make([]*PriorityQueue, maxTurns+1)
	for i := 0; i < len(pqs); i++ {
		pqs[i] = &PriorityQueue{}
	}
	heap.Push(pqs[0], &Item{
		State:    start,
		Priority: 0,
	})
	chokudaiWidth := 1
	timeStart := time.Now()

	for time.Now().Sub(timeStart) < limit {
		var processed int
		for depth := 0; depth < maxTurns; depth++ {
			for w := 0; w < chokudaiWidth; w++ {
				if pqs[depth].Empty() {
					break
				}

				item := heap.Pop(pqs[depth]).(*Item)
				processed++
				state := item.State

				moves := state.PossibleNextMoves()
				if len(moves) == 0 {
					// terminal state
					heap.Push(pqs[maxTurns], item)
					continue
				}

				for _, move := range moves {
					next := state.Apply(move)
					heap.Push(pqs[depth+1], &Item{
						State:    next,
						Priority: next.Evaluation(),
					})
				}
			}
		}
		if processed == 0 {
			// all queues are empty except the final one which contains the end of all the paths
			log.Println("All nodes processed")
			break
		}
		// In my experience increasing chokudai width is a great way to time out
		//chokudaiWidth++
	}

	if pqs[maxTurns].Len() == 0 {
		return nil
	}
	best := heap.Pop(pqs[maxTurns]).(*Item).State

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
