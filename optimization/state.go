package optimization

import (
	"github.com/jakecoffman/graph"
	"github.com/jakecoffman/graph/maze"
)

// State represents the game state at each moment
type State struct {
	World    maze.World
	At       *maze.Node
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
	if len(s.World.FindAll(maze.Goal)) == 0 {
		return nil
	}

	var validMoves []Move
	for _, n := range s.At.Neighbors {
		if n.Kind == maze.Plain || n.Kind == maze.Start || n.Kind == maze.Goal {
			validMoves = append(validMoves, Move{Pos: n.Pos})
		}
	}
	return validMoves
}

func (s *State) NextStates() []*State {
	var states []*State
	for _, move := range s.PossibleNextMoves() {
		states = append(states, s.Apply(move))
	}
	return states
}

// Apply creates a new state that is the result of performing the move on the current state.
func (s *State) Apply(move Move) *State {
	s.At.Visited++
	newState := s.Clone()
	newState.CameFrom = s
	node := newState.World.At(move.Pos.X, move.Pos.Y)
	if node.Kind == maze.Goal {
		newState.Score++
		node.Kind = maze.Plain
	}
	newState.At = node
	newState.Steps++
	return newState
}

func (s *State) Clone() *State {
	newState := &State{
		World: maze.World{
			Width:  s.World.Width,
			Height: s.World.Height,
			Map:    make([]maze.Node, len(s.World.Map)),
		},
		At:    s.At,
		Score: s.Score,
		Steps: s.Steps,
	}
	copy(newState.World.Map, s.World.Map)
	return newState
}

// Evaluation returns an int that represents how good the state is, higher is better.
func (s *State) Evaluation() int {
	return s.Score*100 - s.At.Visited - s.Steps
}
