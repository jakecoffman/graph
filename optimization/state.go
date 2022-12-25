package optimization

import (
	"github.com/jakecoffman/graph/maze"
)

// GameState represents the state of the game at a given point in time.
type GameState[T, U any] interface {
	comparable
	// PossibleNextMoves returns an index of possible moves.
	PossibleNextMoves() []U
	// Evaluation returns a high value for a good state, low value for a bad state.
	Evaluation() int
	// Apply returns a new state with the move applied.
	Apply(U) T
	// CameFrom returns the state that this one was generated from.
	CameFrom() T
	// CreatedBy returns the move that created this state.
	CreatedBy() U
}

// State represents the game state at each moment
type State struct {
	World    maze.World
	At       *maze.Node
	Score    int
	Steps    int
	cameFrom *State
}

// PossibleNextMoves returns a list of moves you can make at a state
func (s *State) PossibleNextMoves() []*maze.Node {
	if len(s.World.FindAll(maze.Goal)) == 0 {
		return nil
	}

	var validMoves []*maze.Node
	for _, n := range s.At.Neighbors {
		if n.Kind == maze.Plain || n.Kind == maze.Start || n.Kind == maze.Goal {
			validMoves = append(validMoves, n)
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
func (s *State) Apply(move *maze.Node) *State {
	s.At.Visited++
	newState := s.Clone()
	newState.cameFrom = s
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

func (s *State) CameFrom() *State {
	return s.cameFrom
}

func (s *State) CreatedBy() *maze.Node {
	return s.cameFrom.At
}
