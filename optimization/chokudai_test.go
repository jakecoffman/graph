package optimization

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestChokudai(t *testing.T) {
	world := NewWorld(map1)
	startNode := world.FindOne(Start)
	state := &State{
		World: *world,
		At:    startNode,
	}
	start := time.Now()
	const (
		width    = 2
		maxTurns = 100
		limit    = 100 * time.Millisecond
	)
	path := Chokudai(state, width, maxTurns, limit)

	log.Println("Took", time.Now().Sub(start))
	fmt.Println(world.RenderPath(path))
	fmt.Println("Path is", len(path))
}
