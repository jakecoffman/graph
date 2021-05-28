package optimization

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestMCTS(t *testing.T) {
	world := NewWorld(map1)
	startNode := world.FindOne(Start)
	state := &State{
		World: *world,
		At:    startNode,
	}
	start := time.Now()
	const (
		simulations = 10
		c           = 1.5
		limit       = 100 * time.Millisecond
	)
	path := MCTS(state, simulations, c, limit)

	log.Println("Took", time.Now().Sub(start))
	fmt.Println(world.RenderPath(path))
	fmt.Println("Path is", len(path))
}
