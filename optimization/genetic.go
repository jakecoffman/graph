package optimization

import (
	"fmt"
	"github.com/jakecoffman/graph"
	"github.com/jakecoffman/graph/maze"
	"log"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Gene is a Goal node
type Gene = *maze.Node

// Chromosome is a route through all Goals
type Chromosome struct {
	Route   []Gene
	Fitness float64
}

// Population is all routes this generation
type Population struct {
	Routes []Chromosome
}

// NewChromosome creates a random sampling of goal nodes
func NewChromosome(w maze.World) Chromosome {
	goals := w.FindAll(maze.Goal)
	permutation := rand.Perm(len(goals))
	var route []*maze.Node
	for i := 0; i < len(goals); i++ {
		route = append(route, goals[permutation[i]])
	}
	return Chromosome{
		Route: append(w.FindAll(maze.Start), route...),
	}
}

// CalcFitness returns 0 for a bad route up to 1 for a good route
func (c *Chromosome) CalcFitness() float64 {
	if c.Fitness != 0 {
		return c.Fitness
	}
	var distance int
	for i := 0; i < len(c.Route)-1; i++ {
		from, to := c.Route[i], c.Route[i+1]
		distance += graph.ManhattanDistance(from.Pos, to.Pos)
	}
	c.Fitness = 1. / float64(distance)
	return c.Fitness
}

// NewPopulation creates the first generation
func NewPopulation(size int, w maze.World) Population {
	var routes []Chromosome
	for i := 0; i < size; i++ {
		routes = append(routes, NewChromosome(w))
	}
	return Population{
		Routes: routes,
	}
}

// Rank sorts the routes by their fitness, best being first
func (p *Population) Rank() {
	sort.Slice(p.Routes, func(i, j int) bool {
		return p.Routes[i].CalcFitness() > p.Routes[j].CalcFitness()
	})
}

// Selection implements Fitness Proportionate Selection, a.k.a Roulette Wheel Selection
func (p *Population) Selection(eliteSize int) []Chromosome {
	selection := p.Routes[:eliteSize]

	p.cumulativeSum()

	for len(selection) < len(p.Routes) {
		pick := rand.Float64()
		for i := len(p.Routes) - 1; i >= 0; i-- {
			if pick <= p.Routes[i].Fitness {
				selection = append(selection, p.Routes[i])
				break
			}
		}
	}

	return selection
}

// cumulativeSum normalizes and calculates cumulative sum
func (p *Population) cumulativeSum() {
	var sumFitness float64
	for i := range p.Routes {
		sumFitness += p.Routes[i].CalcFitness()
	}
	// normalize the fitness
	for i := range p.Routes {
		p.Routes[i].Fitness /= sumFitness
	}
	// calculate cumulative sum (so highest is 1.0) by adding the score of those that come before
	var sum float64
	for i := len(p.Routes) - 1; i >= 0; i-- {
		fitness := p.Routes[i].Fitness
		p.Routes[i].Fitness += sum
		sum += fitness
	}
}

func BreedPopulation(matingPool []Chromosome, eliteSize int) (children Population) {
	length := len(matingPool) - eliteSize
	perm := rand.Perm(len(matingPool))

	for i := 0; i < eliteSize; i++ {
		children.Routes = append(children.Routes, matingPool[i])
	}
	for i := 0; i < length; i++ {
		child := Breed(matingPool[perm[i]], matingPool[perm[len(matingPool)-i-1]])
		children.Routes = append(children.Routes, child)
	}

	return children
}

// Breed implements ordered crossover since the travelling salesman problem we are solving
// involves going through each goal 1 time.
func Breed(parent1, parent2 Chromosome) Chromosome {
	// remove the start for both to avoid mutating it
	startNode := parent1.Route[0]
	parent1Route := parent1.Route[1:]
	parent2Route := parent2.Route[1:]

	start := rand.Intn(len(parent1Route))
	end := rand.Intn(len(parent1Route))

	if start > end {
		start, end = end, start
	}

	var child Chromosome
	child.Route = make([]Gene, 0, len(parent1.Route))
	child.Route = append(child.Route, startNode)
	for i := start; i < end; i++ {
		child.Route = append(child.Route, parent1Route[i])
	}
	for i := range parent2Route {
		goal := parent2Route[i]
		var found bool
		for j := range child.Route {
			if goal == child.Route[j] {
				found = true
				break
			}
		}
		if !found {
			child.Route = append(child.Route, goal)
		}
	}

	return child
}

func (p *Population) MutatePopulation(mutationRate float64) {
	for i := range p.Routes {
		Mutate(&p.Routes[i], mutationRate)
	}
}

// Mutate may modify an individual route with a swap mutation
func Mutate(individual *Chromosome, mutationRate float64) {
	// starting at 1 to skip the start position
	for i := 1; i < len(individual.Route); i++ {
		if rand.Float64() < mutationRate {
			// again, skip the first position
			swap := rand.Intn(len(individual.Route)-1) + 1
			individual.Route[i], individual.Route[swap] = individual.Route[swap], individual.Route[i]
		}
	}
}

func (p *Population) NextGeneration(eliteSize int, mutationRate float64) Population {
	p.Rank()
	matingPool := p.Selection(eliteSize)
	children := BreedPopulation(matingPool, eliteSize)
	children.MutatePopulation(mutationRate)
	return children
}

func GeneticAlgorithm(first *State, populationSize int, eliteSize int, mutationRate float64, limit time.Duration) []*maze.Node {
	start := time.Now()
	p := NewPopulation(populationSize, first.World)
	for time.Now().Sub(start) < limit {
		p = p.NextGeneration(eliteSize, mutationRate)
	}
	p.Rank()
	var total int
	best := p.Routes[0]
	var str strings.Builder
	str.WriteString(fmt.Sprint(best.Route[0].Pos))
	for i := 1; i < len(best.Route); i++ {
		str.WriteString(fmt.Sprint(best.Route[i].Pos))
		total += graph.ManhattanDistance(best.Route[i-1].Pos, best.Route[i].Pos)
	}
	log.Println(str.String(), "Total manhattan distance:", total)
	return p.Routes[0].Route
}
