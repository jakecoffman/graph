package optimization

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// MCTS performs a Monte Carlo Tree Search with Upper Confidence Bound.
func MCTS(first *State, simulations int, c float64, limit time.Duration) []*Node {
	start := time.Now()
	root := &MCTSNode{
		state:        first,
		untriedMoves: first.PossibleNextMoves(),
	}

	for time.Now().Sub(start) < limit {
		node := root

		// Selection - find the node with the highest selection score
		for len(node.untriedMoves) == 0 && len(node.children) > 0 {
			sort.Slice(node.children, func(i, j int) bool {
				return node.children[i].selectionScore > node.children[j].selectionScore
			})
			node = node.children[0]
		}

		// Expansion - make a random move on the optimal node
		if len(node.untriedMoves) > 0 {
			i := rand.Intn(len(node.untriedMoves))
			move := node.untriedMoves[i]
			node.untriedMoves = append(node.untriedMoves[:i], node.untriedMoves[i+1:]...)

			newState := node.state.Apply(move)
			child := &MCTSNode{
				parent:       node,
				state:        newState,
				untriedMoves: newState.PossibleNextMoves(),
			}
			node.children = append(node.children, child)

			node = child
		}

		// Simulation - play randomized games from this new state
		sim := node.state
		for j := 0; j < simulations; j++ {
			moves := sim.PossibleNextMoves()
			if len(moves) == 0 {
				break
			}
			i := rand.Intn(len(moves))
			sim = sim.Apply(moves[i])
		}

		// Backpropagation - update the tree to show the results of the play-outs
		outcome := float64(sim.Evaluation())
		p := node
		for p != nil {
			p.totalOutcome += outcome
			p.visits++
			p = p.parent
		}
		winRatio := node.totalOutcome / float64(node.visits)
		node.selectionScore = winRatio + c*math.Sqrt(2*math.Log(float64(node.parent.visits)/float64(node.visits)))
	}

	var path []*Node
	current := root
	for len(current.children) > 0 {
		sort.Slice(current.children, func(i, j int) bool {
			return current.children[i].visits > current.children[j].visits
		})
		path = append(path, current.children[0].state.At)
		current = current.children[0]
	}
	return path
}

type MCTSNode struct {
	parent         *MCTSNode
	state          *State
	totalOutcome   float64
	visits         uint64
	untriedMoves   []Move
	children       []*MCTSNode
	selectionScore float64
}
