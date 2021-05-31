package optimization

import (
	"github.com/jakecoffman/graph"
	"log"
	"math/rand"
	"sort"
	"time"
)

// Gene is a Goal node
type Gene = *Node

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
func NewChromosome(w World) Chromosome {
	goals := w.FindAll(Goal)
	goals = append(goals, w.FindOne(Start))
	permutation := rand.Perm(len(goals))
	var route []*Node
	for i := 0; i < len(goals); i++ {
		route = append(route, goals[permutation[i]])
	}
	return Chromosome{
		Route: route,
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
func NewPopulation(size int, w World) Population {
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

	var sumFitness float64
	for i := range p.Routes {
		sumFitness += p.Routes[i].CalcFitness()
	}
	// normalize the fitness
	for i := range p.Routes {
		p.Routes[i].Fitness /= sumFitness
	}
	// calculate cumulative sum (so highest is 1.0) by adding the score of those that come after
	for i := range p.Routes {
		for j := i+1; j < len(p.Routes); j++ {
			p.Routes[i].Fitness += p.Routes[j].Fitness
		}
	}

	for i := 0; i < len(p.Routes)-eliteSize; i++ {
		pick := rand.Float64()
		for i := len(p.Routes)-1; i >= 0; i-- {
			if pick <= p.Routes[i].Fitness {
				selection = append(selection, p.Routes[i])
				break
			}
		}
	}

	return selection
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
	var child Chromosome

	start := rand.Intn(len(parent1.Route))
	end := rand.Intn(len(parent1.Route))

	if start > end {
		start, end = end, start
	}

	for i := start; i < end; i++ {
		child.Route = append(child.Route, parent1.Route[i])
	}
	p2Genes := []*Node{}
	for i := range parent2.Route {
		goal := parent2.Route[i]
		var found bool
		for j := range child.Route {
			if goal.Pos == child.Route[j].Pos {
				found = true
				break
			}
		}
		if !found {
			p2Genes = append(p2Genes, goal)
		}
	}

	child.Route = append(child.Route, p2Genes...)
	return child
}

func (p *Population) MutatePopulation(mutationRate float64) {
	for i := range p.Routes {
		Mutate(p.Routes[i], mutationRate)
	}
}

// Mutate may modify an individual route with a swap mutation
func Mutate(individual Chromosome, mutationRate float64) {
	for i := range individual.Route {
		if rand.Float64() < mutationRate {
			swap := rand.Intn(len(individual.Route))
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

func GeneticAlgorithm(first *State, populationSize int, eliteSize int, mutationRate float64, limit time.Duration) []*Node {
	start := time.Now()
	p := NewPopulation(populationSize, first.World)
	for time.Now().Sub(start) < limit {
		p = p.NextGeneration(eliteSize, mutationRate)
	}
	p.Rank()
	//var route []*Node
	// TODO can't use Astar because Node has diverged
	//goals := p.Routes[0].Route
	//node := first.At
	//for i := range goals {
	//	path, found := pathfinding.Astar(node.Node, goals[i])
	//	if !found {
	//		log.Println("Path not found between", node.Pos, "and", goals[i].Pos)
	//	}
	//	route = append(route, path...)
	//}
	for _, r := range p.Routes {
		var t int
		for i := 0; i < 5; i++ {
			t += graph.ManhattanDistance(r.Route[i].Pos, r.Route[i+1].Pos)
		}
		log.Println("TOTAL", t, "FITNESS", r.CalcFitness())
	}
	for i := range p.Routes[0].Route {
		log.Println(p.Routes[0].Route[i].Pos)
	}
	return p.Routes[0].Route
}
