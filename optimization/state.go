package optimization

type State struct {
	Score     int
	Neighbors []*State
}

func (s *State) Priority() int {
	return s.Score
}
