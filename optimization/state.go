package optimization

// State represents the game state at each moment
type State struct {
	move Move
}

// Move represents an action you can take in the game
type Move struct{}

// PossibleNextMoves returns a list of moves you can make at a state
func (s *State) PossibleNextMoves() []Move {
	return []Move{}
}

// Apply creates a new state that is the result of performing the move on the current state.
func (s *State) Apply(move Move) *State {
	newState := s.Clone()
	newState.move = move
	return newState
}

func (s *State) Clone() *State {
	return &State{}
}

// Evaluation returns an int that represents how good the state is, higher is better.
func (s *State) Evaluation() int {
	return 0
}
