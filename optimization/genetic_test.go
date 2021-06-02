package optimization

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

func TestGeneticAlgorithm(t *testing.T) {
	log.SetFlags(0)
	world := NewWorld(map1)
	startNode := world.FindOne(Start)
	state := &State{
		World: *world,
		At:    startNode,
	}
	start := time.Now()
	const (
		populationSize = 100
		eliteSize      = 10
		mutationRate   = .1
		limit          = 100 * time.Millisecond
	)
	path := GeneticAlgorithm(state, populationSize, eliteSize, mutationRate, limit)

	log.Println("Took", time.Now().Sub(start)) // best path is 24 (manhattan)
	fmt.Println(world.RenderPath(path))
}

func TestPopulation_cumulativeSum(t *testing.T) {
	p := &Population{Routes: []Chromosome{{
		Fitness: 120.2,
	},{
		Fitness: 83.5,
	},{
		Fitness: 33.3,
	},{
		Fitness: 10,
	}}}
	p.cumulativeSum()

	if p.Routes[0].Fitness != 1.0 {
		t.Error(p.Routes[0].Fitness)
	}
}

func BenchmarkGeneticAlgorithm(b *testing.B) {
	{
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	log.SetFlags(0)
	world := NewWorld(map1)
	startNode := world.FindOne(Start)
	state := &State{
		World: *world,
		At:    startNode,
	}
	const (
		populationSize = 1000
		eliteSize      = 100
		mutationRate   = .3
		limit          = 100 * time.Millisecond
	)
	for i := 0; i < b.N; i++ {
		GeneticAlgorithm(state, populationSize, eliteSize, mutationRate, limit)
	}
}
