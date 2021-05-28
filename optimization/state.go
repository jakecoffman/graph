package optimization

import (
	"github.com/jakecoffman/graph"
)

// State represents the game state at each moment
type State struct {
	World    World
	At       *Node
	Score    int
	Steps    int
	CameFrom *State
}

// Move represents an action you can take in the game
type Move struct {
	Pos graph.Pos
}

// PossibleNextMoves returns a list of moves you can make at a state
func (s *State) PossibleNextMoves() []Move {
	if len(s.World.FindAll(Goal)) == 0 {
		return nil
	}

	var validMoves []Move
	for _, n := range s.At.Neighbors {
		if n.Kind == Plain || n.Kind == Start || n.Kind == Goal {
			validMoves = append(validMoves, Move{Pos: n.Pos})
		}
	}
	return validMoves
}

// Apply creates a new state that is the result of performing the move on the current state.
func (s *State) Apply(move Move) *State {
	s.At.Visited++
	newState := s.Clone()
	newState.CameFrom = s
	node := newState.World.At(move.Pos.X, move.Pos.Y)
	if node.Kind == Goal {
		newState.Score++
		node.Kind = Plain
	}
	newState.At = node
	newState.Steps++
	return newState
}

func (s *State) Clone() *State {
	newState := &State{
		World: World{
			width:  s.World.width,
			height: s.World.height,
			world:  make([]Node, len(s.World.world)),
		},
		At:    s.At,
		Score: s.Score,
		Steps: s.Steps,
	}
	copy(newState.World.world, s.World.world)
	return newState
}

// Evaluation returns an int that represents how good the state is, higher is better.
func (s *State) Evaluation() int {
	return s.Score*100 - s.At.Visited - s.Steps
}
