package optimization

import (
	"container/heap"
	"log"
	"time"
)

const (
	limit    = 10 * time.Millisecond
	maxTurns = 2
)

// Chokudai is like DFS but considers the highest priority nodes first.
func Chokudai(start, goal *State) (nextStep *State, found bool) {
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

				for _, next := range state.Neighbors {
					if _, seen := cameFrom[next]; !seen {
						pqs[depth+1].Push(&Item{
							State:    next,
							Priority: state.Priority(),
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

	log.Println("FOUND PATHS:", pqs[maxTurns].Len())

	if pqs[maxTurns].Len() == 0 {
		return
	}
	best := pqs[maxTurns].Pop().(*Item).State
	found = true

	log.Println("BUILDING PATH")
	var path []*State
	current := best
	var i int
	for current != start {
		if i == 10 {
			break
		}
		i++
		path = append(path, current)
		current = cameFrom[current]
	}
	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path[0], true
}
