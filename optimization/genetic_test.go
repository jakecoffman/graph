package optimization

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestGeneticAlgorithm(t *testing.T) {
	world := NewWorld(map1)
	startNode := world.FindOne(Start)
	state := &State{
		World: *world,
		At:    startNode,
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

func TestPopulation_Selection(t *testing.T) {
	p := &Population{Routes: []Chromosome{{
		Fitness: 83.5,
	},{
		Fitness: 10,
	},{
		Fitness: 120.2,
	},{
		Fitness: 33.3,
	}}}
	p.Selection(0)
}
