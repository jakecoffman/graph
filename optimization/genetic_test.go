package optimization

import (
	"fmt"
	"github.com/jakecoffman/graph/pathfinding"
	"log"
	"testing"
	"time"
)

func TestGeneticAlgorithm(t *testing.T) {
	world := pathfinding.NewWorld(map1)
	startNode := world.FindOne(Start)
	state := &State{
		World: *world,
		At:    &VisitedNode{Node: startNode},
	}
	start := time.Now()
	const (
		populationSize = 10
		eliteSize      = 2
		mutationRate   = .5
		limit          = 100 * time.Millisecond
	)
	path := GeneticAlgorithm(state, populationSize, eliteSize, mutationRate, limit)

	log.Println("Took", time.Now().Sub(start))
	fmt.Println(world.RenderPath(path))
	fmt.Println("Path is", len(path))
}
