package optimization

import (
	"github.com/jakecoffman/graph"
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
	var route []Gene
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
	c.Fitness = 1. / (float64(distance) + 1)
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

	sumFitness := p.Routes[0].Fitness
	cumulativeSum := make([]float64, len(p.Routes))
	cumulativeSum[0] = sumFitness
	for i := 1; i < len(p.Routes); i++ {
		sumFitness += p.Routes[i].Fitness
		cumulativeSum[i] = p.Routes[i].Fitness + cumulativeSum[i-1]
	}
	percentOfSum := make([]float64, len(p.Routes))
	for i := 0; i < len(p.Routes); i++ {
		percentOfSum[i] = 100 * cumulativeSum[i] / sumFitness
	}

	for i := 0; i < len(p.Routes)-eliteSize; i++ {
		pick := rand.Float64() * 100
		for i := 0; i < len(p.Routes); i++ {
			if pick <= percentOfSum[i] {
				selection = append(selection, p.Routes[i])
				break
			}
		}
	}

	return selection
}

func (p *Population) Breed(matingPool []Chromosome, eliteSize int) (children []Chromosome) {
	length := len(matingPool) - eliteSize
	perm := rand.Perm(len(matingPool))

	for i := 0; i < eliteSize; i++ {
		children = append(children, matingPool[i])
	}
	for i := 0; i < length; i++ {
		child := Breed(matingPool[perm[i]], matingPool[perm[len(matingPool)-i-1]])
		children = append(children, child)
	}
	return children
}

// Breed implements ordered crossover since the travelling salesman problem we are solving
// involves going through each goal 1 time.
func Breed(parent1, parent2 Chromosome) Chromosome {
	var child, p1Genes, p2Genes Chromosome

	geneA := rand.Intn(len(parent1.Route))
	geneB := rand.Intn(len(parent2.Route))

	startGene := graph.Min(geneA, geneB)
	endGene := graph.Max(geneA, geneB)

	p1Genes.Route = parent1.Route[startGene:endGene]
	for i := range parent2.Route {
		goal := parent2.Route[i]
		var found bool
		for j := range parent1.Route {
			if goal == parent1.Route[j] {
				found = true
				break
			}
		}
		if !found {
			p2Genes.Route = append(p2Genes.Route, goal)
		}
	}

	child.Route = append(p1Genes.Route, p2Genes.Route...)
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

func (p *Population) NextGeneration(eliteSize int, mutationRate float64) {
	p.Rank()
	matingPool := p.Selection(eliteSize)
	p.Breed(matingPool, eliteSize)
	p.MutatePopulation(mutationRate)
}

func GeneticAlgorithm(first *State, populationSize int, eliteSize int, mutationRate float64, limit time.Duration) []*Node {
	start := time.Now()
	p := NewPopulation(populationSize, first.World)
	for time.Now().Sub(start) < limit {
		p.NextGeneration(eliteSize, mutationRate)
	}
	p.Rank()
	var route []*Node
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
	return route
}
