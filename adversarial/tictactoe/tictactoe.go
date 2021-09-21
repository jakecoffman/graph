package tictactoe

import (
	"fmt"
	"github.com/jakecoffman/graph/adversarial"
	"log"
	"strings"
)

type Cell int

func (c Cell) String() string {
	if c == CellBlank {
		return " "
	}
	if c == CellX {
		return "X"
	}
	return "O"
}

const (
	CellBlank = Cell(iota)
	CellX
	CellO
)

const (
	width  = 3
	height = 3
)

type State struct {
	board     []Cell
	Current   Cell
	Player    Cell
	CreatedBy int
}

func NewState(player Cell) *State {
	return &State{
		board:     make([]Cell, width*height),
		Player:    player,
		Current:   CellX,
		CreatedBy: -1,
	}
}

func (s *State) String() string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("\n %v │ %v │ %v \n", s.board[0], s.board[1], s.board[2]))
	str.WriteString("───┼───┼───\n")
	str.WriteString(fmt.Sprintf(" %v │ %v │ %v \n", s.board[3], s.board[4], s.board[5]))
	str.WriteString("───┼───┼───\n")
	str.WriteString(fmt.Sprintf(" %v │ %v │ %v ", s.board[6], s.board[7], s.board[8]))
	return str.String()
}

func (s *State) Clone() *State {
	newState := NewState(s.Player)
	copy(newState.board, s.board)
	newState.Current = s.Current
	return newState
}

func (s *State) At(x, y int) Cell {
	return s.board[width*y+x]
}

func (s *State) Set(x, y int, cell Cell) {
	s.board[width*y+x] = cell
}

func (s *State) Index(i int) Cell {
	return s.board[i]
}

func (s *State) Play(i int) *State {
	if s.board[i] != CellBlank {
		log.Panicln("Illegal move", i)
	}
	n := s.Clone()
	n.CreatedBy = i
	n.board[i] = n.Current
	if n.Current == CellX {
		n.Current = CellO
	} else {
		n.Current = CellX
	}
	return n
}

func (s *State) Undo(i int) {
	s.board[i] = CellBlank
	if s.Current == CellX {
		s.Current = CellO
	} else {
		s.Current = CellX
	}
}

func (s *State) IsGameOver() bool {
	score := s.Score()
	return score != 0
}

func (s *State) Score() int {
	x := (s.board[0] == CellX && s.board[1] == CellX && s.board[2] == CellX) ||
		(s.board[3] == CellX && s.board[4] == CellX && s.board[5] == CellX) ||
		(s.board[6] == CellX && s.board[7] == CellX && s.board[8] == CellX) ||
		(s.board[0] == CellX && s.board[3] == CellX && s.board[6] == CellX) ||
		(s.board[1] == CellX && s.board[4] == CellX && s.board[7] == CellX) ||
		(s.board[2] == CellX && s.board[5] == CellX && s.board[8] == CellX) ||
		(s.board[0] == CellX && s.board[4] == CellX && s.board[8] == CellX) ||
		(s.board[2] == CellX && s.board[4] == CellX && s.board[6] == CellX)

	o := (s.board[0] == CellO && s.board[1] == CellO && s.board[2] == CellO) ||
		(s.board[3] == CellO && s.board[4] == CellO && s.board[5] == CellO) ||
		(s.board[6] == CellO && s.board[7] == CellO && s.board[8] == CellO) ||
		(s.board[0] == CellO && s.board[3] == CellO && s.board[6] == CellO) ||
		(s.board[1] == CellO && s.board[4] == CellO && s.board[7] == CellO) ||
		(s.board[2] == CellO && s.board[5] == CellO && s.board[8] == CellO) ||
		(s.board[0] == CellO && s.board[4] == CellO && s.board[8] == CellO) ||
		(s.board[2] == CellO && s.board[4] == CellO && s.board[6] == CellO)

	freeCellsLeft := s.board[0] == CellBlank || s.board[1] == CellBlank || s.board[2] == CellBlank ||
		s.board[3] == CellBlank || s.board[4] == CellBlank || s.board[5] == CellBlank ||
		s.board[6] == CellBlank || s.board[7] == CellBlank || s.board[8] == CellBlank

	switch {
	case x && !o:
		if s.Player == CellX {
			return 10
		} else {
			return -10
		}
	case o && !x:
		if s.Player == CellX {
			return -10
		} else {
			return 10
		}
	case !freeCellsLeft:
		return 1
	default:
		return 0
	}
}

func (s *State) NextStates() []adversarial.GameState {
	var newStates []adversarial.GameState
	for i, cell := range s.board {
		if cell == CellBlank {
			n := s.Play(i)
			newStates = append(newStates, n)
		}
	}
	return newStates
}

func (s *State) BestMove() int {
	best := adversarial.Minimax(s, 1000)
	var legalMoves []int
	for i := range s.board {
		if s.Index(i) == CellBlank {
			legalMoves = append(legalMoves, i)
		}
	}
	return legalMoves[best]
}
