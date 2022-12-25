package optimization

import (
	"fmt"
	"github.com/jakecoffman/graph/maze"
	"log"
	"testing"
	"time"
)

func TestMCTS(t *testing.T) {
	world := maze.NewWorld(map1)
	startNode := world.FindOne(maze.Start)
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
	path := MCTS[*State, *maze.Node](state, simulations, c, limit)

	log.Println("Took", time.Now().Sub(start))
	fmt.Println(world.RenderPath(path))
	fmt.Println("Path is", len(path))
}
